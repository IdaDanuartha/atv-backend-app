package controllers

import (
	"net/http"

	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/utils"
	"github.com/gin-gonic/gin"
)

// FacilityController -> FacilityController
type FacilityController struct {
	service services.FacilityService
}

// NewFacilityController : NewFacilityController
func NewFacilityController(s services.FacilityService) FacilityController {
	return FacilityController{
		service: s,
	}
}

// GetFacilities : GetFacilities controller
func (p FacilityController) GetFacilities(ctx *gin.Context) {
	var facilities models.Facility

	search := ctx.Query("search")

	data, _, err := p.service.FindAll(facilities, search)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find facility")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Facility result set",
		Data:    respArr,
	})
}

// AddFacility : AddFacility controller
func (p *FacilityController) AddFacility(ctx *gin.Context) {
	var entertainmentCategory models.Facility
	ctx.ShouldBindJSON(&entertainmentCategory)

	if entertainmentCategory.Name == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	err := p.service.Save(entertainmentCategory)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create facility")
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Successfully created facility")
}

// GetFacility : get facility by id
func (p *FacilityController) GetFacility(c *gin.Context) {
	idParam := c.Param("id")

	var bus models.Facility
	bus.ID = idParam
	foundBus, err := p.service.Find(bus)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Error finding facility")
		return
	}
	response := foundBus.ResponseMap()

	c.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Result set of facility",
		Data:    &response})

}

// UpdateFacility : get update by id
func (p FacilityController) UpdateFacility(ctx *gin.Context) {
	idParam := ctx.Param("id")

	var entertainmentCategory models.Facility
	entertainmentCategory.ID = idParam

	entertainmentCategoryRecord, err := p.service.Find(entertainmentCategory)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "facility with given id not found")
		return
	}
	ctx.ShouldBindJSON(&entertainmentCategoryRecord)

	if entertainmentCategoryRecord.Name == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.Update(entertainmentCategoryRecord); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to update facility")
		return
	}
	response := entertainmentCategoryRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Successfully updated facility",
		Data:    response,
	})
}

// DeleteFacility : Deletes Facility
func (p *FacilityController) DeleteFacility(c *gin.Context) {
	idParam := c.Param("id")

	err := p.service.Delete(idParam)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to delete facility")
		return
	}
	response := &utils.Response{
		Success: true,
		Message: "Deleted sucessfully"}
	c.JSON(http.StatusOK, response)
}
