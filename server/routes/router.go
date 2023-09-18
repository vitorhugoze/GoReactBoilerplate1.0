package router

import (
	"net/http"
	"start/controller"
	"start/middleware"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	r := gin.Default()
	r.RedirectTrailingSlash = true

	group := r.Group("/")
	{

		group.GET("/", func(ctx *gin.Context) {
			ctx.Status(http.StatusOK)
		})

		group.POST("/login", controller.LoginController)

		group.POST("/signup", controller.SignupController)

	}

	authGroup := r.Group("/", middleware.AuthUser())
	{

		authGroup.GET("/authuser", func(ctx *gin.Context) {
			ctx.Status(http.StatusOK)
		})

	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return r
}
