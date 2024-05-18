package services

import (
	"github.com/IdaDanuartha/teman-bus-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/models"
)

//BusService BusService struct
type BusService struct {
    repository repositories.BusRepository
}

//NewBusService : returns the BusService struct instance
func NewBusService(r repositories.BusRepository) BusService {
    return BusService{
        repository: r,
    }
}

//Save -> calls bus repository save method
func (p BusService) Save(bus models.Bus) error {
    return p.repository.Save(bus)
}

//FindAll -> calls bus repo find all method
func (p BusService) FindAll(bus models.Bus, keyword string) (*[]models.Bus, int64, error) {
    return p.repository.FindAll(bus, keyword)
}

// Update -> calls busrepo update method
func (p BusService) Update(bus models.Bus) error {
    return p.repository.Update(bus)
}

// Delete -> calls bus repo delete method
func (p BusService) Delete(id int64) error {
    var bus models.Bus
    bus.ID = id
    return p.repository.Delete(bus)
}

// Find -> calls bus repo find method
func (p BusService) Find(bus models.Bus) (models.Bus, error) {
    return p.repository.Find(bus)
}