package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type errResponse struct {
	Error string `json:"error"`
}

func makeResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body == nil {
		return
	}
	data, err := json.Marshal(body)
	if err != nil {
		logrus.Errorf("could not create response: %v", err)
	}
	w.Write(data)
}

func makeErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data, err := json.Marshal(errResponse{
		Error: err.Error(),
	})
	if err != nil {
		logrus.Errorf("could not create response: %v", err)
	}
	w.Write(data)
}
