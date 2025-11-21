package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root@tcp(127.0.0.1:3306)/bersihin_express_temp?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Koneksi Database Gagal")
		panic(err)
	}

	fmt.Println("Koneksi Database Berhasil")

	DB = db
}
