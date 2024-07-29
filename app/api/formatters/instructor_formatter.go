package formatters

import (
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

func FormatInstructor(instructor models.Instructor) models.Instructor {
	instructorFormatter := models.Instructor{}
	instructorFormatter.ID = instructor.ID
	instructorFormatter.Name = instructor.Name
	instructorFormatter.EmployeeCode = instructor.EmployeeCode
	instructorFormatter.CreatedAt = instructor.CreatedAt
	instructorFormatter.UpdatedAt = instructor.UpdatedAt
	instructorFormatter.DeletedAt = instructor.DeletedAt

	instructorFormatter.User.ID = instructor.User.ID
	instructorFormatter.User.Username = instructor.User.Username
	instructorFormatter.User.Email = instructor.User.Email
	instructorFormatter.User.Password = instructor.User.Password
	instructorFormatter.User.Role = instructor.User.Role
	instructorFormatter.User.CreatedAt = instructor.User.CreatedAt
	instructorFormatter.User.UpdatedAt = instructor.User.UpdatedAt
	instructorFormatter.User.DeletedAt = instructor.User.DeletedAt

	return instructorFormatter
}

func FormatInstructors(instructors []models.Instructor) []models.Instructor {
	instructorsFormatter := []models.Instructor{}

	for _, instructor := range instructors {
		instructor := FormatInstructor(instructor)
		instructorsFormatter = append(instructorsFormatter, instructor)
	}

	return instructorsFormatter
}
