package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type RouteService interface {
	FindAll(model models.Route, search string, currentPage int, pageSize int) ([]models.Route, int64, int, error)
	Find(input inputs.GetRouteDetailInput) (models.Route, error)
	Save(input inputs.RouteInput) (models.Route, error)
	Update(inputID inputs.GetRouteDetailInput, input inputs.RouteInput) (models.Route, error)
	Delete(inputID inputs.GetRouteDetailInput) (models.Route, error)
}

// RouteService RouteService struct
type routeService struct {
	repository repositories.RouteRepository
}

// NewRouteService : returns the RouteService struct instance
func NewRouteService(repository repositories.RouteRepository) routeService {
	return routeService{repository}
}

// FindAll -> calls Route repo find all method
func (s routeService) FindAll(model models.Route, search string, currentPage int, pageSize int) ([]models.Route, int64, int, error) {
	routes, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return routes, total, currentPage, err
	}

	return routes, total, currentPage, nil
}

// Find -> calls Route repo find method
func (s routeService) Find(input inputs.GetRouteDetailInput) (models.Route, error) {
	route, err := s.repository.Find(input.ID)

	if err != nil {
		return route, err
	}

	return route, nil
}

// Save -> calls Route repository save method
func (s routeService) Save(input inputs.RouteInput) (models.Route, error) {
	route := models.Route{}
	route.Name = input.Name
	route.Address = input.Address

	newRoute, err := s.repository.Save(route)
	if err != nil {
		return newRoute, err
	}

	return newRoute, nil
}

// Update -> calls Route repo update method
func (s routeService) Update(inputID inputs.GetRouteDetailInput, input inputs.RouteInput) (models.Route, error) {
	route, err := s.repository.Find(inputID.ID)
	if err != nil {
		return route, err
	}

	route.Name = input.Name
	route.Address = input.Address

	updatedRoute, err := s.repository.Update(route)
	if err != nil {
		return updatedRoute, err
	}

	return updatedRoute, nil
}

// Delete -> calls Route repo delete method
func (s routeService) Delete(inputID inputs.GetRouteDetailInput) (models.Route, error) {
	route, err := s.repository.Find(inputID.ID)
	if err != nil {
		return route, err
	}

	deletedRoute, err := s.repository.Delete(route)
	if err != nil {
		return deletedRoute, err
	}

	return deletedRoute, nil
}
