package routes

import (
	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authController := controllers.NewAuthController(config.DB)
	r.POST("/login", authController.Login)
}
