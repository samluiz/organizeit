package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samluiz/organizeit/config"
	"github.com/samluiz/organizeit/internal/handlers"
	"github.com/samluiz/organizeit/internal/repositories"
)

func main() {

	db, err := config.NewSQLiteConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := repositories.NewRepository(db)

	userHandler := handlers.UserHandler{Repository: repo}

	http.HandleFunc("/users", userHandler.HandlePostUser)
	http.HandleFunc("/", handler)

	PORT := ":8000"
	errs := http.ListenAndServe(PORT, nil)
	fmt.Println(errs)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}