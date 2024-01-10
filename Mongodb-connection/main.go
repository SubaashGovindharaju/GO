package main

import (
	"fmt"
	"log"
	"mongodb/router"
	"net/http"
)

func main() {
	fmt.Println("Mongodb API")
	r:=router.Router()
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000 ...")


} 


