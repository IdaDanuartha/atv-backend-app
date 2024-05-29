package models

import "time"

// Entertainment Package Model
type EntertainmentPackage struct {
	Base
	Name                       string                       `gorm:"size:100" json:"name"`
	Description                string                       `gorm:"type:text" json:"description"`
	Price                      int32                        `json:"price"`
	ExpiredAt                  time.Time                    `json:"expired_at"`
	EntertainmentPackageDetails []EntertainmentPackageDetail `gorm:"foreignKey:EntertainmentPackageID" json:"facilities"`
}

// TableName method sets table name for Bus model
func (entertainmentPackage *EntertainmentPackage) TableName() string {
	return "entertainment_packages"
}
