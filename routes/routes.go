package routes

import (
	"go-rt01-rw06-fe/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/starter-page", handlers.GetHomeAbout)
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	return router
}
