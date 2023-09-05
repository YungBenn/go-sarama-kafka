package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Hobby string `json:"hobby"`
}
