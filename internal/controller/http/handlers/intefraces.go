package handlers

import (
	"github.com/CracherX/favorite_hist/internal/entity"
	"net/http"
)

type FavoriteUC interface {
	GetUserFavorite(id int) ([]entity.Favorite, error)
	DeleteFavorite(userID, favID int) error
	AddFavorite(userID, prodID int) error
}

type Validator interface {
	Validate(dto interface{}) error
}

type Logger interface {
	Info(msg string, field ...any)
	Error(msg string, field ...any)
	Debug(msg string, field ...any)
}

type Client interface {
	Get(path string, queryParams ...map[string]string) (*http.Response, error)
}
