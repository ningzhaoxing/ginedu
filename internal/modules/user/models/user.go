package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Articles []Article `gorm:"foreignKey:user_id" json:"articles"`
}
