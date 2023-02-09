package database

import (
	"go.tienngay/chatGpt/config"
	"go.tienngay/pkg/mysql/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
)

// ConnectDB connect to db
func MysqlConnectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	db.AutoMigrate(&entities.MistakenVpbankTransaction{})
	db.AutoMigrate(&entities.ReportLogTransaction{})
	log.Println("mysql connected")
	DB = db
}


func MongoConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://"+config.Config("MONGODB_USER")+":"+config.Config("MONGODB_PASSWORD")+"@"+config.Config("MONGODB_HOST")+":"+config.Config("MONGODB_PORT")+"/"+config.Config("MONGODB_DATABASE")+"?connect=direct&authSource="+config.Config("MONGODB_AUTH")+"&authMechanism=SCRAM-SHA-256").SetServerSelectionTimeout(5*time.Second))
	if err != nil {
		cancel()
	}
	log.Println("mongodb connected")
	db := client.Database(config.Config("MONGODB_DATABASE"))
	MG = db
}
