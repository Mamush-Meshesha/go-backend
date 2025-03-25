package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model 
	Title		string `gorm:"not null"`
	Description string
	Completed bool `gorm:"default:false"`
	 UserID      uint `gorm:"not null;default:1"`

}