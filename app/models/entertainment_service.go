package models

// Entertainment Service Model
type EntertainmentService struct {
	Base
	Name        string    `gorm:"size:100" json:"name"`
	Price       int32     `json:"price"`
	ImagePath *string    `gorm:"size:100;" json:"image_path"`

	EntertainmentCategoryID string `gorm:"type:varchar(100);primaryKey;foreignKey:EntertainmentCategoryID" json:"entertainment_category_id,omitempty"`
	EntertainmentCategory EntertainmentCategory `gorm:"foreignKey:EntertainmentCategoryID" json:"entertainment_category"`
	
	RouteID string `gorm:"type:varchar(100);primaryKey;foreignKey:RouteID" json:"route_id,omitempty"`
	Route Route `gorm:"foreignKey:RouteID" json:"route"`

	EntertainmentServiceFacilities []EntertainmentServiceFacility `gorm:"foreignKey:EntertainmentServiceID" json:"facilities"`
	EntertainmentServiceInstructors []EntertainmentServiceInstructor `gorm:"foreignKey:EntertainmentServiceID" json:"instructors"`
	MandatoryLuggageEntertainmentServices []MandatoryLuggageEntertainmentService `gorm:"foreignKey:EntertainmentServiceID" json:"entertainment_services"`
}

// TableName method sets table name for Entertainment Service model
func (entertainmentService *EntertainmentService) TableName() string {
	return "entertainment_services"
}