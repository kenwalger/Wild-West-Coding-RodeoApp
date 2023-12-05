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
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo/mongodriver"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var authHandler *handlers.AuthHandler
var rodeosHandler *handlers.RodeoHandler
var webHandler *handlers.WebHandler
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
	rodeoCollection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("RODEO_COLLECTION"))
	rodeosHandler = handlers.NewRodeoHandler(ctx, rodeoCollection)
	usersCollection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("USERS_COLLECTION"))
	authHandler = handlers.NewAuthHandler(ctx, usersCollection)

}

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	// Session Configuration
	sessionCollection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("SESSION_COLLECTION"))
	store := mongodriver.NewStore(sessionCollection, 1800, true, []byte(os.Getenv("SESSION_SECRET")))
	router.Use(sessions.Sessions("RodeoAppSession", store))

	// Web Routes
	router.GET("/", webHandler.IndexHandler)

	// API Auth routes
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/signout", authHandler.SignOutHandler)

	// API Version 1 endpoints and routes
	version1 := router.Group("/api/v1")
	{
		version1.GET("/rodeos", rodeosHandler.ListRodeosHandler)
		version1.GET("/rodeos/:id", rodeosHandler.ListSingleRodeoHandler)

		authorizedV1 := version1.Group("")
		authorizedV1.Use(authHandler.AuthMiddleware())
		{
			authorizedV1.POST("/rodeos", rodeosHandler.NewRodeoHandler)
			authorizedV1.PUT("/rodeos/:id", rodeosHandler.UpdateRodeoHandler)
			authorizedV1.DELETE("/rodeos/:id", rodeosHandler.DeleteRodeoHandler)
		}
	}

	err := router.Run()
	if err != nil {
		return
	}

}
