package inputs

import "time"

type GetEntertainmentPackageDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type EntertainmentPackageInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int32 `json:"price" binding:"required"`
	ExpiredAt   time.Time `json:"expired_at" binding:"required"`
}
