package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

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

func CreateActor(actor Actor) (*Actor, error) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		`INSERT INTO actors
		(first_name, middle_name, last_name) 
		VALUES
		($1, $2, $3)
		RETURNING
		(id, first_name, middle_name, last_name)`,
		&actor.FirstName, &actor.MiddleName, &actor.LastName,
	)

	err = row.Scan(&actor)
	if err != nil {
		return nil, err
	}
	return &actor, nil
}

func FindFirstActor(id string) (*Actor, error) {
	var actor Actor
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT * FROM actors WHERE id = $1 LIMIT 1`,
		id,
	).Scan(&actor.ID, &actor.FirstName, &actor.MiddleName, &actor.LastName)
	if err != nil {
		return nil, err
	}
	return &actor, nil
}

func FindActors() (*[]Actor, error) {
	var actors []Actor

	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `SELECT * FROM actors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var actor Actor
		err = rows.Scan(&actor.ID, &actor.FirstName, &actor.MiddleName, &actor.LastName)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &actors, nil
}

func UpdateActor(actor Actor) (*Actor, error) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		`UPDATE actors SET 
		first_name=$1, middle_name=$2, last_name=$3
		WHERE id = $4`,
		actor.FirstName, actor.MiddleName, actor.LastName, actor.ID,
	)
	if ct.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	if err != nil {
		return nil, err
	}
	return &actor, nil
}

func DeleteActor(id string) error {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), `DELETE FROM actors WHERE id=$1`, id)
	if ct.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}

func CreateFilm(film Film) (*Film, error) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		`INSERT INTO films
		(title, directed_by, logline, year) 
		VALUES
		($1, $2, $3, $4)
		RETURNING
		(id, title, directed_by, logline, year)`,
		&film.Title, &film.DirectedBy, &film.Logline, &film.Year,
	)

	err = row.Scan(&film)
	if err != nil {
		return nil, err
	}
	return &film, nil
}

func FindFirstFilm(id string) (*Film, error) {
	var film Film
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT * FROM films WHERE id = $1 LIMIT 1`,
		id,
	).Scan(&film.ID, &film.Title, &film.DirectedBy, &film.Logline, &film.Year)
	if err != nil {
		return nil, err
	}
	return &film, nil
}

func FindFilms() (*[]Film, error) {
	var films []Film

	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `SELECT * FROM films`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var film Film
		err = rows.Scan(&film.ID, &film.Title, &film.DirectedBy, &film.Logline, &film.Year)
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &films, nil
}

func UpdateFilm(film Film) (*Film, error) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		`UPDATE films SET 
		title=$1, directed_by=$2, logline=$3, year=$4
		WHERE id = $5`,
		film.Title, film.DirectedBy, film.Logline, film.Year, film.ID,
	)
	if ct.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	if err != nil {
		return nil, err
	}
	return &film, nil
}

func DeleteFilm(id string) error {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), `DELETE FROM films WHERE id=$1`, id)
	if ct.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}

func CreateCharacter(character Character) (*Character, error) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		`INSERT INTO characters
		(name, portrayed_by, featured_in, dies_in_the_end) 
		VALUES
		($1, $2, $3, $4)
		RETURNING
		(id, name, portrayed_by, featured_in, dies_in_the_end)`,
		&character.Name, &character.PortrayedBy, &character.FeaturedIn, &character.DiesInTheEnd,
	)

	err = row.Scan(&character)
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func FindFirstCharacter(id string) (*Character, error) {
	var character Character
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT * FROM characters WHERE id = $1 LIMIT 1`,
		id,
	).Scan(&character.ID, &character.Name, &character.PortrayedBy, &character.FeaturedIn, &character.DiesInTheEnd)
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func FindCharacters() (*[]Character, error) {
	var characters []Character

	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `SELECT * FROM characters`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var character Character
		err = rows.Scan(&character.ID, &character.Name, &character.PortrayedBy, &character.FeaturedIn, &character.DiesInTheEnd)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &characters, nil
}

func FindCharactersByFilm(filmId string) (*[]Character, error) {
	var characters []Character

	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `SELECT * FROM characters WHERE featured_in=$1`, filmId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var character Character
		err = rows.Scan(&character.ID, &character.Name, &character.PortrayedBy, &character.FeaturedIn, &character.DiesInTheEnd)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &characters, nil
}

func UpdateCharacter(character Character) (*Character, error) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		`UPDATE characters SET 
		name=$1, portrayed_by=$2, featured_in=$3, dies_in_the_end=$4
		WHERE id = $5`,
		character.Name, character.PortrayedBy, character.FeaturedIn, character.DiesInTheEnd, character.ID,
	)
	if ct.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func DeleteCharacter(id string) error {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), `DELETE FROM characters WHERE id=$1`, id)
	if ct.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}
