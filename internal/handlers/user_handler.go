package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samluiz/organizeit/internal/repositories"
	t "github.com/samluiz/organizeit/internal/types"
)

type UserHandler struct {
	Repository *repositories.UserRepository
}

func (handler *UserHandler) HandlePostUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user t.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	newUser, err := handler.Repository.CreateUser(user.Email, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}