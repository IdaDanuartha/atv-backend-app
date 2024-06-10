package models

// Entertainment Service Facility Model
type EntertainmentServiceFacility struct {
	EntertainmentServiceID string   `gorm:"type:varchar(100);foreignKey:EntertainmentServiceID" json:"entertainment_service_id,omitempty"`
	Facility               Facility `gorm:"foreignKey:FacilityID" json:"facility"`
	FacilityID             string   `gorm:"type:varchar(100);foreignKey:FacilityID" json:"facility_id,omitempty"`
}

// TableName method sets table name for Entertainment Service Facility model
func (entertainmentServiceFacility *EntertainmentServiceFacility) TableName() string {
	return "entertainment_service_facilities"
}
