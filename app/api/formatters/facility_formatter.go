package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatFacility(facility models.Facility) models.Facility {
	facilityFormatter := models.Facility{}
	facilityFormatter.ID = facility.ID
	facilityFormatter.Name = facility.Name

	return facilityFormatter
}

func FormatFacilities(facilities []models.Facility) []models.Facility {
	facilitiesFormatter := []models.Facility{}

	for _, facility := range facilities {
		facility := FormatFacility(facility)
		facilitiesFormatter = append(facilitiesFormatter, facility)
	}

	return facilitiesFormatter
}
