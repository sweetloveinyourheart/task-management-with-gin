package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string         `gorm:"type:varchar(100);not null;unique;index:idx_name"`
	Username string         `gorm:"type:varchar(50);not null;unique;index:idx_name"`
	Password string         `gorm:"type:varchar(100);not null"`
	FullName sql.NullString `gorm:"type:varchar(100)"`
	Boards   []Board
	Comments []Comment
}
