package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.CreateRouter()

	port := config.Porta

	fmt.Println("Rodando API com GO na porta ", port, config.StringConexaoBanco)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
