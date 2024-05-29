package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatEntertainmentService(entertainmentService models.EntertainmentService) models.EntertainmentService {
	entertainmentServiceFormatter := models.EntertainmentService{}
	entertainmentServiceFormatter.ID = entertainmentService.ID
	entertainmentServiceFormatter.Name = entertainmentService.Name
	entertainmentServiceFormatter.Price = entertainmentService.Price
	entertainmentServiceFormatter.CreatedAt = entertainmentService.CreatedAt
	entertainmentServiceFormatter.UpdatedAt = entertainmentService.UpdatedAt
	entertainmentServiceFormatter.DeletedAt = entertainmentService.DeletedAt

	entertainmentServiceFormatter.EntertainmentCategory.ID = entertainmentService.EntertainmentCategoryID
	entertainmentServiceFormatter.EntertainmentCategory.Name = entertainmentService.EntertainmentCategory.Name
	entertainmentServiceFormatter.EntertainmentCategory.CreatedAt = entertainmentService.EntertainmentCategory.CreatedAt
	entertainmentServiceFormatter.EntertainmentCategory.UpdatedAt = entertainmentService.EntertainmentCategory.UpdatedAt
	entertainmentServiceFormatter.EntertainmentCategory.DeletedAt = entertainmentService.EntertainmentCategory.DeletedAt

	entertainmentServiceFormatter.Route.ID = entertainmentService.RouteID
	entertainmentServiceFormatter.Route.StartingRoute = entertainmentService.Route.StartingRoute
	entertainmentServiceFormatter.Route.FinalRoute = entertainmentService.Route.FinalRoute
	entertainmentServiceFormatter.Route.Duration = entertainmentService.Route.Duration
	entertainmentServiceFormatter.Route.Distance = entertainmentService.Route.Distance
	entertainmentServiceFormatter.Route.CreatedAt = entertainmentService.Route.CreatedAt
	entertainmentServiceFormatter.Route.UpdatedAt = entertainmentService.Route.UpdatedAt
	entertainmentServiceFormatter.Route.DeletedAt = entertainmentService.Route.DeletedAt

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
