package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Category(mr *mux.Router, h FavoriteHandler) *mux.Router {
	r := mr.PathPrefix("/favorite").Subrouter()
	r.HandleFunc("", h.GetUserFavorite).Methods(http.MethodGet)
	r.HandleFunc("", h.DeleteFavorite).Methods(http.MethodDelete)
	r.HandleFunc("", h.AddFavorite).Methods(http.MethodPost)
	return r
}
