package router

import "github.com/gorilla/mux"

// Setup устанавливает главный роутер
func Setup() *mux.Router {
	r := mux.NewRouter()
	return r
}
