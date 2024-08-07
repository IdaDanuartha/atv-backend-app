package inputs

import (
	"time"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type GetEntertainmentPackageDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type EntertainmentPackageInput struct {
	Name        string                              `json:"name" binding:"required"`
	Description string                              `json:"description" binding:"required"`
	Price       int32                               `json:"price" binding:"required"`
	Duration    int32                               `json:"duration" binding:"required"`
	ExpiredAt   time.Time                           `json:"expired_at" binding:"required"`
	Services    []models.EntertainmentPackageDetail `json:"services" binding:"required"`
}
