package config

import (
	"fmt"
	"restful_api_testing/models"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DB_USER = "root"
const DB_PASS = ""
const DB_HOST = "host.docker.internal"
const DB_PORT = "3306"
const DB_NAME = "crud_go"

const DB_USER_TEST = "root"
const DB_PASS_TEST = ""
const DB_HOST_TEST = "host.docker.internal"
const DB_PORT_TEST = "3306"
const DB_NAME_TEST = "crud_go"

func InitDB() {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	initialMigration()
}

func initialMigration() {
	DB.AutoMigrate(&models.Blog{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
}

func InitDBTest() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER_TEST,
		DB_PASS_TEST,
		DB_HOST_TEST,
		DB_PORT_TEST,
		DB_NAME_TEST,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrationTest()
}

func InitMigrationTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
	DB.Migrator().DropTable(&models.Blog{})
	DB.AutoMigrate(&models.Blog{})
	DB.Migrator().DropTable(&models.Book{})
	DB.AutoMigrate(&models.Book{})
}
