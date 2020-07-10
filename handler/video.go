package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a API) videoList(w http.ResponseWriter, r *http.Request) {
	res, err := a.service.VideoList(r.Context())
	if err != nil {
		makeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	makeResponse(w, res, http.StatusOK)
}

func (a API) videoByID(w http.ResponseWriter, r *http.Request) {
	videoID := mux.Vars(r)["id"]
	res, err := a.service.VideoURLByID(r.Context(), videoID)
	if err != nil {
		makeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	makeResponse(w, res, http.StatusOK)
}
