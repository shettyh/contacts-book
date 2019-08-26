package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shettyh/contacts-book/pkg/db/dao"
	"github.com/shettyh/contacts-book/pkg/db/model"
)

type UserController struct{}

func (*UserController) Register(ctx *gin.Context) {
	var user model.User

	// gin will automatically handle the json error and send 400 error
	ctx.BindJSON(&user)

	userDao := new(dao.UserDao)
	err := userDao.Add(&user)

	//TODO: check gin error handling guideline
	if err != nil {
		log.Printf("Failed to register user, %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	log.Printf("User %s registered successfully", user.Email)
	ctx.JSON(http.StatusOK, "Registered successfully")
}
