package router

import "net/http"

type FavoriteHandler interface {
	GetUserFavorite(w http.ResponseWriter, r *http.Request)
	DeleteFavorite(w http.ResponseWriter, r *http.Request)
	AddFavorite(w http.ResponseWriter, r *http.Request)
}
