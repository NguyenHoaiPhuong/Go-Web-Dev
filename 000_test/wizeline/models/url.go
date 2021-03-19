package models

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	Origin  string `gorm:"unique_index;not null;column:origin" json:"origin"`
	Shorten string `gorm:"unique_index;not null;column:shorten" json:"shorten"`
}
