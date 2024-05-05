package controllers

import (
	"net/http"
	models "stockManagment/DbStore/Models"
	dtos "stockManagment/Dtos"
	utilities "stockManagment/Utilities"
	utilities_controllers "stockManagment/Utilities/Controllers"
	cookies_utilities "stockManagment/Utilities/Cookies"

	"github.com/gin-gonic/gin"
)

func GetAllWearhouses(c *gin.Context) []models.Warehouse {
	currentCompanyId := cookies_utilities.GetCurrentCompanyId(c)
	wearhouses := utilities.GetAllWearhouses(currentCompanyId)
	return wearhouses
}

func CreateWearhouse(c *gin.Context) {
	var req string

	err := c.ShouldBind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	userId := cookies_utilities.RetriveCurrentUserId(c)
	companyId := cookies_utilities.GetCurrentCompanyId(c)
	success := utilities.CreateWearhouse(req, companyId, userId)

	if success {
		c.IndentedJSON(http.StatusOK, true)
		return
	}

	c.IndentedJSON(http.StatusOK, false)
}

func DeleteWarehouse(c *gin.Context) {
	var req string

	err := c.ShouldBind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err = utilities.DeleteWarehouse(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, true)
}

func GetWarehouseName(c *gin.Context) {
	id := c.Param("id")

	warehouse, err := utilities.GetWarehouseById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, warehouse.Name)
}

func EditWarehouse(c *gin.Context) {
	validate := utilities_controllers.InitValidator()
	var req dtos.WarehouseInfo

	err := c.ShouldBind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid request body.")
		return
	}

	err = validate.Struct(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = utilities.EditWarehouse(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}
