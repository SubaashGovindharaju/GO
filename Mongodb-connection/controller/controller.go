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
const dbName = "netflix"
const colName = "watchlist"

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

//insertOneMovie

func insertOneMovie(movie model.Netflix	){
	inserted ,err:=collection.InsertOne(context.Background(),movie)	

	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in database with Id:",inserted.InsertedID)
}

//updateOneMovie

func updateOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}

	result,err:=collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)

	}
fmt.Println("modified count:",result.ModifiedCount)
}

//DeleteOneMovie

func deleteOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}

	deleteCount,err:=collection.DeleteOne(context.Background(),filter)

	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("delete count:",deleteCount)

}

//DeleteAllMovie

func deleteAllMovie()int64{
	
	deleteResult,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)

	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println(" Number of movies delete:",deleteResult)
	return deleteResult.DeletedCount

}

//get all movies from database

func getAllMovies() []primitive.M{
	cur,err:=collection.Find(context.Background(),bson.D{{}})
	if err !=nil{
		log.Fatal(err)
	}

	var movies[]primitive.M
	
	for cur.Next(context.Background()){
	var movie bson.M
	err := cur.Decode(&movie)
	if err !=nil{
		log.Fatal(err)
	}
	movies=append(movies,movie)
	}

	defer cur.Close(context.Background())
	return movies
}


//ACtual contrpller - file


func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
	
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")

	var movie model.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")

	params:=mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "DELETE")
	params:=mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}