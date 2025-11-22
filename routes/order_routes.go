package routes

import (
	"github.com/GeZaM8/laundry-be/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	ctrl := controllers.OrderController{}

	r.GET("/order", ctrl.GetAll)
	r.GET("/order/:id", ctrl.GetByID)
	r.POST("/order", ctrl.Create)
	r.PUT("/order/:id", ctrl.Update)
	r.DELETE("order/:id", ctrl.Delete)
}
