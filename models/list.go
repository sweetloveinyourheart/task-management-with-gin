package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255);not null"`
	Board    Board  `gorm:"embedded"`
	Position *int   `gorm:"type:int"`
}
