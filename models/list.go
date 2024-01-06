package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255);not null"`
	BoardId  uint
	Board    Board `gorm:"foreignKey:BoardId"`
	Position *int  `gorm:"type:int"`
}
