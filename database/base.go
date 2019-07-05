package database

import (
	"fmt"
	"log"
	
	"github.com/jinzhu/gorm"
	// Postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	
	"notez/config"
	"notez/models"
)

var client *gorm.DB

func NewDatabase(conf *config.Config) *gorm.DB {
	
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.DB,
	)
	
	newClient, err := gorm.Open("postgres", connect)
	if err != nil {
		log.Panicln(err)
	}
	
	newClient.AutoMigrate(
		&models.User{},
		&models.Note{},
	)
	
	newClient.LogMode(conf.Debug)
	
	client = newClient
	
	return newClient
}

func GetDB() *gorm.DB {
	return client
}
