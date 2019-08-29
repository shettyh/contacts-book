package util

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPaginationDetailsFromCtx(t *testing.T) {
	ctx := gin.Context{}

	u := url.URL{}

	q := u.Query()
	// Set query params
	q.Set("pageNo", "1")
	q.Set("pageSize", "20")

	u.RawQuery = q.Encode()

	// Set the URL to the ctx
	ctx.Request = new(http.Request)
	ctx.Request.URL = &u

	pageNo, pageSize := GetPaginationDetailsFromCtx(&ctx)
	assert.Equal(t, 1, pageNo)
	assert.Equal(t, 20, pageSize)
}

func TestGetPaginationDetailsFromCtxDefault(t *testing.T) {
	ctx := gin.Context{}

	u := url.URL{}
	q := u.Query()
	u.RawQuery = q.Encode()

	// Set the URL to the ctx
	ctx.Request = new(http.Request)
	ctx.Request.URL = &u

	pageNo, pageSize := GetPaginationDetailsFromCtx(&ctx)
	assert.Equal(t, 0, pageNo)
	assert.Equal(t, 10, pageSize)
}
