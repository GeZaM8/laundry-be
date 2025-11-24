package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
    AuthRoutes(r)
	OrderRoutes(r)
	CategoryRoutes(r)
	ReportRoutes(r)
}
