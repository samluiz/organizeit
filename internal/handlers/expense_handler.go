package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/samluiz/organizeit/internal/adapters"
	"github.com/samluiz/organizeit/internal/common"
	t "github.com/samluiz/organizeit/internal/types"
)

type ExpenseHandler struct {
	ExpenseAdapter *adapters.ExpenseAdapter
	UserAdapter *adapters.UserAdapter
}

func (handler *ExpenseHandler) HandleCreateExpense(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		common.JSONError(w, common.NewError(http.StatusMethodNotAllowed, "Method not allowed", r.URL.Path), http.StatusMethodNotAllowed)
		return
	}

	var expense t.Expense

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&expense); err != nil && err.Error() != "EOF" {
		common.JSONError(w, common.NewError(http.StatusBadRequest, err.Error(), r.URL.Path), http.StatusBadRequest)
		return
	}

	userIdParam := r.URL.Query().Get("userId")
	if userIdParam == "" {
		common.JSONError(w, common.NewError(http.StatusBadRequest, "userId is a required parameter.", r.URL.Path), http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		common.JSONError(w, common.NewError(http.StatusBadRequest, "userId must be a number.", r.URL.Path), http.StatusBadRequest)
		return
	}

	newExpense, err := handler.ExpenseAdapter.CreateExpense(&expense, userId, handler.UserAdapter)

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

	response, err := common.JSONResponse(w, newExpense, http.StatusCreated)

	if err != nil {
		common.JSONError(w, common.NewErrorFromError(err, r.URL.Path), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}