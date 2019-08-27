package controller

import (
	"log"
	"net/http"

	"github.com/shettyh/contacts-book/pkg/api"

	"github.com/shettyh/contacts-book/pkg/db/dao"

	"github.com/shettyh/contacts-book/pkg/db/model"

	"github.com/gin-gonic/gin"
)

type ContactController struct{}

func (*ContactController) Add(ctx *gin.Context) {
	var contact model.Contact
	err := ctx.BindJSON(&contact)
	if err != nil {
		log.Printf("Invalid request json, %v", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// Get user id
	userId, _ := ctx.Get("user_id")
	contact.UserId = userId.(string)

	contactDao := new(dao.ContactDao)
	err = contactDao.Add(&contact)
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
	userId, _ := ctx.Get("user_id")

	// Get the request JSON mapping
	var contact model.Contact
	ctx.BindJSON(&contact)

	contact.UserId = userId.(string)

	contactDao := new(dao.ContactDao)
	err := contactDao.Update(&contact)
	if err != nil {
		log.Printf("Failed to update contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func (*ContactController) Delete(ctx *gin.Context) {
	// Get contact id
	contactId := ctx.Param("contact_id")

	// Get user id
	userId, _ := ctx.Get("user_id")

	contactDao := new(dao.ContactDao)
	err := contactDao.Delete(contactId, userId.(string))
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
	userId, _ := ctx.Get("user_id")

	// Get pagination data
	pageNo, pageSize := api.GetPaginationDetailsFromCtx(ctx)

	// Calculate the offset
	offset := pageNo * pageSize

	contactDao := new(dao.ContactDao)
	contacts, err := contactDao.GetAll(userId.(string), offset, pageSize)
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
	userId, _ := ctx.Get("user_id")

	// Get query params
	emailId := ctx.Query("emailId")
	name := ctx.Query("name")

	// Get pagination data
	pageNo, pageSize := api.GetPaginationDetailsFromCtx(ctx)

	// Calculate the offset
	offset := pageNo * pageSize

	contactsDao := new(dao.ContactDao)
	contacts, err := contactsDao.Search(userId.(string), emailId, name, offset, pageSize)
	if err != nil {
		log.Printf("Failed to update contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, contacts)
	return
}
