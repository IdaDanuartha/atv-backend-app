package models

// Mandatory Luggage Model
type MandatoryLuggage struct {
	Base
	Name string `gorm:"size:100;uniqueIndex" json:"name"`
}

// TableName method sets table name for Bus model
func (MandatoryLuggage *MandatoryLuggage) TableName() string {
	return "mandatory_luggages"
}
