package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	CardId  uint
	Card    Card `gorm:"foreignKey:CardId"`
	UserId  uint
	User    User `gorm:"foreignKey:UserId"`
}
