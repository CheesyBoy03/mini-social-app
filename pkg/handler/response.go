package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string
}

func JSONError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	response := map[string]string{
		"message": err.Error(),
	}
	json.NewEncoder(w).Encode(response)
}
