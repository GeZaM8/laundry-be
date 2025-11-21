package routes

import (
	"github.com/GeZaM8/laundry-be/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	ctrl := controllers.OrderController{}

	r.GET("/order", ctrl.GetAllOrder)
	r.GET("/order/:id", ctrl.GetOrder)
	r.POST("/order", ctrl.CreateOrder)
	r.PUT("/order/:id", ctrl.UpdateOrder)
	r.DELETE("order/:id", ctrl.DeleteOrder)
}
