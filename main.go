package main

import (
	"go-test/latest/middleware"
	"log"
	"net/http"
)

func findByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("recieved request for item: " + id + "\n"))
}


func coolest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting the coolest!\n"))
}


func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /item/{id}", findByID)
	router.HandleFunc("POST /item/{id}", findByID)
	router.HandleFunc("DELETE /item/{id}", findByID)
	router.HandleFunc("GET godotglobes.su/cool", coolest)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	stack := middleware.CreateStack(
		middleware.Logging,
		// middleware.AllowCors,
		// middleware.IsAuthed,
		// middleware.CheckPermissions,
		// Examples
	)

	server := http.Server{
		Addr:		":8080",
		Handler:	stack(router),
	}
	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
