package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255);not null"`
	BoardId  uint
	Board    Board         `gorm:"foreignKey:BoardId"`
	Position sql.NullInt64 `gorm:"type:int"`
}
