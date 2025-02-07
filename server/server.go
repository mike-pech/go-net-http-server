package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/swaggo/http-swagger"

	_ "go-test/docs"
	"go-test/middleware"

	"log"
	"net/http"
	"time"
)

var validate *validator.Validate

func Setup(host string) {
	router := http.NewServeMux()

	router.HandleFunc("POST /directors/", postDirector)
	router.HandleFunc("GET /directors/{id}", getDirectorById)
	router.HandleFunc("GET /directors/", getDirectors)
	router.HandleFunc("PATCH /directors/", patchDirector)
	router.HandleFunc("DELETE /directors/{id}", deleteDirector)

	router.HandleFunc("POST /actors/", postActor)
	router.HandleFunc("GET /actors/{id}", getActorById)
	router.HandleFunc("GET /actors/", getActors)
	router.HandleFunc("PATCH /actors/", patchActor)
	router.HandleFunc("DELETE /actors/{id}", deleteActor)

	router.HandleFunc("POST /films/", postFilm)
	router.HandleFunc("GET /films/{id}", getFilmById)
	router.HandleFunc("GET /films/", getFilms)
	router.HandleFunc("PATCH /films/", patchFilm)
	router.HandleFunc("DELETE /films/{id}", deleteFilm)

	router.HandleFunc("POST /characters/", postCharacter)
	router.HandleFunc("GET /characters/{id}", getCharacterById)
	router.HandleFunc("GET /filmCharacters/{filmId}", getCharacterByFilmId)
	router.HandleFunc("GET /characters/", getCharacters)
	router.HandleFunc("PATCH /characters/", patchCharacter)
	router.HandleFunc("DELETE /characters/{id}", deleteCharacter)

	router.HandleFunc("GET /docs/", httpSwagger.Handler(
		httpSwagger.URL("/docs/doc.json"),
		httpSwagger.UIConfig(map[string]string{
			"defaultModelRendering":    `"example"`,
			"defaultModelsExpandDepth": "3",
		}),
	))

	validate = validator.New(validator.WithRequiredStructEnabled())

	stack := middleware.CreateStack(
		middleware.Logging,
		// middleware.AllowCors,
		// middleware.IsAuthed,
		// middleware.CheckPermissions,
	)

	server := http.Server{
		Addr:              host,
		ReadHeaderTimeout: 5000 * time.Millisecond,
		ReadTimeout:       5000 * time.Millisecond,
		Handler:           http.TimeoutHandler(stack(router), 5*time.Second, ""),
	}

	log.Printf("Starting server on port %s\n", host)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n", err)
	}
}
