package models

// Entertainment Service Facility Model
type EntertainmentServiceFacility struct {
	EntertainmentServiceID string   `gorm:"type:varchar(100);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"entertainment_service_id,omitempty"`
	Facility               Facility `gorm:"foreignKey:FacilityID" json:"facility"`
	FacilityID             string   `gorm:"type:varchar(100);" json:"facility_id,omitempty"`
}

// TableName method sets table name for Entertainment Service Facility model
func (entertainmentServiceFacility *EntertainmentServiceFacility) TableName() string {
	return "entertainment_service_facilities"
}
