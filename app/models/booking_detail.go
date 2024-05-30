package models

// Booking Detail Model
type BookingDetail struct {
	Base
	BookingID              string               `gorm:"type:varchar(100);primaryKey;foreignKey:BookingID" json:"booking_id,omitempty"`
	EntertainmentService   EntertainmentService `gorm:"foreignKey:EntertainmentServiceID" json:"entertainment_service"`
	EntertainmentServiceID string               `gorm:"type:varchar(100);primaryKey;foreignKey:EntertainmentServiceID" json:"entertainment_service_id,omitempty"`
	Price                  int32                `json:"price"`
	Qty                    int16                `json:"qty"`
}

// TableName method sets table name for Booking Detail model
func (entertainmentPackBookingDetail *BookingDetail) TableName() string {
	return "booking_details"
}
