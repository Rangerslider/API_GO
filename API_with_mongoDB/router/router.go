package router

import (
	"github.com/Rangerslider/API/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/allm", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/mo", controller.Createmovie).Methods("POST")
	router.HandleFunc("/api/mo/{id}", controller.Markwatch).Methods("PUT")

	return router

}
