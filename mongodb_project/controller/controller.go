package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
"mongodb_project/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection string and database/collection names

const connectionSring = "Enter your mongodb connection "
const dbname = "Netflix"
const colname = "Watchlist"

// Most Important
var collection *mongo.Collection

// connect with mongoDb
func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionSring)
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection success")
	collection = client.Database(dbname).Collection(colname)
	// collection instance
	fmt.Println("Collection instance is ready")
}

// GetAllMovies - gets all the movies from the MongoDB collection
func GetAllMovies() []primitive.M {
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

// mongodb helpes to connect with the mongodb and perform CRUD operations
// insert one record
func InsertOneMovie(movie model.Netflix) (*mongo.InsertOneResult, error) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
	return inserted, nil
}

// now we update the movie

// UpdateOneMovie - exported function name should be capitalized
func UpdateOneMovie(movieId string, movie model.Netflix) (*mongo.UpdateResult, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": movie}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	fmt.Println("Updated 1 movie in db with id:", result.UpsertedID)
	return result, nil
}

// DeleteOneMovie - exported function name should be capitalized
func DeleteOneMovie(movieID string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	fmt.Println("Deleted 1 movie in db with id:", result.DeletedCount)
	return result, nil
}

// DeleteAllMovies - exported function name should be capitalized
func DeleteAllMovies() (*mongo.DeleteResult, error) {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	fmt.Println("Deleted all movies in db with count:", result.DeletedCount)
	return result, nil
}

// Remove duplicate getAllMovies function since we already have GetAllMovies

// HTTP handler functions
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := GetAllMovies() // use the existing function
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST") // fixed header name

	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inserted, err := InsertOneMovie(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(inserted)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT") // fixed header name

	params := mux.Vars(r)
	result, err := UpdateOneMovie(params["id"], model.Netflix{Watched: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func DeletAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") // fixed header name

	params := mux.Vars(r)
	result, err := DeleteOneMovie(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func DeleteAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") // fixed header name

	result, err := DeleteAllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
