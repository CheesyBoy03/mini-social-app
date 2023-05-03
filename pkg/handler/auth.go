package handler

import (
	"encoding/json"
	"net/http"
)

const (
	maxUploadSize = 5 << 20 // 5 megabytes
)

type authSignInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var inp authSignInInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusNotFound)
		return
	}

	token, err := h.services.SignIn(inp.Login, inp.Password)
	if err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"token": token,
	}
	json.NewEncoder(w).Encode(response)
}

type authSignUpInput struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var inp authSignUpInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.services.SignUp(inp.Email, inp.Password, inp.FirstName, inp.LastName)
	if err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
