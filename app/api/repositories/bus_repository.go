package repositories

import (
	"github.com/IdaDanuartha/teman-bus-backend-app/app/config"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/models"
)

//BusRepository -> BusRepository
type BusRepository struct {
    db config.Database
}

// NewBusRepository : fetching database
func NewBusRepository(db config.Database) BusRepository {
    return BusRepository{
        db: db,
    }
}

//Save -> Method for saving bus to database
func (p BusRepository) Save(bus models.Bus) error {
    return p.db.DB.Create(&bus).Error
}

//FindAll -> Method for fetching all buses from database
func (p BusRepository) FindAll(bus models.Bus, keyword string) (*[]models.Bus, int64, error) {
    var buses []models.Bus
    var totalRows int64 = 0

    queryBuider := p.db.DB.Order("created_at desc").Model(&models.Bus{})

    // Search parameter
    if keyword != "" {
        queryKeyword := "%" + keyword + "%"
        queryBuider = queryBuider.Where(
            p.db.DB.Where("bus.title LIKE ? ", queryKeyword))
    }

    err := queryBuider.
        Where(bus).
        Find(&buses).
        Count(&totalRows).Error
    return &buses, totalRows, err
}

//Update -> Method for updating bus
func (p BusRepository) Update(bus models.Bus) error {
    return p.db.DB.Save(&bus).Error
}

//Find -> Method for fetching bus by id
func (p BusRepository) Find(bus models.Bus) (models.Bus, error) {
    var buses models.Bus
    err := p.db.DB.
        Debug().
        Model(&models.Bus{}).
        Where(&bus).
        Take(&buses).Error
    return buses, err
}

//Delete Deletes bus
func (p BusRepository) Delete(bus models.Bus) error {
    return p.db.DB.Delete(&bus).Error
}