package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	Card    Card   `gorm:"embedded"`
	User    User   `gorm:"embedded"`
}
