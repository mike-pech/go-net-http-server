package server

import (
	"go-test/middleware"

	"log"
	"net/http"
	"time"
)

func Setup(host string) {
	router := http.NewServeMux()
	router.HandleFunc("POST /directors/", postDirector)
	router.HandleFunc("GET /directors/{id}", getDirectorById)
	router.HandleFunc("GET /directors/", getDirectors)
	router.HandleFunc("PATCH /directors/", patchDirector)
	router.HandleFunc("DELETE /directors/{id}", deleteDirector)

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
