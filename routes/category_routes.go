package routes

import (
	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/controllers"
	"github.com/GeZaM8/laundry-be/middleware"
	"github.com/gin-gonic/gin"
)



func CategoryRoutes(r *gin.RouterGroup) {
	ctrl := controllers.NewCategoryController(config.DB)

	r.GET("/category", ctrl.GetAll)
	r.GET("/category/:id", ctrl.GetByID)
	r.POST("/category", middleware.AuthMiddleware(), ctrl.Create)
	r.PUT("/category/:id", middleware.AuthMiddleware(), ctrl.Update)
	r.DELETE("category/:id", middleware.AuthMiddleware(), ctrl.Delete)
}
