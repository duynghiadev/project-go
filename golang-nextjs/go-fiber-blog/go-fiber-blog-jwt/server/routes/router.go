package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neerajbg/go-gin-auth/controller"
	"github.com/neerajbg/go-gin-auth/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	// r.GET("/logout", controller.Logout)

	// Blog routes
	r.GET("/", controller.BlogList)
	r.GET("/:id", controller.BlogDetail)
	r.POST("/", controller.BlogCreate)
	r.PUT("/:id", controller.BlogUpdate)
	r.DELETE("/:id", controller.BlogDelete)

	// Private routes
	private := r.Group("/private")
	private.Use(middleware.Authenticate)

	private.GET("/refreshtoken", controller.RefreshToken)
	private.GET("/profile", controller.Profile)
}
