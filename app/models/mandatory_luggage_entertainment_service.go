package models

// Mandatory Luggage Entertainment Service Model
type MandatoryLuggageEntertainmentService struct {
	EntertainmentServiceID string           `gorm:"type:varchar(100);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"entertainment_service_id,omitempty"`
	MandatoryLuggage       MandatoryLuggage `gorm:"foreignKey:MandatoryLuggageID" json:"mandatory_luggage"`
	MandatoryLuggageID     string           `gorm:"type:varchar(100);" json:"mandatory_luggage_id,omitempty"`
}

// TableName method sets table name for Mandatory Luggage Entertainment Service model
func (mandatoryLuggagEentertainmentService *MandatoryLuggageEntertainmentService) TableName() string {
	return "mandatory_luggage_entertainment_services"
}
