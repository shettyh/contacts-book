package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shettyh/contacts-book/pkg/api/controller"
	"github.com/shettyh/contacts-book/pkg/api/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Create controllers
	userController := new(controller.UserController)
	contactController := new(controller.ContactController)

	api := r.Group("/api/v1")
	{
		// No auth endpoints
		{
			api.PUT("/register", userController.Register)
		}

		// auth endpoints
		{
			contactAPI := api.Group("user/contacts")
			contactAPI.Use(middleware.AuthHandler)
			{
				// TODO: check what methods to use like PUT or POST
				contactAPI.GET("/", contactController.GetAll)
				contactAPI.PUT("/add", contactController.Add)
				contactAPI.POST("/update", contactController.Update)
				contactAPI.DELETE("/:contact_id", contactController.Delete)
				contactAPI.GET("/search", contactController.Search)
			}
		}
	}

	return r
}
