package formatters

import (
	"fmt"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

func FormatEntertainmentPackage(entertainmentPackage models.EntertainmentPackage) models.EntertainmentPackage {
	entertainmentPackageFormatter := models.EntertainmentPackage{}
	entertainmentPackageFormatter.ID = entertainmentPackage.ID
	entertainmentPackageFormatter.Name = entertainmentPackage.Name
	entertainmentPackageFormatter.Description = entertainmentPackage.Description
	entertainmentPackageFormatter.Price = entertainmentPackage.Price
	entertainmentPackageFormatter.Duration = entertainmentPackage.Duration
	entertainmentPackageFormatter.ImagePath = entertainmentPackage.ImagePath
	entertainmentPackageFormatter.ExpiredAt = entertainmentPackage.ExpiredAt
	entertainmentPackageFormatter.CreatedAt = entertainmentPackage.CreatedAt
	entertainmentPackageFormatter.UpdatedAt = entertainmentPackage.UpdatedAt
	entertainmentPackageFormatter.DeletedAt = entertainmentPackage.DeletedAt

	packageDetails := make([]models.EntertainmentPackageDetail, 0)

	for _, packageDetail := range entertainmentPackage.Services {
		fmt.Println(packageDetail.EntertainmentService.Instructors)
		newPackageDetail := models.EntertainmentPackageDetail{}

		newPackageDetail.ID = packageDetail.ID
		newPackageDetail.CreatedAt = packageDetail.CreatedAt
		newPackageDetail.UpdatedAt = packageDetail.UpdatedAt
		newPackageDetail.DeletedAt = packageDetail.DeletedAt

		newPackageDetail.EntertainmentService.ID = packageDetail.EntertainmentService.ID
		newPackageDetail.EntertainmentService.Name = packageDetail.EntertainmentService.Name
		newPackageDetail.EntertainmentService.Price = packageDetail.EntertainmentService.Price
		newPackageDetail.EntertainmentService.Duration = packageDetail.EntertainmentService.Duration
		newPackageDetail.EntertainmentService.Description = packageDetail.EntertainmentService.Description
		newPackageDetail.EntertainmentService.ImagePath = packageDetail.EntertainmentService.ImagePath
		newPackageDetail.EntertainmentService.CreatedAt = packageDetail.EntertainmentService.CreatedAt
		newPackageDetail.EntertainmentService.UpdatedAt = packageDetail.EntertainmentService.UpdatedAt
		newPackageDetail.EntertainmentService.DeletedAt = packageDetail.EntertainmentService.DeletedAt

		newPackageDetail.EntertainmentService.EntertainmentCategory.ID = packageDetail.EntertainmentService.EntertainmentCategory.ID
		newPackageDetail.EntertainmentService.EntertainmentCategory.Name = packageDetail.EntertainmentService.EntertainmentCategory.Name
		newPackageDetail.EntertainmentService.EntertainmentCategory.CreatedAt = packageDetail.EntertainmentService.EntertainmentCategory.CreatedAt
		newPackageDetail.EntertainmentService.EntertainmentCategory.UpdatedAt = packageDetail.EntertainmentService.EntertainmentCategory.UpdatedAt
		newPackageDetail.EntertainmentService.EntertainmentCategory.DeletedAt = packageDetail.EntertainmentService.EntertainmentCategory.DeletedAt

		packageDetails = append(packageDetails, newPackageDetail)

		// entertainmentInstructors := []models.EntertainmentServiceInstructor{}
		// for _, instructor := range packageDetail.EntertainmentService.Instructors {
		// 	newInstructors := models.EntertainmentServiceInstructor{}

		// 	newInstructors.Instructor.Name = instructor.Instructor.Name
		// 	newInstructors.Instructor.CreatedAt = instructor.Instructor.CreatedAt
		// 	newInstructors.Instructor.UpdatedAt = instructor.Instructor.UpdatedAt
		// 	newInstructors.Instructor.DeletedAt = instructor.Instructor.DeletedAt

		// 	entertainmentInstructors = append(entertainmentInstructors, newInstructors)
		// }
		// newPackageDetail.EntertainmentService.Instructors = entertainmentInstructors

	}
	entertainmentPackageFormatter.Services = packageDetails

	return entertainmentPackageFormatter
}

func FormatEntertainmentPackages(entertainmentPackages []models.EntertainmentPackage) []models.EntertainmentPackage {
	entertainmentPackagesFormatter := []models.EntertainmentPackage{}

	for _, entertainmentPackage := range entertainmentPackages {
		entertainmentPackage := FormatEntertainmentPackage(entertainmentPackage)
		entertainmentPackagesFormatter = append(entertainmentPackagesFormatter, entertainmentPackage)
	}

	return entertainmentPackagesFormatter
}

func FormatEntertainmentPackageDetail(entertainmentPackageDetail models.EntertainmentPackageDetail) models.EntertainmentPackageDetail {
	entertainmentPackageFormatterDetail := models.EntertainmentPackageDetail{}
	entertainmentPackageFormatterDetail.ID = entertainmentPackageDetail.ID
	entertainmentPackageFormatterDetail.EntertainmentPackageID = entertainmentPackageDetail.EntertainmentPackageID
	entertainmentPackageFormatterDetail.EntertainmentServiceID = entertainmentPackageDetail.EntertainmentServiceID
	entertainmentPackageFormatterDetail.CreatedAt = entertainmentPackageDetail.CreatedAt
	entertainmentPackageFormatterDetail.UpdatedAt = entertainmentPackageDetail.UpdatedAt
	entertainmentPackageFormatterDetail.DeletedAt = entertainmentPackageDetail.DeletedAt

	return entertainmentPackageFormatterDetail
}
