package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100);not null"`
	UserId    uint
	CreatedBy User   `gorm:"foreignKey:UserId"`
	Members   []User `gorm:"many2many:board_members;"`
	Lists     []List
}
