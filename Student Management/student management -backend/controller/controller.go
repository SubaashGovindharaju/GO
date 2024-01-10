package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodb/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "GoStudentManagement"
const colName = "Students"

var collection *mongo.Collection

func init() {
	//client option

	clientOption := options.Client().ApplyURI(connectionString)

	client,err := mongo.Connect(context.TODO(), clientOption)

	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance created successfully")

}

//insertOneStudent

func insertOneStudent(Name model.StudentManage){
	inserted ,err:=collection.InsertOne(context.Background(),Name)	

	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 Student in database with Id:",inserted.InsertedID)
}

//updateOneStudent

// func updateOneStudent(studentId string){
// 	id,_:=primitive.ObjectIDFromHex(studentId)
// 	filter := bson.M{"_id":id}
// 	update:=bson.M{"$set":bson.M{"watched":true}}

// 	result,err:=collection.UpdateOne(context.Background(),filter,update)
// 	if err!=nil{
// 		log.Fatal(err)

// 	}
// fmt.Println("modified count:",result.ModifiedCount)
// }

//DeleteOneStudent

func deleteOneStudent(studentId string){
	id,_:=primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id":id}

	deleteCount,err:=collection.DeleteOne(context.Background(),filter)

	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("delete count:",deleteCount)

}

//DeleteAllStudent

func deleteAllStudent()int64{
	
	deleteResult,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)

	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println(" Number of Students delete:",deleteResult)
	return deleteResult.DeletedCount

}

//get all Students from database

func getAllStudents() []primitive.M{
	cur,err:=collection.Find(context.Background(),bson.D{{}})
	if err !=nil{
		log.Fatal(err)
	}

	var Students[]primitive.M
	
	for cur.Next(context.Background()){
	var Student bson.M
	err := cur.Decode(&Student)
	if err !=nil{
		log.Fatal(err)
	}
	Students=append(Students,Student)
	}

	defer cur.Close(context.Background())
	return Students
}


func getStudentByID(id primitive.ObjectID) (model.StudentManage, error) {
	var student model.StudentManage

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&student)

	return student, err
}

//ACtual contrpller - file


func GetMyAllStudents(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allStudents := getAllStudents()
	json.NewEncoder(w).Encode(allStudents)
	
}


func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	studentID := params["id"]

	objectID, err := primitive.ObjectIDFromHex(studentID)
	if err != nil {
		http.Error(w, "Invalid ObjectID", http.StatusBadRequest)
		return
	}

	student, err := getStudentByID(objectID)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

func CreateStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")

	var Student model.StudentManage
	_=json.NewDecoder(r.Body).Decode(&Student)

	// You can now marshal the struct into JSON with the desired key capitalization
	jsonData, err := json.Marshal(Student)
	
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Use jsonData as needed (e.g., logging, storing in the database, etc.)
	fmt.Println("JSON Data:", string(jsonData))
	
	insertOneStudent(Student)
	json.NewEncoder(w).Encode(Student)
}

// func MarkAsWatched(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Allow-Content-Allow-Methods", "POST")

// 	params:=mux.Vars(r)
// 	updateOneStudent(params["id"])
// 	json.NewEncoder(w).Encode(params["id"])
// }


// updateOneStudent updates a student by ID
func UpdateOneStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	studentID := params["id"]

	objectID, err := primitive.ObjectIDFromHex(studentID)
	if err != nil {
		http.Error(w, "Invalid ObjectID", http.StatusBadRequest)
		return
	}

	// Decode the request body into a StudentManage struct
	var updatedStudent model.StudentManage
	err = json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Update the student in the database
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedStudent}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, "Error updating student", http.StatusInternalServerError)
		return
	}

	// Check if the student was found and updated
	if result.MatchedCount == 0 {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedStudent)
}


func DeleteAStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "DELETE")
	params:=mux.Vars(r)
	deleteOneStudent(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "DELETE")
	count := deleteAllStudent()
	json.NewEncoder(w).Encode(count)
}