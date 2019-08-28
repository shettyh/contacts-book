package middleware

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shettyh/contacts-book/pkg/db/dao"
	"golang.org/x/crypto/bcrypt"
)

const (
	errAuthFailed = "failed to authorize user"
)

// AuthHandler is a middleware. It will check and validate the BasicAuth header in the request.
// This will be called for all the routes for which authentication is required.
func AuthHandler(ctx *gin.Context) {
	auth := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errAuthFailed})
		ctx.Abort()
		return
	}

	// Decode and validate credentials
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !validateCredentials(pair[0], pair[1]) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errAuthFailed})
		ctx.Abort()
		return
	}

	// Set the user in the context so the upstream can get the user details.
	ctx.Set("user_id", pair[0])
	ctx.Next()
}

// validateCredentials will check the user credentials against database
func validateCredentials(username, password string) bool {
	// Get the user details from DB
	userDao := new(dao.UserDao)
	user, err := userDao.Get(username)
	if err != nil {
		log.Printf("Not able to get the user details, %v", err)
		return false
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("Password verification failed, %v", err)
		return false
	}

	return true
}
