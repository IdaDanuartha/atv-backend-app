package controllers

import (
	"net/http"
	"strconv"

	"github.com/IdaDanuartha/teman-bus-backend-app/app/api/services"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/models"
	"github.com/IdaDanuartha/teman-bus-backend-app/utils"
	"github.com/gin-gonic/gin"
)

//BusController -> BusController
type BusController struct {
    service services.BusService
}

//NewBusController : NewBusController
func NewBusController(s services.BusService) BusController {
    return BusController{
        service: s,
    }
}

// GetBuses : GetBuses controller
func (p BusController) GetBuses(ctx *gin.Context) {
    var buses models.Bus

    keyword := ctx.Query("keyword")

    data, _, err := p.service.FindAll(buses, keyword)

    if err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find bus")
        return
    }
    respArr := make([]map[string]interface{}, 0, 0)

    for _, n := range *data {
        resp := n.ResponseMap()
        respArr = append(respArr, resp)
    }

    ctx.JSON(http.StatusOK, &utils.Response{
        Success: true,
        Message: "Bus result set",
        Data: respArr,
    })
}

// AddBus : AddBus controller
func (p *BusController) AddBus(ctx *gin.Context) {
    var bus models.Bus
    ctx.ShouldBindJSON(&bus)

    if bus.Name == "" {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
        return
    }
    if bus.LicensePlate == "" {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "License Plate is required")
        return
    }
    err := p.service.Save(bus)
    if err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create bus")
        return
    }
    utils.SuccessJSON(ctx, http.StatusCreated, "Successfully Created bus")
}

//GetBus : get bus by id
func (p *BusController) GetBus(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "id invalid")
        return
    }
    var bus models.Bus
    bus.ID = id
    foundBus, err := p.service.Find(bus)
    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "Error Finding Bus")
        return
    }
    response := foundBus.ResponseMap()

    c.JSON(http.StatusOK, &utils.Response{
        Success: true,
        Message: "Result set of Bus",
        Data:    &response})

}

//UpdateBus : get update by id
func (p BusController) UpdateBus(ctx *gin.Context) {
    idParam := ctx.Param("id")

    id, err := strconv.ParseInt(idParam, 10, 64)

    if err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
        return
    }
    var bus models.Bus
    bus.ID = id

    busRecord, err := p.service.Find(bus)

    if err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Bus with given id not found")
        return
    }
    ctx.ShouldBindJSON(&busRecord)

    if busRecord.Name == "" {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
        return
    }
    if busRecord.LicensePlate == "" {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "License Plate is required")
        return
    }

    if err := p.service.Update(busRecord); err != nil {
        utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Bus")
        return
    }
    response := busRecord.ResponseMap()

    ctx.JSON(http.StatusOK, &utils.Response{
        Success: true,
        Message: "Successfully Updated Bus",
        Data:    response,
    })
}

//DeleteBus : Deletes Bus
func (p *BusController) DeleteBus(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "id invalid")
        return
    }
    err = p.service.Delete(id)

    if err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Bus")
        return
    }
    response := &utils.Response{
        Success: true,
        Message: "Deleted Sucessfully"}
    c.JSON(http.StatusOK, response)
}