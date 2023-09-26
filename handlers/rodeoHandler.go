package handlers

import (
	"RodeoApp/models"
	context2 "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"net/http"
	"time"
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
//	    description: Successful operation
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
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the rodeo
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
//	'500':
//	    description: Internal Server Error
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
//   - name: id
//     in: path
//     description: ID of the rodeo
//     required: true
//     type: string
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

// swagger:operation POST /rodeos rodeos insertRodeo
// Insert a new rodeo into the database from JSON received in
// the request body, returns the new rodeo as a response.
// ---
// produce:
// - application/json
// responses:
//
//	'201':
//	    description: Successful creation of resource
//	'400':
//	    description: Invalid input
//	'500':
//	    description: Internal Server Error
func (handler *RodeoHandler) NewRodeoHandler(c *gin.Context) {
	var rodeo models.Rodeo
	if err := c.ShouldBindJSON(&rodeo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error - bad request: ": err.Error(),
		})
		return
	}

	rodeo.ID = primitive.NewObjectID()
	rodeo.PublishedAt = time.Now()
	rodeo.UpdatedAt = time.Now()

	_, err := handler.collection.InsertOne(handler.ctx, rodeo)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error:": "Error while inserting a new rodeo.",
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, rodeo)
}

// swagger:operation PUT /rodeos/{id} rodeos updateRodeo
// Update an existing rodeo whose ID value matches the ID parameter
// sent by the client with the request body from the PUT request, then
// return the updated rodeo as a response.
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the rodeo
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
//	'400':
//	    description: Invalid input
//	'404':
//	    description: Invalid rodeo ID
func (handler *RodeoHandler) UpdateRodeoHandler(c *gin.Context) {
	id := c.Param("id")
	var rodeo models.Rodeo

	if err := c.ShouldBindJSON(&rodeo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	// Take the updated information from the request body, update the
	// rodeo with the updated_at field and current time, save it to the database.
	_, err := handler.collection.UpdateOne(handler.ctx, bson.M{"_id": id},
		bson.D{{"$set", bson.D{
			{"name", rodeo.Name},
			{"pro_rodeo", rodeo.ProRodeo},
			{"start_date", rodeo.StartDate},
			{"end_date", rodeo.EndDate},
			{"venue", rodeo.Venue},
			{"events", rodeo.Events},
			{"updated_at", time.Now()},
		}}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error updating rodeo: ": err.Error()})
	}

	c.JSON(http.StatusOK, rodeo)
}
