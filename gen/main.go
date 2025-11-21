package main

import (
	"github.com/GeZaM8/laundry-be/config"
	"gorm.io/gen"
)

func main() {
	config.ConnectDatabase()

	g := gen.NewGenerator(gen.Config{
		OutPath: "./models",
	})
	g.UseDB(config.DB)

	g.GenerateAllTable()

	g.Execute()
}
