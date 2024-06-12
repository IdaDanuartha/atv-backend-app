package models

// Entertainment Service Model
type EntertainmentService struct {
	Base
	Name      string  `gorm:"size:100;uniqueIndex" json:"name"`
	Price     int32   `json:"price"`
	ImagePath *string `gorm:"size:150;" json:"image_path"`

	EntertainmentCategoryID string                `gorm:"type:varchar(100);foreignKey:EntertainmentCategoryID" json:"entertainment_category_id,omitempty"`
	EntertainmentCategory   EntertainmentCategory `gorm:"foreignKey:EntertainmentCategoryID" json:"entertainment_category"`

	Routes            []EntertainmentServiceRoute            `gorm:"foreignKey:EntertainmentServiceID;" json:"routes"`
	Facilities        []EntertainmentServiceFacility         `gorm:"foreignKey:EntertainmentServiceID" json:"facilities"`
	Instructors       []EntertainmentServiceInstructor       `gorm:"foreignKey:EntertainmentServiceID" json:"instructors"`
	MandatoryLuggages []MandatoryLuggageEntertainmentService `gorm:"foreignKey:EntertainmentServiceID" json:"mandatory_luggages"`
}

// TableName method sets table name for Entertainment Service model
func (entertainmentService *EntertainmentService) TableName() string {
	return "entertainment_services"
}
