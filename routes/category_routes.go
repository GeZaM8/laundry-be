package routes

import (
	"github.com/GeZaM8/laundry-be/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	ctrl := controllers.CategoryController{}

	r.GET("/category", ctrl.GetAll)
	r.GET("/category/:id", ctrl.GetByID)
	r.POST("/category", ctrl.Create)
	r.PUT("/category/:id", ctrl.Update)
	r.DELETE("category/:id", ctrl.Delete)
}
