package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string  `gorm:"type:varchar(100);not null;unique;index:idx_name"`
	Username string  `gorm:"type:varchar(50);not null;unique;index:idx_name"`
	Password string  `gorm:"type:varchar(100);not null"`
	FullName *string `gorm:"type:varchar(100)"`
}
