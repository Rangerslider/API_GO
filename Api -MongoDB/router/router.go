package router

import (
	"github.com/gorilla/mux"
	"github.com/joy-adhikary/API/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/allm", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/mo", controller.Createmovie).Methods("POST")
	router.HandleFunc("/api/mo/{id}", controller.Markwatch).Methods("PUT")

	return router

}
