package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Title       string        `gorm:"type:varchar(255);not null"`
	Description string        `gorm:"type:text"`
	Position    sql.NullInt64 `gorm:"type:int"`
	DueDate     sql.NullTime
	ListId      uint
	List        List `gorm:"foreignKey:ListId"`
	Comments    []Comment
}
