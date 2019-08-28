package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shettyh/contacts-book/pkg/db/dao"
	"github.com/shettyh/contacts-book/pkg/db/model"
)

// UserController will handle all the User related API's
type UserController struct{}

// Register will register a new user to the contacts book.
// User needs to be registered first to use all other API's
func (*UserController) Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userDao := new(dao.UserDao)
	if err := userDao.Add(&user); err != nil {
		log.Printf("Failed to register user, %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	log.Printf("User %s registered successfully", user.Email)
	ctx.JSON(http.StatusOK, gin.H{"status": "Registered successfully"})
}
