package models

import (
	// "gorm.io/driver/postgres"

	"github.com/Hosein110011/go-radar/pkg/config"
	"gorm.io/gorm"
	// pq "github.com/lib/pq"
)

var db *gorm.DB

type Tabler interface {
	TableName() string
}

func init() {
	config.Connect()
	db = config.GetDB()
	// db.AutoMigrate(&User{})
}
