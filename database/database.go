package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/uchennaemeruche/ecommerce-with-go-fiber/model"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "root:@tcp(127.0.0.1:3308)/fiberecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}

	log.Println("DB connection succcessful")
	db.Logger = logger.Default.LogMode(logger.Info)

	//  Add migrations
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})

	Database = DbInstance{
		Db: db,
	}

}
