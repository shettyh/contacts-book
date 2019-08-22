package handler

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthHandler is a middleware. It will check and validate the BasicAuth header in the request.
// This will be called for all the routes for which authentication is required.
func AuthHandler(ctx *gin.Context) {
	auth := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		ctx.JSON(http.StatusUnauthorized, "Authorization failed")
		ctx.Abort()
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !validateCredentials(pair[0], pair[1]) {
		ctx.JSON(http.StatusUnauthorized, "Authorization failed")
		ctx.Abort()
		return
	}

	// Set the user in the context so the upstream handler can get the user details.
	ctx.Set("user_id", pair[0])
	ctx.Next()
}
