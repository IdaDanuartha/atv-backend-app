package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatEntertainmentCategory(entertainmentCategory models.EntertainmentCategory) models.EntertainmentCategory {
	entertainmentCategoryFormatter := models.EntertainmentCategory{}
	entertainmentCategoryFormatter.ID = entertainmentCategory.ID
	entertainmentCategoryFormatter.Name = entertainmentCategory.Name
	entertainmentCategoryFormatter.CreatedAt = entertainmentCategory.CreatedAt
	entertainmentCategoryFormatter.UpdatedAt = entertainmentCategory.UpdatedAt
	entertainmentCategoryFormatter.DeletedAt = entertainmentCategory.DeletedAt

	return entertainmentCategoryFormatter
}

func FormatEntertainmentCategories(entertainmentCategories []models.EntertainmentCategory) []models.EntertainmentCategory {
	entertainmentCategoriesFormatter := []models.EntertainmentCategory{}

	for _, entertainmentCategory := range entertainmentCategories {
		entertainmentCategory := FormatEntertainmentCategory(entertainmentCategory)
		entertainmentCategoriesFormatter = append(entertainmentCategoriesFormatter, entertainmentCategory)
	}

	return entertainmentCategoriesFormatter
}