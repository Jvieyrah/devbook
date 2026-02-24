package router

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	router := mux.NewRouter() // gerator de rotas
	return router
}
