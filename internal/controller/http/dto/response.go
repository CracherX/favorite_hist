package dto

import (
	"encoding/json"
	"github.com/CracherX/favorite_hist/internal/entity"
	"net/http"
)

type UserFavoriteResponse struct {
	Favorites []entity.Favorite `json:"favorites"`
}

type e struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Response возвращает сообщение об успехе или ошибке клиенту в json формате.
func Response(w http.ResponseWriter, status int, msg string, details ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := e{
		Status:  status,
		Error:   http.StatusText(status),
		Message: msg,
	}
	if len(details) > 0 {
		errorResponse.Details = details[0]
	}
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(errorResponse)
}

type AuthClientResponse struct {
	UserID int `json:"id"`
}
