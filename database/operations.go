package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// TODO: Map cast (set of actors) to a film — FindCastByFilm

func CreateDirector(director Director) (*Director, error) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		`INSERT INTO directors
		(first_name, middle_name, last_name) 
		VALUES
		($1, $2, $3)
		RETURNING
		(id, first_name, middle_name, last_name)`,
		&director.FirstName, &director.MiddleName, &director.LastName,
	)

	err = row.Scan(&director)
	if err != nil {
		return nil, err
	}
	return &director, nil
}

func FindFirstDirector(id string) (*Director, error) {
	var director Director
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT * FROM directors WHERE id = $1 LIMIT 1`,
		id,
	).Scan(&director.ID, &director.FirstName, &director.MiddleName, &director.LastName)
	if err != nil {
		return nil, err
	}
	return &director, nil
}

func FindDirectors() (*[]Director, error) {
	var directors []Director

	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `SELECT * FROM directors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var director Director
		err = rows.Scan(&director.ID, &director.FirstName, &director.MiddleName, &director.LastName)
		if err != nil {
			return nil, err
		}
		directors = append(directors, director)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &directors, nil
}

func UpdateDirector(director Director) (*Director, error) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		`UPDATE directors SET 
		first_name=$1, middle_name=$2, last_name=$3
		WHERE id = $4`,
		director.FirstName, director.MiddleName, director.LastName, director.ID,
	)
	if ct.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	if err != nil {
		return nil, err
	}
	return &director, nil
}

func DeleteDirector(id string) error {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), `DELETE FROM directors WHERE id=$1`, id)
	if ct.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}
