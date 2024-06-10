package inputs

import (
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"time"
)

type GetBookingDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type BookingInput struct {
	CustomerID    string                 `json:"customer_id" binding:"required"`
	Name          string                 `json:"name" binding:"required"`
	PhoneNumber   string                 `json:"phone_number" binding:"required"`
	PaymentMethod string                 `json:"payment_method" binding:"required"`
	Date          time.Time              `json:"date" binding:"required"`
	TotalPrice    int32                  `json:"total_price" binding:"required"`
	TotalPay      int32                  `json:"total_pay" binding:"required"`
	TotalChange   int32                  `json:"total_change" binding:"required"`
	Details       []models.BookingDetail `json:"details" binding:"required"`
}
