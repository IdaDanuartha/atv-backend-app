package formatters

import (
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

func FormatEntertainmentService(entertainmentService models.EntertainmentService) models.EntertainmentService {
	entertainmentServiceFormatter := models.EntertainmentService{}
	entertainmentServiceFormatter.ID = entertainmentService.ID
	entertainmentServiceFormatter.Name = entertainmentService.Name
	entertainmentServiceFormatter.Price = entertainmentService.Price
	entertainmentServiceFormatter.Duration = entertainmentService.Duration
	entertainmentServiceFormatter.Description = entertainmentService.Description
	entertainmentServiceFormatter.ImagePath = entertainmentService.ImagePath
	entertainmentServiceFormatter.CreatedAt = entertainmentService.CreatedAt
	entertainmentServiceFormatter.UpdatedAt = entertainmentService.UpdatedAt
	entertainmentServiceFormatter.DeletedAt = entertainmentService.DeletedAt

	entertainmentServiceFormatter.EntertainmentCategory.ID = entertainmentService.EntertainmentCategoryID
	entertainmentServiceFormatter.EntertainmentCategory.Name = entertainmentService.EntertainmentCategory.Name
	entertainmentServiceFormatter.EntertainmentCategory.CreatedAt = entertainmentService.EntertainmentCategory.CreatedAt
	entertainmentServiceFormatter.EntertainmentCategory.UpdatedAt = entertainmentService.EntertainmentCategory.UpdatedAt
	entertainmentServiceFormatter.EntertainmentCategory.DeletedAt = entertainmentService.EntertainmentCategory.DeletedAt

	routes := make([]models.EntertainmentServiceRoute, 0)

	for _, route := range entertainmentService.Routes {
		newRoute := models.EntertainmentServiceRoute{}

		newRoute.Route.ID = route.Route.ID
		newRoute.Route.Name = route.Route.Name
		newRoute.Route.Address = route.Route.Address
		newRoute.Route.CreatedAt = route.Route.CreatedAt
		newRoute.Route.UpdatedAt = route.Route.UpdatedAt
		newRoute.Route.DeletedAt = route.Route.DeletedAt

		routes = append(routes, newRoute)
	}

	entertainmentServiceFormatter.Routes = routes

	facilities := make([]models.EntertainmentServiceFacility, 0)

	for _, facility := range entertainmentService.Facilities {
		newFacility := models.EntertainmentServiceFacility{}

		newFacility.Facility.ID = facility.Facility.ID
		newFacility.Facility.Name = facility.Facility.Name
		newFacility.Facility.CreatedAt = facility.Facility.CreatedAt
		newFacility.Facility.UpdatedAt = facility.Facility.UpdatedAt
		newFacility.Facility.DeletedAt = facility.Facility.DeletedAt

		facilities = append(facilities, newFacility)
	}

	entertainmentServiceFormatter.Facilities = facilities

	instructors := make([]models.EntertainmentServiceInstructor, 0)

	for _, instructor := range entertainmentService.Instructors {
		newInstructor := models.EntertainmentServiceInstructor{}

		newInstructor.Instructor.ID = instructor.Instructor.ID
		newInstructor.Instructor.EmployeeCode = instructor.Instructor.EmployeeCode
		newInstructor.Instructor.Name = instructor.Instructor.Name
		newInstructor.Instructor.CreatedAt = instructor.Instructor.CreatedAt
		newInstructor.Instructor.UpdatedAt = instructor.Instructor.UpdatedAt
		newInstructor.Instructor.DeletedAt = instructor.Instructor.DeletedAt

		newInstructor.Instructor.User.ID = instructor.Instructor.User.ID
		newInstructor.Instructor.User.Username = instructor.Instructor.User.Username
		newInstructor.Instructor.User.Email = instructor.Instructor.User.Email
		newInstructor.Instructor.User.Role = instructor.Instructor.User.Role
		newInstructor.Instructor.User.ProfilePath = instructor.Instructor.User.ProfilePath
		newInstructor.Instructor.User.CreatedAt = instructor.Instructor.User.CreatedAt
		newInstructor.Instructor.User.UpdatedAt = instructor.Instructor.User.UpdatedAt
		newInstructor.Instructor.User.DeletedAt = instructor.Instructor.User.DeletedAt

		instructors = append(instructors, newInstructor)
	}

	entertainmentServiceFormatter.Instructors = instructors

	mandatoryLuggages := make([]models.MandatoryLuggageEntertainmentService, 0)

	for _, mandatoryLuggage := range entertainmentService.MandatoryLuggages {
		newMandatoryLuggage := models.MandatoryLuggageEntertainmentService{}

		newMandatoryLuggage.MandatoryLuggage.ID = mandatoryLuggage.MandatoryLuggage.ID
		newMandatoryLuggage.MandatoryLuggage.Name = mandatoryLuggage.MandatoryLuggage.Name
		newMandatoryLuggage.MandatoryLuggage.CreatedAt = mandatoryLuggage.MandatoryLuggage.CreatedAt
		newMandatoryLuggage.MandatoryLuggage.UpdatedAt = mandatoryLuggage.MandatoryLuggage.UpdatedAt
		newMandatoryLuggage.MandatoryLuggage.DeletedAt = mandatoryLuggage.MandatoryLuggage.DeletedAt

		mandatoryLuggages = append(mandatoryLuggages, newMandatoryLuggage)
	}

	entertainmentServiceFormatter.MandatoryLuggages = mandatoryLuggages

	return entertainmentServiceFormatter
}

func FormatEntertainmentServices(entertainmentServices []models.EntertainmentService) []models.EntertainmentService {
	entertainmentServiceFormatter := []models.EntertainmentService{}

	for _, entertainmentService := range entertainmentServices {
		entertainmentService := FormatEntertainmentService(entertainmentService)
		entertainmentServiceFormatter = append(entertainmentServiceFormatter, entertainmentService)
	}

	return entertainmentServiceFormatter
}
