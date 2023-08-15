package handlers

import (
	"RodeoApp/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type RodeoHandler struct {
	//collection *mongo.Collection
	ctx context.Context
}

func (handler *RodeoHandler) ListRodeosHandler(c *gin.Context) {

	var rodeos = make([]models.Rodeo, 0)

	// Open our jsonFile
	jsonFile, err := os.Open("rodeos.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Error opening rodeo list: ", err.Error())
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &rodeos)
	if err != nil {
		fmt.Println("Error converting file: ", err.Error())
	}

	c.IndentedJSON(http.StatusOK, rodeos)
}
