package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API com GO")
	r := router.CreateRouter()

	port := 5000

	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(port), r))
}
