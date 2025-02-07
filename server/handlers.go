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

// @Summary	Creates a new actor record.
// @Tags		Actors
// @Accept		application/json
// @Produce	application/json
// @Param		Actor	body		database.Actor	true	"Create Actor record"
// @Success	200		{object}	ResponseHTTP{data=database.Actor}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/actors/ [post]
func postActor(w http.ResponseWriter, r *http.Request) {
	var actor operations.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error in PostActor handler\n"))
		log.Printf("Error in PostActor handler \n%s", err)
		return
	}

	newActor, err := operations.CreateActor(actor)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in CreateActor operation\n"))
		log.Printf("Error in CreateActor operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(newActor)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in postActor operation\n"))
		log.Printf("Error in postActor operation \n%s", err)
		return
	}
}

// @Summary	Fetches actor record by id.
// @Tags		Actors
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		string	true	"Get a actor record by ID"
// @Success	200		{object}	ResponseHTTP{data=database.Actor}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/actors/{id} [get]
func getActorById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	actor, err := operations.FindFirstActor(id)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Actor not found!\n"))
		log.Printf("Error: Actor not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstActor operation\n"))
		log.Printf("Error in FindFirstActor operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getActorById operation\n"))
		log.Printf("Error in getActorById operation \n%s", err)
		return
	}
}

// @Summary	Fetches all actors.
// @Tags		Actors
// @Accept		application/json
// @Produce	application/json
// @Success	200		{object}	ResponseHTTP{data=database.Actor}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/actors/ [get]
func getActors(w http.ResponseWriter, r *http.Request) {
	actor, err := operations.FindActors()
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Actors not found!\n"))
		log.Printf("Error: Actors not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindActors operation\n"))
		log.Printf("Error in FindActors operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getActors operation\n"))
		log.Printf("Error in getActors operation \n%s", err)
		return
	}
}

// @Summary	Updates a Actor record.
// @Tags		Actors
// @Accept		application/json
// @Produce	application/json
// @Param		Actor	body		database.Actor	true	"Update Actor record"
// @Success	200		{object}	ResponseHTTP{data=database.Actor}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/actors/ [patch]
func patchActor(w http.ResponseWriter, r *http.Request) {
	var actor operations.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error in patchActor handler\n"))
		log.Printf("Error in patchActor handler \n%s", err)
		return
	}

	updActor, err := operations.UpdateActor(actor)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Actors not found!\n"))
		log.Printf("Error: Actors not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in UpdateActor operation\n"))
		log.Printf("Error in UpdateActor operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(updActor)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in patchActor handler\n"))
		log.Printf("Error in patchActor handler \n%s", err)
		return
	}
}

// @Summary	Updates a Actor record.
// @Tags		Actors
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		string	true	"Delete a actor record by ID"
// @Success	200		{object}	ResponseHTTP{data=database.Actor}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/actors/{id} [delete]
func deleteActor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := operations.DeleteActor(id)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Actor not found!\n"))
		log.Printf("Error: Actor not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstActor operation\n"))
		log.Printf("Error in FindFirstActor operation \n%s", err)
		return
	}
}

// @Summary	Creates a new film record.
// @Tags		Films
// @Accept		application/json
// @Produce	application/json
// @Param		Film	body		database.Film	true	"Create Film record"
// @Success	200		{object}	ResponseHTTP{data=database.Film}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/films/ [post]
func postFilm(w http.ResponseWriter, r *http.Request) {
	var film operations.Film
	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error in PostFilm handler\n"))
		log.Printf("Error in PostFilm handler \n%s", err)
		return
	}

	newFilm, err := operations.CreateFilm(film)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in CreateFilm operation\n"))
		log.Printf("Error in CreateFilm operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(newFilm)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in postFilm operation\n"))
		log.Printf("Error in postFilm operation \n%s", err)
		return
	}
}

// @Summary	Fetches film record by id.
// @Tags		Films
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		string	true	"Get a film record by ID"
// @Success	200		{object}	ResponseHTTP{data=database.Film}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/films/{id} [get]
func getFilmById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	film, err := operations.FindFirstFilm(id)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Film not found!\n"))
		log.Printf("Error: Film not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstFilm operation\n"))
		log.Printf("Error in FindFirstFilm operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getFilmById operation\n"))
		log.Printf("Error in getFilmById operation \n%s", err)
		return
	}
}

// @Summary	Fetches all films.
// @Tags		Films
// @Accept		application/json
// @Produce	application/json
// @Success	200		{object}	ResponseHTTP{data=database.Film}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/films/ [get]
func getFilms(w http.ResponseWriter, r *http.Request) {
	film, err := operations.FindFilms()
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Films not found!\n"))
		log.Printf("Error: Films not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFilms operation\n"))
		log.Printf("Error in FindFilms operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getFilms operation\n"))
		log.Printf("Error in getFilms operation \n%s", err)
		return
	}
}

// @Summary	Updates a Film record.
// @Tags		Films
// @Accept		application/json
// @Produce	application/json
// @Param		Film	body		database.Film	true	"Update Film record"
// @Success	200		{object}	ResponseHTTP{data=database.Film}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/films/ [patch]
func patchFilm(w http.ResponseWriter, r *http.Request) {
	var film operations.Film
	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error in patchFilm handler\n"))
		log.Printf("Error in patchFilm handler \n%s", err)
		return
	}

	updFilm, err := operations.UpdateFilm(film)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Films not found!\n"))
		log.Printf("Error: Films not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in UpdateFilm operation\n"))
		log.Printf("Error in UpdateFilm operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(updFilm)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in patchFilm handler\n"))
		log.Printf("Error in patchFilm handler \n%s", err)
		return
	}
}

// @Summary	Updates a Film record.
// @Tags		Films
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		string	true	"Delete a film record by ID"
// @Success	200		{object}	ResponseHTTP{data=database.Film}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/films/{id} [delete]
func deleteFilm(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := operations.DeleteFilm(id)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Film not found!\n"))
		log.Printf("Error: Film not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstFilm operation\n"))
		log.Printf("Error in FindFirstFilm operation \n%s", err)
		return
	}
}

// @Summary	Creates a new character record.
// @Tags		Characters
// @Accept		application/json
// @Produce	application/json
// @Param		Character	body		database.Character	true	"Create Character record"
// @Success	200		{object}	ResponseHTTP{data=database.Character}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/characters/ [post]
func postCharacter(w http.ResponseWriter, r *http.Request) {
	var character operations.Character
	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error in PostCharacter handler\n"))
		log.Printf("Error in PostCharacter handler \n%s", err)
		return
	}

	newCharacter, err := operations.CreateCharacter(character)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in CreateCharacter operation\n"))
		log.Printf("Error in CreateCharacter operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(newCharacter)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in postCharacter operation\n"))
		log.Printf("Error in postCharacter operation \n%s", err)
		return
	}
}

// @Summary	Fetches character record by id.
// @Tags		Characters
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		string	true	"Get a character record by ID"
// @Success	200		{object}	ResponseHTTP{data=database.Character}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/characters/{id} [get]
func getCharacterById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	character, err := operations.FindFirstCharacter(id)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Character not found!\n"))
		log.Printf("Error: Character not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstCharacter operation\n"))
		log.Printf("Error in FindFirstCharacter operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(character)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getCharacterById operation\n"))
		log.Printf("Error in getCharacterById operation \n%s", err)
		return
	}
}

// @Summary	Fetches all characters.
// @Tags		Characters
// @Accept		application/json
// @Produce	application/json
// @Success	200		{object}	ResponseHTTP{data=database.Character}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/characters/ [get]
func getCharacters(w http.ResponseWriter, r *http.Request) {
	character, err := operations.FindCharacters()
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Characters not found!\n"))
		log.Printf("Error: Characters not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindCharacters operation\n"))
		log.Printf("Error in FindCharacters operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(character)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getCharacters operation\n"))
		log.Printf("Error in getCharacters operation \n%s", err)
		return
	}
}

// @Summary	Fetches character record by film id.
// @Tags		Characters
// @Accept		application/json
// @Produce	application/json
// @Param		filmId	path		string	true	"Get a character record by Film ID"
// @Success	200		{object}	ResponseHTTP{data=database.Character}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/filmCharacters/{filmId} [get]
func getCharacterByFilmId(w http.ResponseWriter, r *http.Request) {
	filmId := r.PathValue("filmId")

	character, err := operations.FindCharactersByFilm(filmId)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Characters not found!\n"))
		log.Printf("Error: Characters not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstCharacter operation\n"))
		log.Printf("Error in FindFirstCharacter operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(character)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in getCharacterById operation\n"))
		log.Printf("Error in getCharacterById operation \n%s", err)
		return
	}
}

// @Summary	Updates a Character record.
// @Tags		Characters
// @Accept		application/json
// @Produce	application/json
// @Param		Character	body		database.Character	true	"Update Character record"
// @Success	200		{object}	ResponseHTTP{data=database.Character}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/characters/ [patch]
func patchCharacter(w http.ResponseWriter, r *http.Request) {
	var character operations.Character
	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error in patchCharacter handler\n"))
		log.Printf("Error in patchCharacter handler \n%s", err)
		return
	}

	updCharacter, err := operations.UpdateCharacter(character)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Characters not found!\n"))
		log.Printf("Error: Characters not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in UpdateCharacter operation\n"))
		log.Printf("Error in UpdateCharacter operation \n%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(updCharacter)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in patchCharacter handler\n"))
		log.Printf("Error in patchCharacter handler \n%s", err)
		return
	}
}

// @Summary	Updates a Character record.
// @Tags		Characters
// @Accept		application/json
// @Produce	application/json
// @Param		id	path		string	true	"Delete a character record by ID"
// @Success	200		{object}	ResponseHTTP{data=database.Character}
// @Failure	400		{object}	ResponseHTTP{}
// @Failure	418		{object}	ResponseHTTP{}
// @Failure	500		{object}	ResponseHTTP{}
// @Router		/characters/{id} [delete]
func deleteCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := operations.DeleteCharacter(id)
	if err == pgx.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("Error: Character not found!\n"))
		log.Printf("Error: Character not found!\n%s", err)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error in FindFirstCharacter operation\n"))
		log.Printf("Error in FindFirstCharacter operation \n%s", err)
		return
	}
}
