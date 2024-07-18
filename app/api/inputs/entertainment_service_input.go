package inputs

import "github.com/IdaDanuartha/atv-backend-app/app/models"

type GetEntertainmentServiceDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type EntertainmentServiceInput struct {
	Name                    string                                        `json:"name" binding:"required"`
	Price                   int32                                         `json:"price" binding:"required"`
	EntertainmentCategoryID string                                        `json:"entertainment_category_id" binding:"required"`
	Routes                  []models.EntertainmentServiceRoute            `json:"routes" binding:"required"`
	Facilities              []models.EntertainmentServiceFacility         `json:"facilities" binding:"required"`
	Instructors             []models.EntertainmentServiceInstructor       `json:"instructors" binding:"required"`
	MandatoryLuggages       []models.MandatoryLuggageEntertainmentService `json:"mandatory_luggages" binding:"required"`
}
