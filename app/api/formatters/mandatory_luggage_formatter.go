package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatMandatoryLuggage(mandatoryLuggage models.MandatoryLuggage) models.MandatoryLuggage {
	mandatoryLuggageFormatter := models.MandatoryLuggage{}
	mandatoryLuggageFormatter.ID = mandatoryLuggage.ID
	mandatoryLuggageFormatter.Name = mandatoryLuggage.Name
	mandatoryLuggageFormatter.CreatedAt = mandatoryLuggage.CreatedAt
	mandatoryLuggageFormatter.UpdatedAt = mandatoryLuggage.UpdatedAt
	mandatoryLuggageFormatter.DeletedAt = mandatoryLuggage.DeletedAt

	return mandatoryLuggageFormatter
}

func FormatMandatoryLuggages(mandatoryLuggages []models.MandatoryLuggage) []models.MandatoryLuggage {
	mandatoryLuggagesFormatter := []models.MandatoryLuggage{}

	for _, mandatoryLuggage := range mandatoryLuggages {
		mandatoryLuggage := FormatMandatoryLuggage(mandatoryLuggage)
		mandatoryLuggagesFormatter = append(mandatoryLuggagesFormatter, mandatoryLuggage)
	}

	return mandatoryLuggagesFormatter
}
