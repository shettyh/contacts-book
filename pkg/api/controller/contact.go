package controller

import (
	"log"
	"net/http"

	"github.com/shettyh/contacts-book/pkg/api"

	"github.com/shettyh/contacts-book/pkg/db/dao"

	"github.com/shettyh/contacts-book/pkg/db/model"

	"github.com/gin-gonic/gin"
)

// ContactController will handle all `api/v1//user/contacts` APIs
type ContactController struct{}

// Add will adds the new contact to the Users contact book
func (*ContactController) Add(ctx *gin.Context) {
	userId := ctx.GetString("user_id")

	// Get the request JSON
	var contact model.Contact
	if err := ctx.BindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Set the user ID
	contact.UserId = userId

	// Add to DB
	contactDao := new(dao.ContactDao)
	if err := contactDao.Add(&contact); err != nil {
		log.Printf("Failed to add contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// Update will try to update the existing contact for the specific user
func (*ContactController) Update(ctx *gin.Context) {
	userId := ctx.GetString("user_id")

	// Get the request JSON mapping
	var contact model.Contact
	if err := ctx.BindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Set the user ID
	contact.UserId = userId

	contactDao := new(dao.ContactDao)
	if err := contactDao.Update(&contact); err != nil {
		log.Printf("Failed to update contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// Delete will take the contact email id as argument in the URL
// and tries to delete that contact.
func (*ContactController) Delete(ctx *gin.Context) {
	userId := ctx.GetString("user_id")

	// Get contact id
	contactId := ctx.Param("contact_id")

	contactDao := new(dao.ContactDao)
	if err := contactDao.Delete(contactId, userId); err != nil {
		log.Printf("Failed to delete contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// GetAll will returns all the contacts for the user.
func (*ContactController) GetAll(ctx *gin.Context) {
	userId := ctx.GetString("user_id")

	// Get pagination data
	pageNo, pageSize := api.GetPaginationDetailsFromCtx(ctx)
	// Calculate the offset
	offset := pageNo * pageSize

	contactDao := new(dao.ContactDao)
	contacts, err := contactDao.GetAll(userId, offset, pageSize)
	if err != nil {
		log.Printf("Failed to update contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

// Search will search for the contact details.
// Search can work with EmailId or Name of the contact or Both.
func (*ContactController) Search(ctx *gin.Context) {
	userId := ctx.GetString("user_id")

	// Get query params
	emailId := ctx.Query("emailId")
	name := ctx.Query("name")

	// Get pagination data
	pageNo, pageSize := api.GetPaginationDetailsFromCtx(ctx)
	// Calculate the offset
	offset := pageNo * pageSize

	contactsDao := new(dao.ContactDao)
	contacts, err := contactsDao.Search(userId, emailId, name, offset, pageSize)
	if err != nil {
		log.Printf("Failed to update contact. %v", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}
