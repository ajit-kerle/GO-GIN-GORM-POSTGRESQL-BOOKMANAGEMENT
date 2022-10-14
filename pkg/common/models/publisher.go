package models

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	PublisherName    string `gorm:"not null" json:"pname"`
	PublisherAddress string `gorm:"not null" json:"paddress"`
}
