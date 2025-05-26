package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserId uint   `json:"userId" gorm:"column:user_id"`
	Title  string `json:"title" gorm:"type:varchar(50);"`
}
