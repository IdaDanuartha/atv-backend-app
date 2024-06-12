package models

// Facility Model
type Facility struct {
	Base
	Name string `gorm:"size:100;uniqueIndex" json:"name"`
}

// TableName method sets table name for Bus model
func (Facility *Facility) TableName() string {
	return "facilities"
}
