// Rodeos API
//
// This is a sample API about rodeos. Additional information can be found at
// https://github.com/kenwalger/Wild-West-Coding-RodeoApp
//
// Schemes: http
// Host: localhost:8080
// Basepath: /
// Version: 1.0.0
// Contact: Ken W. Alger <kealger@cisco.com>
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// swagger:meta
package main

//goland:noinspection ALL
import (
	handlers "RodeoApp/handlers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var rodeosHandler *handlers.RodeoHandler
var ctx context.Context
var client *mongo.Client

func init() {
	err := godotenv.Load("instance/.env")
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB Atlas.")
	rodeoCollection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION"))
	rodeosHandler = handlers.NewRodeoHandler(ctx, rodeoCollection)

}

func IndexHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello World, welcome to Wild West Coding.",
	})
}

func NameHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.GET("/:name", NameHandler)

	// API Version 1 endpoints and routes
	version1 := router.Group("/api/v1")
	{
		version1.GET("/rodeos", rodeosHandler.ListRodeosHandler)
		version1.POST("/rodeos", rodeosHandler.NewRodeoHandler)
		version1.GET("/rodeos/:id", rodeosHandler.ListSingleRodeoHandler)
		//router.PUT("/rodeos/:id", rodeosHandler.UpdateRodeoHandler)
		version1.DELETE("/rodeos/:id", rodeosHandler.DeleteRodeoHandler)
	}

	router.Run()
}
