package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100);not null"`
	CreatedBy User   `gorm:"embedded"`
	User      []User `gorm:"many2many:board_members;"`
}
