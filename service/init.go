package service

import (
	"gin-app/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func init() {
	DbConnection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	DbConnection.AutoMigrate(&model.Todo{})

}