package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shettyh/contacts-book/pkg/api/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		// No auth endpoints
		{
			api.POST("/register", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
		}

		// auth endpoints
		{
			basicAuth := api.Group("/")
			basicAuth.Use(middleware.AuthHandler)
			{
				basicAuth.GET("/test", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"message": "oho",
					})
				})
			}
		}
	}

	return r
}
