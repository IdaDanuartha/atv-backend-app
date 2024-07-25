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
	
	services := []models.EntertainmentServiceInstructor{}

	for _, service := range instructor.Services {
		newService := models.EntertainmentServiceInstructor{}
		
		newService.EntertainmentService.ID = service.EntertainmentService.ID
		newService.EntertainmentService.Name = service.EntertainmentService.Name
		newService.EntertainmentService.Price = service.EntertainmentService.Price
		newService.EntertainmentService.Duration = service.EntertainmentService.Duration
		newService.EntertainmentService.ImagePath = service.EntertainmentService.ImagePath
		newService.EntertainmentService.Description = service.EntertainmentService.Description
		newService.EntertainmentService.CreatedAt = service.EntertainmentService.CreatedAt
		newService.EntertainmentService.UpdatedAt = service.EntertainmentService.UpdatedAt
		newService.EntertainmentService.DeletedAt = service.EntertainmentService.DeletedAt

		newService.EntertainmentService.EntertainmentCategory.ID = service.EntertainmentService.EntertainmentCategory.ID
		newService.EntertainmentService.EntertainmentCategory.Name = service.EntertainmentService.EntertainmentCategory.Name
		newService.EntertainmentService.EntertainmentCategory.CreatedAt = service.EntertainmentService.EntertainmentCategory.CreatedAt
		newService.EntertainmentService.EntertainmentCategory.UpdatedAt = service.EntertainmentService.EntertainmentCategory.UpdatedAt
		newService.EntertainmentService.EntertainmentCategory.DeletedAt = service.EntertainmentService.EntertainmentCategory.DeletedAt

		services = append(services, newService)
	}

	instructorFormatter.Services = services

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
