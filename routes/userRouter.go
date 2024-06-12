package routes

import (
	"github.com/DKJohn92/go-gin-jwt-auth/middleware"

	controller "github.com/DKJohn92/go-gin-jwt-auth/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("/users/:user_id", controller.GetUser())
}
