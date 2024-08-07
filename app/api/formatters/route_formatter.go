package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatRoute(route models.Route) models.Route {
	routeFormatter := models.Route{}
	routeFormatter.ID = route.ID
	routeFormatter.Name = route.Name
	routeFormatter.Address = route.Address
	routeFormatter.CreatedAt = route.CreatedAt
	routeFormatter.UpdatedAt = route.UpdatedAt
	routeFormatter.DeletedAt = route.DeletedAt

	return routeFormatter
}

func FormatRoutes(routes []models.Route) []models.Route {
	routesFormatter := []models.Route{}

	for _, route := range routes {
		route := FormatRoute(route)
		routesFormatter = append(routesFormatter, route)
	}

	return routesFormatter
}
