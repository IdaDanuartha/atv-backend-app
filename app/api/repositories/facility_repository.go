package repositories

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/gin-gonic/gin"
)

type FacilityRepository interface {
	FindAll(facility models.Facility, search string, currentPage int, pageSize int) ([]models.Facility, int64, int, error)
	ExportToExcel(facility models.Facility, ctx *gin.Context) error
	Find(ID string) (models.Facility, error)
	Save(facility models.Facility) (models.Facility, error)
	Update(facility models.Facility) (models.Facility, error)
	Delete(facility models.Facility) (models.Facility, error)
}

type facilityRepository struct {
	db config.Database
}

// NewFacilityRepository : fetching database
func NewFacilityRepository(db config.Database) facilityRepository {
	return facilityRepository{db}
}

// FindAll -> Method for fetching all Facility from database
func (r facilityRepository) FindAll(facility models.Facility, search string, currentPage int, pageSize int) ([]models.Facility, int64, int, error) {
	var facilities []models.Facility
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.Facility{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("facilities.name LIKE ? ", querySearch))
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(facility).
			Count(&totalRows).Error

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Where(facility).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&facilities).Error

		return facilities, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Where(facility).
			Find(&facilities).
			Count(&totalRows).Error

		return facilities, 0, 0, err
	}
}

// ExportToExcel -> Method for exporting all data to excel file
func (r facilityRepository) ExportToExcel(facility models.Facility, ctx *gin.Context) error {
	var facilities []models.Facility

	// Retrieve data from the database
	rows := r.db.DB.Find(&facilities)
	if rows.Error != nil {
		return rows.Error
	}

	// Create an Excel file
	xlsx := excelize.NewFile()
	sheetName := "Facilities"
	xlsx.SetSheetName("Facilities", sheetName)

	// Add headers
	xlsx.SetCellValue(sheetName, "A1", "No")
	xlsx.SetCellValue(sheetName, "B1", "Name")

	// Fill data from the database into an Excel file
	rowIndex := 1
	for _, facility := range facilities {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), rowIndex)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), facility.Name)
		rowIndex++
	}

	// Save the file to a local path for debugging (optional)
	err := xlsx.SaveAs("facilities_export.xlsx")
	if err != nil {
		fmt.Println("Error saving file locally:", err)
		return err
	}

	// Set header response
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename=facilities_export.xlsx")
	ctx.Header("Content-Type", "application/octet-stream")

	// Write the Excel file to the response
	err = xlsx.Write(ctx.Writer)
	if err != nil {
		return err
	}

	return nil
}

// Find -> Method for fetching Facility by id
func (r facilityRepository) Find(ID string) (models.Facility, error) {
	var facilities models.Facility
	err := r.db.DB.
		Debug().
		Model(&models.Facility{}).
		Where("id = ?", ID).
		Find(&facilities).Error
	return facilities, err
}

// Save -> Method for saving Facility to database
func (r facilityRepository) Save(facility models.Facility) (models.Facility, error) {
	err := r.db.DB.Create(&facility).Error
	if err != nil {
		return facility, err
	}

	return facility, nil
}

// Update -> Method for updating Facility
func (r *facilityRepository) Update(facility models.Facility) (models.Facility, error) {
	err := r.db.DB.Save(&facility).Error

	if err != nil {
		return facility, err
	}

	return facility, nil
}

// Delete -> Method for deleting Facility
func (r facilityRepository) Delete(facility models.Facility) (models.Facility, error) {
	err := r.db.DB.Delete(&facility).Error

	if err != nil {
		return facility, err
	}

	return facility, nil
}
