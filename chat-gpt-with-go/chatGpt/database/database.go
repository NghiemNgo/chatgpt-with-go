package database

import (
	"gorm.io/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)


// DB gorm connector
var DB *gorm.DB
var MG *mongo.Database