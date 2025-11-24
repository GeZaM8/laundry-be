package routes

import (
	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/controllers"
	"github.com/GeZaM8/laundry-be/middleware"
	"github.com/gin-gonic/gin"
)
func ReportRoutes(r *gin.RouterGroup) {
	ctrl := controllers.NewReportController(config.DB)

	reports := r.Group("/reports") 
	{
		reports.GET("/daily", middleware.AuthMiddleware(), ctrl.Daily)
		reports.GET("/revenue", middleware.AuthMiddleware(), ctrl.Revenue)
		reports.GET("/revenue-all", middleware.AuthMiddleware(), ctrl.RevenueAll)
		reports.GET("/items", middleware.AuthMiddleware(), ctrl.Items)
		reports.GET("/items-all", middleware.AuthMiddleware(), ctrl.ItemsAll)

		reports.GET("/chart-monthly", middleware.AuthMiddleware(), ctrl.ChartMonthly)
	}
}
