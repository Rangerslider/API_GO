package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://nuri:nurinuri@cluster0.wdytyko.mongodb.net/?retryWrites=true&w=majority"

const bdName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption) // kon type seita nah janle just use todo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb connection sucessful")

	collection = client.Database(bdName).Collection(colName) // collection er moddhe database r coloum name store kore rakchi jate khujty nah hoi bar bar

}

// helper

func insertOnemovie(movie model.Netflix) { // movie jeita add korbo ==> netflix struct ta koi ashe seitar location
	inserted, err := collection.InsertOne(context.Background(), movie) // database er oparetion korlei context use korty hobe  movie add korbo tai movie ashbe

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted done with id ", inserted.InsertedID) // sdu id dekhbo .. inserted pura struct ke contain korbe
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // movieId string convert korbe objectid te ..jeita mongodb bujhbe
	//bson.B jkn ordered
	// bson.M rest of the time
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}} // filter set kore update korbo ... watched ke ok kore dibo
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count : ", result.ModifiedCount)

}

func deleteOneMoive(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	delete, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted iteam : ", delete.DeletedCount)
}

func getallMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies

}

func GetAllMyMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	all := getallMovies()
	json.NewEncoder(w).Encode(all)

}

func Createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOnemovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func Markwatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
