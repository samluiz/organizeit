package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samluiz/organizeit/internal/adapters"
	"github.com/samluiz/organizeit/internal/common"
	t "github.com/samluiz/organizeit/internal/types"
)

type UserHandler struct {
	Adapter *adapters.UserAdapter
}

func (handler *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		common.JSONError(w, common.NewError(http.StatusMethodNotAllowed, "Method not allowed", r.URL.Path), http.StatusMethodNotAllowed)
		return
	}

	var user t.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil && err.Error() != "EOF" {
		common.JSONError(w, common.NewError(http.StatusBadRequest, err.Error(), r.URL.Path), http.StatusBadRequest)
		return
	}

	newUser, err := handler.Adapter.CreateUser(user.Email, user.Password)

	if err != nil {
		switch e := err.(type) {
		case *adapters.BusinessRuleError:
				common.JSONError(w, common.NewError(e.StatusCode, e.Error(), r.URL.Path), e.StatusCode)
				return
		default:
			common.JSONError(w, common.NewErrorFromError(err, r.URL.Path), http.StatusInternalServerError)
			return
		}
	}

	response, err := common.JSONResponse(w, newUser, http.StatusCreated)

	if err != nil {
		common.JSONError(w, common.NewErrorFromError(err, r.URL.Path), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (handler *UserHandler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		common.JSONError(w, common.NewError(http.StatusMethodNotAllowed, "Method not allowed", r.URL.Path), http.StatusMethodNotAllowed)
		return
	}

	users, err := handler.Adapter.GetAllUsers()

	if err != nil {
		common.JSONError(w, common.NewErrorFromError(err, r.URL.Path), http.StatusInternalServerError)
		return
	}

	response, err := common.JSONResponse(w, users, http.StatusOK)

	if err != nil {
		common.JSONError(w, common.NewErrorFromError(err, r.URL.Path), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

