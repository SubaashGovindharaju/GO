package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"coureseid"`
	CourseName  string  ` json:"couresename"`
	CoursePrice int     `json:"price"`
	Author       *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {
fmt.Println("API Documentation")
r :=mux.NewRouter()

courses =append(courses, Course{CourseId: "2", CourseName: "reactjs", CoursePrice: 299,Author: &Author{FullName: "subaash",Website: "http://reactjs"}})
courses =append(courses, Course{CourseId: "4", CourseName: "Mern", CoursePrice: 199,Author: &Author{FullName: "subaash",Website: "http://Mern.com"}})

r.HandleFunc("/", serveHome).Methods("GET")
r.HandleFunc("/courses", getAllCourses).Methods("GET")
r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
r.HandleFunc("/courses", createOneCourse).Methods("POST")
r.HandleFunc("/courses/{id}", UpdateOneCourse).Methods("PUT")
r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")


//listen to port
log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Api by GO </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func  getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("get course one")
		w.Header().Set("Content-Type", "application/json")
params:=mux.Vars(r)

for _,course := range courses{
	if course.CourseId==params["id"]{
	json.NewEncoder(w).Encode(course)
	return	
	}
}
json.NewEncoder(w).Encode("no courses")
return
}

func createOneCourse(w http.ResponseWriter, r *http.Request){

fmt.Println("Create course one")
		w.Header().Set("Content-Type", "application/json")

		if r.Body==nil{
			json.NewEncoder(w).Encode("please send data")
		}

 var course Course
 _ = json.NewDecoder(r.Body).Decode(&course)
 

 if course.IsEmpty(){
				json.NewEncoder(w).Encode("no data inside the json")
 }

// Generate a random CourseID
 course.CourseId = generateRandomCourseID()
 courses = append(courses,course)
 json.NewEncoder(w).Encode(course)
 json.NewEncoder(w).Encode("Added")

 return
}

func generateRandomCourseID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Intn(100))
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request){


	fmt.Println("Update One Course")
		w.Header().Set("Content-Type", "application/json")


		//first - grabid from req
		params :=mux.Vars(r)

		for index,course := range courses{
			if course.CourseId == params["id"]{
				courses = append(courses[:index],courses[index+1:]...)
				var course Course
				_=json.NewDecoder(r.Body).Decode(&course)
				course.CourseId = params["id"]
				courses = append(courses,course)
				json.NewEncoder(w).Encode(course)
				return
			}
}

}


func deleteOneCourse(w http.ResponseWriter, r *http.Request){

	
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params :=mux.Vars(r)

	for index, course :=range courses{
		if course.CourseId== params["id"]{
		courses = append(courses[:index], courses[index+1:]...)
break
}
json.NewEncoder(w).Encode("Deleted json")
	}

}
