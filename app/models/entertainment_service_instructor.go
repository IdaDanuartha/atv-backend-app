package models

// Entertainment Service Instructor Model
type EntertainmentServiceInstructor struct {
	EntertainmentServiceID string     `gorm:"type:varchar(100);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"entertainment_service_id,omitempty"`
	Instructor             Instructor `gorm:"foreignKey:InstructorID" json:"instructor"`
	InstructorID           string     `gorm:"type:varchar(100);" json:"instructor_id,omitempty"`
}

// TableName method sets table name for Entertainment Service Instructor model
func (entertainmentServiceInstructor *EntertainmentServiceInstructor) TableName() string {
	return "entertainment_service_instructors"
}
