package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shettyh/contacts-book/pkg/db/model"

	"github.com/shettyh/contacts-book/pkg/db"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
}

func (*ContactController) Add(ctx *gin.Context) {
	var contact model.Contact
	err := ctx.BindJSON(&contact)
	if err != nil {
		log.Printf("Invalid request json, %v", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// Get user id
	userId, ok := ctx.Get("user_id")
	// not authenticated request
	if !ok {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	dbSession := db.GetSession()
	contact.UserId = userId.(string)
	err = dbSession.Create(&contact).Error
	if err != nil {
		log.Printf("Failed to add contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func (*ContactController) Update(ctx *gin.Context) {
	// Get user id
	userId, ok := ctx.Get("user_id")
	// not authenticated request
	if !ok {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	dbSession := db.GetSession()
	var contact model.Contact
	ctx.BindJSON(&contact)

	contact.UserId = userId.(string)
	err := dbSession.Update(&contact).Error
	if err != nil {
		log.Printf("Failed to update contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func (*ContactController) Delete(ctx *gin.Context) {
	dbSession := db.GetSession()
	// Get contact id
	contactId := ctx.Param("contact_id")

	// Get user id
	userId, ok := ctx.Get("user_id")
	// not authenticated request
	if !ok {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	err := dbSession.Delete(&model.Contact{Email: contactId, UserId: userId.(string)}).Error
	if err != nil {
		log.Printf("Failed to delete contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
	return
}

// TODO: Pagination
func (*ContactController) GetAll(ctx *gin.Context) {
	// Get user id
	userId, ok := ctx.Get("user_id")
	// not authenticated request
	if !ok {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	dbSession := db.GetSession()

	// Pagination query params

	var contacts []model.Contact
	err := dbSession.Where(&model.Contact{UserId: userId.(string)}).Offset(0).Limit(10).Find(&contacts).Error
	if err != nil {
		log.Printf("Failed to update contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

// Search
func (*ContactController) Search(ctx *gin.Context) {
	// Get user id
	userId, ok := ctx.Get("user_id")
	// not authenticated request
	if !ok {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	// Get query params
	emailId := ctx.Query("emailId")
	name := ctx.Query("name")

	dbSession := db.GetSession()

	var contacts []model.Contact
	err := dbSession.Where("user_id = ? AND name LIKE ? AND email like ?",
		userId,
		fmt.Sprintf("%%%s%%", name),
		fmt.Sprintf("%%%s%%", emailId)).Find(&contacts).Error

	if err != nil {
		log.Printf("Failed to update contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, contacts)
	return
}
