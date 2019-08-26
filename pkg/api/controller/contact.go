package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
}

func (*ContactController) Add(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
	return
}

func (*ContactController) Update(ctx *gin.Context) {

}

func (*ContactController) Delete(ctx *gin.Context) {

}

// TODO: Pagination
func (*ContactController) GetAll(ctx *gin.Context) {

}

// Search
func (*ContactController) Search(ctx *gin.Context) {

}
