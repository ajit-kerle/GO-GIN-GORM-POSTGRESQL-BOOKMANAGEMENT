package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model 
	Name    string `gorm:"not null" json:"name"`
	Address string `gorm:"not null" json:"address"`
}


