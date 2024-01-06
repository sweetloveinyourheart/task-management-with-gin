package models

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Position    *int   `gorm:"type:int"`
	DueDate     *time.Time
	ListId      uint
	List        List `gorm:"foreignKey:ListId"`
	Comments    []Comment
}
