package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter() // gerator de rotas
	rotas.Configurar(router)
	return router
}
