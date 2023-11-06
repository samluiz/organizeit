package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samluiz/organizeit/config"
	"github.com/samluiz/organizeit/internal/adapters"
	"github.com/samluiz/organizeit/internal/handlers"
)

func main() {

	db, err := config.NewSQLiteConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userAdapter := adapters.NewUserAdapter(db)
	expenseAdapter := adapters.NewExpenseAdapter(db)

	userHandler := handlers.UserHandler{Adapter: userAdapter}
	expenseHandler := handlers.ExpenseHandler{ExpenseAdapter: expenseAdapter, UserAdapter: userAdapter}

	http.HandleFunc("/api", handler)
	http.HandleFunc("/api/users/create", userHandler.HandleCreateUser)
	http.HandleFunc("/api/users/getAll", userHandler.HandleGetUsers)
	http.HandleFunc("/api/users/get", userHandler.HandleGetUserById)

	http.HandleFunc("/api/expenses/create", expenseHandler.HandleCreateExpense)

	PORT := ":8000"
	log.Default().Println("Server running on port", PORT)
	err = http.ListenAndServe(PORT, nil)
	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}