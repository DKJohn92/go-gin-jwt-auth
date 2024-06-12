package routes

import (
	controller "github.com/DKJohn92/go-gin-jwt-auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/users/signup", controller.Signup())
	incommingRoutes.POST("/users/login", controller.Login())
}
