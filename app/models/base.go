package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

// Base Models
type Base struct {
	// ID           string     `gorm:"type:uuid;primary_key;" json:"id"`
	ID           string     `gorm:"primary_key;" json:"id"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (b *EntertainmentCategory) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *EntertainmentPackage) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *Facility) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *MandatoryLuggage) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}