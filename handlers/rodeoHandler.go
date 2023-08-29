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
