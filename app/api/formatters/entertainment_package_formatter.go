package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatEntertainmentPackage(entertainmentPackage models.EntertainmentPackage) models.EntertainmentPackage {
	entertainmentPackageFormatter := models.EntertainmentPackage{}
	entertainmentPackageFormatter.ID = entertainmentPackage.ID
	entertainmentPackageFormatter.Name = entertainmentPackage.Name
	entertainmentPackageFormatter.CreatedAt = entertainmentPackage.CreatedAt
	entertainmentPackageFormatter.UpdatedAt = entertainmentPackage.UpdatedAt
	entertainmentPackageFormatter.DeletedAt = entertainmentPackage.DeletedAt

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
