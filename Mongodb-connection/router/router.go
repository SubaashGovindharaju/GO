package router

import (
	"mongodb/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router:=mux.NewRouter()

	router.HandleFunc("/api/movie",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controller.DeleteAllMovie).Methods("DELETE")


	return	router
}