package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type RouteRepository interface {
	FindAll(route models.Route, search string) ([]models.Route, int64, error)
	Find(ID string) (models.Route, error)
	Save(route models.Route) (models.Route, error)
	Update(route models.Route) (models.Route, error)
	Delete(route models.Route) (models.Route, error)
}

type routeRepository struct {
	db config.Database
}

// NewRouteRepository : fetching database
func NewRouteRepository(db config.Database) routeRepository {
	return routeRepository{db}
}

// FindAll -> Method for fetching all route from database
func (r routeRepository) FindAll(route models.Route, search string) ([]models.Route, int64, error) {
	var routes []models.Route
	var totalRows int64 = 0

	queryBuider := r.db.DB.Order("created_at desc").Model(&models.Route{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			r.db.DB.Where("routes.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Where(route).
		Find(&routes).
		Count(&totalRows).Error
	return routes, totalRows, err
}

// Find -> Method for fetching route by id
func (r routeRepository) Find(ID string) (models.Route, error) {
	var routes models.Route
	err := r.db.DB.
		Debug().
		Model(&models.Route{}).
		Where("id = ?", ID).
		Find(&routes).Error
	return routes, err
}

// Save -> Method for saving route to database
func (r routeRepository) Save(route models.Route) (models.Route, error) {
	err := r.db.DB.Create(&route).Error
	if err != nil {
		return route, err
	}

	return route, nil
}

// Update -> Method for updating route
func (r *routeRepository) Update(route models.Route) (models.Route, error) {
	err := r.db.DB.Save(&route).Error

	if err != nil {
		return route, err
	}

	return route, nil
}

// Delete -> Method for deleting route
func (r routeRepository) Delete(route models.Route) (models.Route, error) {
	err := r.db.DB.Delete(&route).Error

	if err != nil {
		return route, err
	}

	return route, nil
}
