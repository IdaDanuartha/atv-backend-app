package controllers

import (
	"net/http"
	"strconv"

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
    var buses models.EntertainmentCategory

    keyword := ctx.Query("keyword")

    data, _, err := p.service.FindAll(buses, keyword)

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
    var bus models.EntertainmentCategory
    ctx.ShouldBindJSON(&bus)

    if bus.Name == "" {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
        return
    }
    utils.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Entertainment Category")
}

//GetEntertainmentCategory : get entertainment category by id
func (p *EntertainmentCategoryController) GetEntertainmentCategory(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "id invalid")
        return
    }
    var bus models.EntertainmentCategory
    bus.ID = id
    foundBus, err := p.service.Find(bus)
    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "Error Finding Entertainment Category")
        return 
    }
    response := foundBus.ResponseMap()

    c.JSON(http.StatusOK, &utils.Response{
        Success: true,
        Message: "Result set of Entertainment Category",
        Data:    &response})

}

//UpdateEntertainmentCategory : get update by id
func (p EntertainmentCategoryController) UpdateEntertainmentCategory(ctx *gin.Context) {
    idParam := ctx.Param("id")

    id, err := strconv.ParseInt(idParam, 10, 64)

    if err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
        return
    }
    var entertainmentCategory models.EntertainmentCategory
    entertainmentCategory.ID = id

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
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store entertainment category")
        return
    }
    response := entertainmentCategoryRecord.ResponseMap()

    ctx.JSON(http.StatusOK, &utils.Response{
        Success: true,
        Message: "Successfully Updated Entertainment Category",
        Data:    response,
    })
}

//DeleteEntertainmentCategory : Deletes Bus
func (p *EntertainmentCategoryController) DeleteEntertainmentCategory(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "id invalid")
        return
    }
    err = p.service.Delete(id)

    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "Failed to delete entertainment category")
        return
    }
    response := &utils.Response{
        Success: true,
        Message: "Deleted Sucessfully"}
    c.JSON(http.StatusOK, response)
}