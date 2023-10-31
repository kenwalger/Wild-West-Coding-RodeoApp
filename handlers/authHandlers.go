package handlers

import (
	"RodeoApp/models"
	"RodeoApp/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
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
//	'500':
//	    description: Internal Server Error
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

	expirationTime := time.Now().Add(2 * time.Minute)
	claims := &Claims{
		Username:         user.Username,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}

	c.JSON(http.StatusOK, jwtOutput)
}

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if token == nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

// swagger:operation POST /refreshToken auth refresh
// Refresh the JWT
// ---
// parameter:
//   - name: token
//     in: JSON
//     description: Expired JSON Web Token from the Authorization value in the POST header
//     required: true
//     type: string
//
// produces:
// - application/json
// reponses:
//
//	'200':
//	    description: Successful token refresh
//	'400':
//	    description: Bad Request
//	'401':
//	    description: Unauthorized
//	'500':
//	    description: Internal Server Error
func (handler *AuthHandler) RefreshTokenHandler(c *gin.Context) {
	tokenValue := c.GetHeader("Authorization")
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(
		tokenValue,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error: ": err.Error()})
		return
	}

	if tkn == nil || !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error: ": "invalid token"})
		return
	}

	if claims.ExpiresAt.Sub(time.Now()) > 60*time.Second {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": "token has not yet expired"})
		return
	}

	expirationTime := time.Now().Add(2 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}

	c.JSON(http.StatusOK, jwtOutput)

}
