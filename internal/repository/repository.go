package repository

import (
	"go-sarama-kafka/internal/model"

	"gorm.io/gorm"
)

func InsertUser(db *gorm.DB, user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}