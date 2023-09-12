package handlers

import (
	"RodeoApp/models"
	context2 "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"net/http"
)

type RodeoHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewRodeoHandler(ctx context.Context, collection *mongo.Collection) *RodeoHandler {
	return &RodeoHandler{
		collection: collection,
		ctx:        ctx,
	}
}

// swagger:operation GET /rodeos rodeos ListRodeos
// Responds with the list of all rodeos as JSON.
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	   description: Successful operation
func (handler *RodeoHandler) ListRodeosHandler(c *gin.Context) {
	cursor, err := handler.collection.Find(handler.ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	defer func(cusor *mongo.Cursor, ctx context2.Context) {
		err := cusor.Close(ctx)
		if err != nil {
			fmt.Println("Error closing cursor.")
		}
	}(cursor, handler.ctx)

	rodeos := make([]models.Rodeo, 0)
	// Iterate through the data cursor of rodeos and append them to the rodeos slice
	for cursor.Next(handler.ctx) {
		var rodeo models.Rodeo
		err := cursor.Decode(&rodeo)
		if err != nil {
			return
		}
		rodeos = append(rodeos, rodeo)
	}

	c.IndentedJSON(http.StatusOK, rodeos)
}

// swagger:operation GET /rodeos/{id} rodeos oneRodeo
// Locate the rodeo whose ID value matches the ID parameter
// sent by the client, then returns the rodeo as a response.
// --
// parameter:
//   - name: id
//     in: path
//     description: ID of the rodeo
//     required: true
//     type: string
//
// produces:
// - application/json
// response:
//
//	'200':
//	   description: Successful operation
//	'500':
//	   description: Internal Server Error
func (handler *RodeoHandler) ListSingleRodeoHandler(c *gin.Context) {
	id := c.Param("id")
	var rodeo models.Rodeo

	cursor := handler.collection.FindOne(handler.ctx, bson.M{"_id": id})

	err := cursor.Decode(&rodeo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error server: ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rodeo)
}

// swagger:operation DELETE /rodeos/{id} rodeos deleteRodeo
// Locate the rodeo whose ID value matches the ID parameter
// sent by the client, delete the rodeo from the database,
// return a successful deletion message.
// ---
// parameter:
// - name: id
// in: path
// description: ID of the rodeo
// required: true
// type: string
//
// produce:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
//	'500':
//	    description: Internal Server Error
func (handler *RodeoHandler) DeleteRodeoHandler(c *gin.Context) {
	id := c.Param("id")

	_, err := handler.collection.DeleteOne(handler.ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error, internal server: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rodeo has been deleted."})
}
