package main

import (
	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8000")
}
