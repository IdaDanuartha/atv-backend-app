package controllers

import (
	"net/http"

	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/utils"
	"github.com/gin-gonic/gin"
)

//EntertainmentCategoryController -> EntertainmentCategoryController
type EntertainmentCategoryController struct {
    service services.EntertainmentCategoryService
}

//NewEntertainmentCategoryController : NewEntertainmentCategoryController
func NewEntertainmentCategoryController(s services.EntertainmentCategoryService) EntertainmentCategoryController {
    return EntertainmentCategoryController{
        service: s,
    }
}

// GetEntertainmentCategories : GetEntertainmentCategories controller
func (p EntertainmentCategoryController) GetEntertainmentCategories(ctx *gin.Context) {
    var entertainment_categories models.EntertainmentCategory

    search := ctx.Query("search")

    data, _, err := p.service.FindAll(entertainment_categories, search)

    if err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find entertainment category")
        return
    }
    respArr := make([]map[string]interface{}, 0, 0)

    for _, n := range *data {
        resp := n.ResponseMap()
        respArr = append(respArr, resp)
    }

    ctx.JSON(http.StatusOK, &utils.Response{
        Success: true,
        Message: "Entertainment category result set",
        Data: respArr,
    })
}

// AddEntertainmentCategory : AddEntertainmentCategory controller
func (p *EntertainmentCategoryController) AddEntertainmentCategory(ctx *gin.Context) {
    var entertainmentCategory models.EntertainmentCategory
    ctx.ShouldBindJSON(&entertainmentCategory)

    if entertainmentCategory.Name == "" {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
        return
    }

    err := p.service.Save(entertainmentCategory)
    if err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create entertainment category")
        return
    }

    utils.SuccessJSON(ctx, http.StatusCreated, "Successfully created entertainment category")
}

//GetEntertainmentCategory : get entertainment category by id
func (p *EntertainmentCategoryController) GetEntertainmentCategory(c *gin.Context) {
    idParam := c.Param("id")
    
    var bus models.EntertainmentCategory
    bus.ID = idParam
    foundBus, err := p.service.Find(bus)
    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "Error finding entertainment category")
        return 
    }
    response := foundBus.ResponseMap()

    c.JSON(http.StatusOK, &utils.Response{
        Success: true,
        Message: "Result set of entertainment category",
        Data:    &response})

}

//UpdateEntertainmentCategory : get update by id
func (p EntertainmentCategoryController) UpdateEntertainmentCategory(ctx *gin.Context) {
    idParam := ctx.Param("id")

    var entertainmentCategory models.EntertainmentCategory
    entertainmentCategory.ID = idParam

    entertainmentCategoryRecord, err := p.service.Find(entertainmentCategory)

    if err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Entertainment category with given id not found")
        return
    }
    ctx.ShouldBindJSON(&entertainmentCategoryRecord)

    if entertainmentCategoryRecord.Name == "" {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
        return
    }

    if err := p.service.Update(entertainmentCategoryRecord); err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to update entertainment category")
        return
    }
    response := entertainmentCategoryRecord.ResponseMap()

    ctx.JSON(http.StatusOK, &utils.Response{
        Success: true,
        Message: "Successfully updated entertainment category",
        Data:    response,
    })
}

//DeleteEntertainmentCategory : Deletes Entertainment Category
func (p *EntertainmentCategoryController) DeleteEntertainmentCategory(c *gin.Context) {
    idParam := c.Param("id")
    
    err := p.service.Delete(idParam)

    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "Failed to delete entertainment category")
        return
    }
    response := &utils.Response{
        Success: true,
        Message: "Deleted sucessfully"}
    c.JSON(http.StatusOK, response)
}