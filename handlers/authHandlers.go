package handlers

import (
	"RodeoApp/models"
	"RodeoApp/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"net/http"
)

type AuthHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthHandler(ctx context.Context, collection *mongo.Collection) *AuthHandler {
	return &AuthHandler{
		collection: collection,
		ctx:        ctx,
	}
}

// swagger:operation POST /signin auth signIn
// Login with username and password
// ---
// parameters:
//   - name: username
//     in: json
//     description: username
//     required: true
//     type: string
//   - name: password
//     in: json
//     description: password
//     required: true
//     type: string
//
// produce:
// - application/json
// responses:
//
//	'200':
//	    description: Successful sign in operation
//	'400':
//	    description: Bad Request
//	'401':
//	    description: Unauthorized - Invalid credentials
func (handler *AuthHandler) SignInHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad request error: ": err.Error()})
		return
	}

	var foundUser models.User

	err := handler.collection.FindOne(handler.ctx, bson.M{"username": user.Username}).Decode(&foundUser)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error: ": "invalid username or password"})
		return
	}

	passwordIsValid := utils.CheckPasswordMatch(user.Password, foundUser.Password)
	if passwordIsValid != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error: ": "invalid username or password"})
		return
	}

	// Sessions
	sessionToken := xid.New().String()
	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("token", sessionToken)
	err = session.Save()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message: ": "Howdy, you've been signed in."})
}

// swagger:operation POST /signout auth signout
// Sign out of the application
// ---
// produces:
// - application/json
// response:
//
//	'200':
//	    description: Successful signout from the application
func (handler *AuthHandler) SignOutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message: ": "Thanks for stopping by. Happy Trails!"})
}

// AuthMiddleware Middleware for API authorization functionality
func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionToken := session.Get("token")
		if sessionToken == nil {
			c.JSON(http.StatusForbidden, gin.H{"message: ": "Sorry, you're not logged in."})
			c.Abort()
		}
		c.Next()
	}
}
