package main

import (
	"go-rt01-rw06-fe/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.InitRoutes()
	log.Println("Server Started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
