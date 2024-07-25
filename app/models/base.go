package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base Models
type Base struct {
	// ID           string     `gorm:"type:uuid;primary_key;" json:"id"`
	ID        string         `gorm:"type:varchar(100);primaryKey;" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return
}

func (b *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now()
	return
}

// Delete overrides the Delete method to set DeletedAt before actual deletion
func (base *Base) Delete(tx *gorm.DB) error {
	base.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	return tx.Delete(base).Error
}

// ----------------------------------------------------------------

// BeforeCreate will set a UUID rather than numeric ID.
func (b *EntertainmentCategory) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *EntertainmentPackage) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *EntertainmentPackageDetail) BeforeCreate(tx *gorm.DB) (err error) {
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
func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	b.User.ID = uuid.New().String()
	b.ID = uuid.New().String()
	return
}
func (b *Staff) BeforeCreate(tx *gorm.DB) (err error) {
	b.User.ID = uuid.New().String()
	b.ID = uuid.New().String()
	return
}
func (b *Instructor) BeforeCreate(tx *gorm.DB) (err error) {
	b.User.ID = uuid.New().String()
	b.ID = uuid.New().String()
	return
}
func (b *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	b.User.ID = uuid.New().String()
	b.ID = uuid.New().String()
	return
}
func (b *Route) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *EntertainmentService) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *BookingDetail) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}