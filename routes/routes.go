package routes

import (
	"github.com/abdulovia/hoff-go/controller"
	"github.com/abdulovia/hoff-go/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PongHandle)

	publicRoutes := r.Group("/auth")
    publicRoutes.POST("/register", controller.Register)
    publicRoutes.POST("/login", controller.Login)
	// r.GET("/profile", handlers.GetProfile)

	// r.POST("/login", handlers.Login)
	// r.POST("/sign-up", handlers.SignUp)
}
