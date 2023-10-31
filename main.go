package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	DBConnection "github.com/task-manager/DbConnection"
	Auth "github.com/task-manager/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// @title Orders API
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath
func main() {
	dbClient := DBConnection.Db_Connection()
	if err := dbClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		fmt.Println("DB Connected Successfully... 🚀🚀🚀🚀")
		taskTodoDb := dbClient.Database("Todo-Task")
		userCollection := taskTodoDb.Collection("Users")
		// taskCollection := taskTodoDb.Collection("Tasks")

		document := bson.M{"name": "John Doe", "age": 30, "email": "awaisniaz1995@gmail.com"}

		// Insert the document into the collection
		_, err = userCollection.InsertOne(ctx, document)
		if err != nil {
			log.Fatal(err)
		}

	}

	route := mux.NewRouter()
	route.PathPrefix("api/v1")
	route.HandleFunc("/login", Auth.Login).Methods("GET")
	route.HandleFunc("/register", Auth.Register).Methods("GET")
	fmt.Println("Server listening on port 8080... 🚀🚀🚀🚀")
	http.ListenAndServe(":8080", route)
}
