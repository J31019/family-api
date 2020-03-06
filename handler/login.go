package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MonkeyBuisness/family-api/model"
)

func (a API) login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		makeResponse(w, nil, http.StatusBadRequest)
		return
	}
	res := a.service.Login(r.Context(), req)
	statusCode := http.StatusUnauthorized
	if res.Result {
		statusCode = http.StatusOK
	}
	makeResponse(w, res, statusCode)
}
