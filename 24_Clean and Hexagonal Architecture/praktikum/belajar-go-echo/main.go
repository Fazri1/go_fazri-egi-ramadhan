package main

import (
	"belajar-go-echo/route"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	DB_USER := "root"
	DB_PASS := ""
	DB_HOST := "127.0.0.1"
	DB_PORT := "3306"
	DB_NAME := "crud_go"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_PORT,
		DB_NAME)
	var err error
	dbConn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	// var db *gorm.DB
	// db.AutoMigrate(
	// 	model.User{},
	// )

	app := echo.New()
	route.NewRoute(app, dbConn)
	app.Logger.Fatal(app.Start(":8080"))
}
