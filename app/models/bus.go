package models

import (
	"time"
	"gorm.io/gorm"
)

// Bus Model
type Bus struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name         string    `gorm:"size:100" json:"name"`
	LicensePlate string    `gorm:"size:20" json:"license_plate" `
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName method sets table name for Bus model
func (bus *Bus) TableName() string {
	return "bus"
}

//ResponseMap -> response map method of Bus
func (Bus *Bus) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = Bus.ID
	resp["name"] = Bus.Name
	resp["license_plate"] = Bus.LicensePlate
	resp["created_at"] = Bus.CreatedAt
	resp["updated_at"] = Bus.UpdatedAt
	resp["deleted_at"] = Bus.DeletedAt

	return resp
}
