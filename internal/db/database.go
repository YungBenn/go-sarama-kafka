package db

import (
	"fmt"
	"go-sarama-kafka/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func ConnectDB(c *ConfigDB) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
