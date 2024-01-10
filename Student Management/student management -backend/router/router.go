package router

import (
	"mongodb/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router:=mux.NewRouter()

	router.HandleFunc("/api/student",controller.GetMyAllStudents).Methods("GET")
	router.HandleFunc("/api/student",controller.CreateStudent).Methods("POST")
	router.HandleFunc("/api/student/{id}",controller.GetStudentByID).Methods("GET")

	router.HandleFunc("/api/student/{id}",controller.UpdateOneStudent).Methods("PUT")
	router.HandleFunc("/api/student/{id}",controller.DeleteAStudent).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controller.DeleteAllStudent).Methods("DELETE")
	return	router
}