package models

import "time"

// Entertainment Package Model
type EntertainmentPackage struct {
	Base
	Name                        string                       `gorm:"size:100" json:"name"`
	Description                 string                       `gorm:"type:text" json:"description"`
	Price                       int32                        `json:"price"`
	Duration                    int32                        `json:"duration"`
	ExpiredAt                   time.Time                    `json:"expired_at"`
	ImagePath                   *string                      `gorm:"size:150;" json:"image_path"`
	EntertainmentPackageDetails []EntertainmentPackageDetail `gorm:"foreignKey:EntertainmentPackageID" json:"services"`
}

// TableName method sets table name for Bus model
func (entertainmentPackage *EntertainmentPackage) TableName() string {
	return "entertainment_packages"
}
