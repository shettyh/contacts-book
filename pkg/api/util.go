package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultPageSize = 10
)

func GetPaginationDetailsFromCtx(ctx *gin.Context) (pageNo, pageSize int) {
	// Get the page number
	pageNo, err := strconv.Atoi(ctx.Query("pageNo"))
	if err != nil {
		pageNo = 0
	}

	// get the page size
	pageSize, err = strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		pageSize = defaultPageSize
	}

	return
}
