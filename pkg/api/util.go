package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultPageSize = 10
)

// GetPaginationDetailsFromCtx will get the query params related to the
// pagination from the Gin Context.
func GetPaginationDetailsFromCtx(ctx *gin.Context) (pageNo, pageSize int) {
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
