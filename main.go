package main

import (
	handlers "RodeoApp/handlers"
	"github.com/gin-gonic/gin"
)

var rodeosHandler *handlers.RodeoHandler

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

	router.GET("/rodeos", rodeosHandler.ListRodeosHandler)

	router.Run()
}
