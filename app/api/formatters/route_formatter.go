package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatRoute(route models.Route) models.Route {
	routeFormatter := models.Route{}
	routeFormatter.ID = route.ID
	routeFormatter.StartingRoute = route.StartingRoute
	routeFormatter.FinalRoute = route.FinalRoute
	routeFormatter.Duration = route.Duration
	routeFormatter.Distance = route.Distance

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
