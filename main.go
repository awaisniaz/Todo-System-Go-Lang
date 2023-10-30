package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	DBConnection "github.com/task-manager/DbConnection"
	Auth "github.com/task-manager/auth"
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
		fmt.Println("DB Connected Successfully... 🚀🚀🚀🚀")
		dbClient.Database("testing").Collection("users")
	}

	route := mux.NewRouter()
	// route.PathPrefix("/swagger/*").Handler(httpSwagger.Handler(
	// 	httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The Swagger JSON document URL
	// ))
	route.PathPrefix("api/v1")
	route.HandleFunc("/login", Auth.Login).Methods("GET")
	route.HandleFunc("/register", Auth.Register).Methods("GET")
	fmt.Println("Server listening on port 8080... 🚀🚀🚀🚀")
	http.ListenAndServe(":8080", route)
}
