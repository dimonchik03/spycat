package models

import (
	"gorm.io/gorm"
)

type Target struct {
	gorm.Model
	MissionID uint    `json:"mission_id" validate:"required" gorm:"not null"`
	Mission   Mission `json:"mission" gorm:"foreignKey:MissionID"`
	Name      string  `json:"name" validate:"required"`
	Country   string  `json:"country" validate:"required"`
	Notes     string  `json:"notes" validate:"required"`
	Complete  bool    `json:"complete" validate:"required"`
}
