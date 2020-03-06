package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func makeResponse(w http.ResponseWriter, body interface{}, statusCode int) {
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
