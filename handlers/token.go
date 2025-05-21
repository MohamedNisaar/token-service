package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	resp := TokenResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
