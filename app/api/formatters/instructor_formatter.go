package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatInstructor(intructor models.Instructor) models.Instructor {
	intructorFormatter := models.Instructor{}
	intructorFormatter.ID = intructor.ID
	intructorFormatter.Name = intructor.Name
	intructorFormatter.EmployeeCode = intructor.EmployeeCode
	intructorFormatter.UserID = intructor.User.ID
	intructorFormatter.CreatedAt = intructor.CreatedAt
	intructorFormatter.UpdatedAt = intructor.UpdatedAt
	intructorFormatter.DeletedAt = intructor.DeletedAt

	intructorFormatter.User.ID = intructor.User.ID
	intructorFormatter.User.Username = intructor.User.Username
	intructorFormatter.User.Email = intructor.User.Email
	intructorFormatter.User.Password = intructor.User.Password
	intructorFormatter.User.Role = intructor.User.Role
	intructorFormatter.User.CreatedAt = intructor.User.CreatedAt
	intructorFormatter.User.UpdatedAt = intructor.User.UpdatedAt
	intructorFormatter.User.DeletedAt = intructor.User.DeletedAt

	return intructorFormatter
}

func FormatInstructors(intructors []models.Instructor) []models.Instructor {
	intructorsFormatter := []models.Instructor{}

	for _, intructor := range intructors {
		intructor := FormatInstructor(intructor)
		intructorsFormatter = append(intructorsFormatter, intructor)
	}

	return intructorsFormatter
}
