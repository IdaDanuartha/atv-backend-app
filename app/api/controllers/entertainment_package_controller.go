package controllers

import (
	"net/http"

	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/utils"
	"github.com/gin-gonic/gin"
)

// EntertainmentPackageController -> EntertainmentPackageController
type EntertainmentPackageController struct {
	service services.EntertainmentPackageService
}

// NewEntertainmentPackageController : NewEntertainmentPackageController
func NewEntertainmentPackageController(s services.EntertainmentPackageService) EntertainmentPackageController {
	return EntertainmentPackageController{
		service: s,
	}
}

// GetEntertainmentPackages : GetEntertainmentPackages controller
func (p EntertainmentPackageController) GetEntertainmentPackages(ctx *gin.Context) {
	var entertainment_packages models.EntertainmentPackage

	search := ctx.Query("search")

	data, _, err := p.service.FindAll(entertainment_packages, search)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find entertainment package")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Entertainment package result set",
		Data:    respArr,
	})
}

// AddEntertainmentPackage : AddEntertainmentPackage controller
func (p *EntertainmentPackageController) AddEntertainmentPackage(ctx *gin.Context) {
	var entertainmentPackage models.EntertainmentPackage
	ctx.ShouldBindJSON(&entertainmentPackage)

	if entertainmentPackage.Name == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	err := p.service.Save(entertainmentPackage)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create entertainment package")
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Successfully created entertainment package")
}

// GetEntertainmentPackage : get entertainment package by id
func (p *EntertainmentPackageController) GetEntertainmentPackage(c *gin.Context) {
	idParam := c.Param("id")

	var bus models.EntertainmentPackage
	bus.ID = idParam
	foundBus, err := p.service.Find(bus)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Error finding entertainment package")
		return
	}
	response := foundBus.ResponseMap()

	c.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Result set of entertainment package",
		Data:    &response})

}

// UpdateEntertainmentPackage : get update by id
func (p EntertainmentPackageController) UpdateEntertainmentPackage(ctx *gin.Context) {
	idParam := ctx.Param("id")

	var entertainmentPackage models.EntertainmentPackage
	entertainmentPackage.ID = idParam

	entertainmentPackageRecord, err := p.service.Find(entertainmentPackage)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Entertainment package with given id not found")
		return
	}
	ctx.ShouldBindJSON(&entertainmentPackageRecord)

	if entertainmentPackageRecord.Name == "" {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.Update(entertainmentPackageRecord); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to update entertainment package")
		return
	}
	response := entertainmentPackageRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "Successfully updated entertainment package",
		Data:    response,
	})
}

// DeleteEntertainmentPackage : Deletes Entertainment Package
func (p *EntertainmentPackageController) DeleteEntertainmentPackage(c *gin.Context) {
	idParam := c.Param("id")

	err := p.service.Delete(idParam)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to delete entertainment package")
		return
	}
	response := &utils.Response{
		Success: true,
		Message: "Deleted sucessfully"}
	c.JSON(http.StatusOK, response)
}
