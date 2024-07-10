package models

import (
	"gorm.io/gorm"
)

type Mission struct {
	gorm.Model
	CatID    uint     `json:"cat_id"`
	Cat      Cat      `json:"cat" gorm:"foreignKey:CatID"`
	Targets  []Target `json:"targets"`
	Complete bool     `json:"complete"`
}
