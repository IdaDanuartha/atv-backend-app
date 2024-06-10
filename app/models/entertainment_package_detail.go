package models

// Entertainment Package Detail Facility Model
type EntertainmentPackageDetail struct {
	Base
	EntertainmentPackageID string               `gorm:"type:varchar(100);foreignKey:EntertainmentPackageID" json:"entertainment_package_id,omitempty"`
	EntertainmentService   EntertainmentService `gorm:"foreignKey:EntertainmentServiceID" json:"entertainment_service"`
	EntertainmentServiceID string               `gorm:"type:varchar(100);foreignKey:EntertainmentServiceID" json:"entertainment_service_id,omitempty"`
}

// TableName method sets table name for Entertainment Package Detail Facility model
func (entertainmentPackEntertainmentPackageDetail *EntertainmentPackageDetail) TableName() string {
	return "entertainment_package_details"
}
