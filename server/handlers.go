package server

import (
	"encoding/json"
	operations "go-test/database"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// @Summary	Creates a new director record.
// @Tags		Directors
// @Accept		application/json
// @Produce	application/json
// @Param		Director	body		database.Director	true	"Create Director record"
// @Success	200		{object}	ResponseHTTP{data=database.Director}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/directors/ [post]
func postDirector(w http.ResponseWriter, r *http.Request) {
	var director operations.Director
	err := json.NewDecoder(r.Body).Decode(&director)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error in PostDirector handler\n"))
		log.Printf("Error in PostDirector handler \n%s", err)
		return
	}

	newDirector, err := operations.CreateDirector(director)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in CreateDirector operation\n"))
		log.Printf("Error in CreateDirector operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(newDirector)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in postDirector operation\n"))
		log.Printf("Error in postDirector operation \n%s", err)
		return
	}
}

// @Summary	Fetches director record by id.
// @Tags		Directors
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		string	true	"Get a director record by ID"
// @Success	200		{object}	ResponseHTTP{data=database.Director}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/directors/{id} [get]
func getDirectorById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	director, err := operations.FindFirstDirector(id)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Director not found!\n"))
		log.Printf("Error: Director not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstDirector operation\n"))
		log.Printf("Error in FindFirstDirector operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(director)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getDirectorById operation\n"))
		log.Printf("Error in getDirectorById operation \n%s", err)
		return
	}
}

// @Summary	Fetches all directors.
// @Tags		Directors
// @Accept		application/json
// @Produce	application/json
// @Success	200		{object}	ResponseHTTP{data=database.Director}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/directors/ [get]
func getDirectors(w http.ResponseWriter, r *http.Request) {
	director, err := operations.FindDirectors()
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Directors not found!\n"))
		log.Printf("Error: Directors not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindDirectors operation\n"))
		log.Printf("Error in FindDirectors operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(director)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getDirectors operation\n"))
		log.Printf("Error in getDirectors operation \n%s", err)
		return
	}
}

// @Summary	Updates a Director record.
// @Tags		Directors
// @Accept		application/json
// @Produce	application/json
// @Param		Director	body		database.Director	true	"Update Director record"
// @Success	200		{object}	ResponseHTTP{data=database.Director}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/directors/ [patch]
func patchDirector(w http.ResponseWriter, r *http.Request) {
	var director operations.Director
	err := json.NewDecoder(r.Body).Decode(&director)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error in patchDirector handler\n"))
		log.Printf("Error in patchDirector handler \n%s", err)
		return
	}

	updDirector, err := operations.UpdateDirector(director)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Directors not found!\n"))
		log.Printf("Error: Directors not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in UpdateDirector operation\n"))
		log.Printf("Error in UpdateDirector operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(updDirector)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in patchDirector handler\n"))
		log.Printf("Error in patchDirector handler \n%s", err)
		return
	}
}

// @Summary	Updates a Director record.
// @Tags		Directors
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		string	true	"Delete a director record by ID"
// @Success	200		{object}	ResponseHTTP{data=database.Director}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/directors/{id} [delete]
func deleteDirector(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := operations.DeleteDirector(id)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Director not found!\n"))
		log.Printf("Error: Director not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstDirector operation\n"))
		log.Printf("Error in FindFirstDirector operation \n%s", err)
		return
	}
}
