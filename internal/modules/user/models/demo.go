package models

import "gorm.io/gorm"

type Demo struct {
	gorm.Model
	Name string `gorm:"type:varchar(20)"`
	Age  int    `gorm:"type:int(11);default:0"`
}
