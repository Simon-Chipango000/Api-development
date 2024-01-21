// main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Simon-Chipango000/Api-Development/db"
	"github.com/go-chi/chi"
)

func main() {

	db.InitializeDB()
	defer db.CloseDB()


	// Example using go-chi
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, go-chi!"))
	})

	// Start the HTTP server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
