package controllers

import (
	"net/http"

	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/utils"
	"github.com/gin-gonic/gin"
)

// MandatoryLuggageController -> MandatoryLuggageController
type MandatoryLuggageController struct {
	service services.MandatoryLuggageService
}

// NewMandatoryLuggageController : NewMandatoryLuggageController
func NewMandatoryLuggageController(s services.MandatoryLuggageService) MandatoryLuggageController {
	return MandatoryLuggageController{
		service: s,
	}
}

// GetMandatoryLuggages : GetMandatoryLuggages controller
func (p *MandatoryLuggageController) GetMandatoryLuggages(ctx *gin.Context) {
	var mandatory_luggages models.MandatoryLuggage

	search := ctx.Query("search")

	data, _, err := p.service.FindAll(mandatory_luggages, search)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find mandatory luggage")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Mandatory luggage result set",
		Data:    respArr,
	})
}

// AddMandatoryLuggage : AddMandatoryLuggage controller
func (p *MandatoryLuggageController) AddMandatoryLuggage(ctx *gin.Context) {
	var entertainmentCategory models.MandatoryLuggage
	ctx.ShouldBindJSON(&entertainmentCategory)

	if entertainmentCategory.Name == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	err := p.service.Save(entertainmentCategory)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create mandatory luggage")
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Successfully created mandatory luggage")
}

// GetMandatoryLuggage : get mandatory luggage by id
func (p *MandatoryLuggageController) GetMandatoryLuggage(c *gin.Context) {
	idParam := c.Param("id")

	var bus models.MandatoryLuggage
	bus.ID = idParam
	foundBus, err := p.service.Find(bus)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Error finding mandatory luggage")
		return
	}
	response := foundBus.ResponseMap()

	c.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Result set of mandatory luggage",
		Data:    &response})

}

// UpdateMandatoryLuggage : get update by id
func (p *MandatoryLuggageController) UpdateMandatoryLuggage(ctx *gin.Context) {
	idParam := ctx.Param("id")

	var entertainmentCategory models.MandatoryLuggage
	entertainmentCategory.ID = idParam

	entertainmentCategoryRecord, err := p.service.Find(entertainmentCategory)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Mandatory luggage with given id not found")
		return
	}
	ctx.ShouldBindJSON(&entertainmentCategoryRecord)

	if entertainmentCategoryRecord.Name == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.Update(entertainmentCategoryRecord); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to update mandatory luggage")
		return
	}
	response := entertainmentCategoryRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Successfully updated mandatory luggage",
		Data:    response,
	})
}

// DeleteMandatoryLuggage : Deletes Mandatory Luggage
func (p *MandatoryLuggageController) DeleteMandatoryLuggage(c *gin.Context) {
	idParam := c.Param("id")

	err := p.service.Delete(idParam)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to delete mandatory luggage")
		return
	}
	response := &utils.Response{
		Success: true,
		Message: "Deleted sucessfully"}
	c.JSON(http.StatusOK, response)
}
