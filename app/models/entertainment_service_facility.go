package models

// Entertainment Service Facility Model
type EntertainmentServiceFacility struct {
	Base
	EntertainmentServiceID string   `gorm:"type:varchar(100);primaryKey;foreignKey:EntertainmentServiceID" json:"entertainment_service_id,omitempty"`
	Facility               Facility `gorm:"foreignKey:FacilityID" json:"facility"`
	FacilityID             string   `gorm:"type:varchar(100);primaryKey;foreignKey:FacilityID" json:"facility_id,omitempty"`
}

// TableName method sets table name for Entertainment Service Facility model
func (entertainmentServiceFacility *EntertainmentServiceFacility) TableName() string {
	return "entertainment_service_facilities"
}