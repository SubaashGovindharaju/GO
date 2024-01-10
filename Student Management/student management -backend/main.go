package main

import (
	"fmt"
	"log"
	"mongodb/router"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	fmt.Println("Mongodb API")
	r:=router.Router()

	// Enable CORS for all origins
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"}) // Allow all origins (not recommended for production)

	// Enable CORS middleware
	corsHandler := handlers.CORS(headers, methods, origins)(r)

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":4000",corsHandler))
	fmt.Println("Listening on port 4000 ...")


} 


