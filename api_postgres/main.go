package main

import (
	"api_postgres/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 8080....")

	log.Fatal(http.ListenAndServe(":8080", r))
}
