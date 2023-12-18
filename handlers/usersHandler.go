package handlers

import (
	"RodeoApp/models"
	"RodeoApp/utils"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

type UserHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserHandler(ctx context.Context, collection *mongo.Collection) *UserHandler {
	return &UserHandler{
		collection: collection,
		ctx:        ctx,
	}
}

// hashing parameters for the Argon2 hashing algorithm
// for passwords
var hashingParams = &utils.Argon2Parameters{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

// swagger:operation GET /u/register user register
// User registration web page
// ---
//
// produces:
// - application/html
// responses:
//
//	'200':
//	    description: Successful web page rendering
func (handler *UserHandler) ShowRegistrationPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"register.tmpl",
		gin.H{"title": "Register",
			"year": year,
		},
	)
}

// swagger:operation POST /u/register user registerUser
// Processed user registration web page
// ---
// produces:
// - application/html
// responses:
//
//	'200':
//	    description: Successful user registration
//	'400':
//	    description: Unable to register user
func (handler *UserHandler) RegisterUser(c *gin.Context) {
	// Obtain form data
	username := c.PostForm("userName")
	email := c.PostForm("email")
	password := c.PostForm("psw")
	repeatPassword := c.PostForm("psw-repeat")

	if _, err := handler.registerNewUser(username, email, password, repeatPassword); err == nil {
		// Session information
		sessionToken := xid.New().String()
		session := sessions.Default(c)
		session.Set("username", username)
		session.Set("token", sessionToken)
		session.Set("is_logged_in", true)
		err := session.Save()
		if err != nil {
			return
		}
		c.HTML(http.StatusOK,
			"successful-login.tmpl",
			gin.H{"title": "Successful Login",
				"year": year})
	} else {
		c.HTML(http.StatusBadRequest,
			"register.tmpl",
			gin.H{
				"title":        "Register",
				"year":         year,
				"ErrorTitle":   "Registration Failed",
				"ErrorMessage": err.Error(),
			})
	}
}

// registerNewUser adds a user to the database
// if successful returns the user object and no error
// if unsuccessful returns no user object and the specific error as to why the registration doesn't work
func (handler *UserHandler) registerNewUser(username, email, password, repeatpassword string) (*models.User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password cannot be empty")
	} else if strings.TrimSpace(password) != strings.TrimSpace(repeatpassword) {
		return nil, errors.New("passwords don't match")
	} else if !handler.isUserNameAvailable(username) {
		return nil, errors.New("username already exists")
	}

	_, err := handler.collection.InsertOne(handler.ctx, bson.M{
		"username": username,
		"email":    email,
		"password": utils.HashPassword(password, hashingParams),
	})
	if err != nil {
		return nil, errors.New("unable to register user")
	}

	u := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	return &u, nil
}

// isUserNameAvailable checks to see if the username already exists in the database
// returns true if the username is available
// returns false if the username is already taken
func (handler *UserHandler) isUserNameAvailable(username string) bool {
	var foundUser models.User

	err := handler.collection.FindOne(handler.ctx, bson.M{"username": username}).Decode(&foundUser)

	if err != nil {
		return true
	}

	return false
}
