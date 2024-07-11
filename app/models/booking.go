package models

import "time"

// Entertainment Booking Model
type Booking struct {
	Base
	Customer       Customer        `gorm:"foreignKey:CustomerID" json:"customer"`
	CustomerID     string          `gorm:"type:varchar(100);foreignKey:CustomerID" json:"customer_id,omitempty"`
	Code           string          `gorm:"size:100" json:"code"`
	Name           string          `gorm:"size:100" json:"name"`
	PhoneNumber    string          `gorm:"size:15" json:"phone_number"`
	PaymentMethod  string          `gorm:"size:50" json:"payment_method"`
	Date           time.Time       `json:"date"`
	TotalPrice     int32           `json:"total_price"`
	TotalPay       int32           `json:"total_pay"`
	TotalChange    int32           `json:"total_change"`
	BookingDetails []BookingDetail `gorm:"foreignKey:BookingID" json:"details"`
}

// TableName method sets table name for Booking model
func (Booking *Booking) TableName() string {
	return "bookings"
}
