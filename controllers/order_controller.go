package controllers

import (
	"net/http"

	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/model"
	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func (OrderController) GetAllOrder(c *gin.Context) {
	var orders []model.Order

	result := config.DB.Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Data Berhasil Diambil",
		Data:    orders,
	})
}

func (OrderController) GetOrder(c *gin.Context) {
	id := c.Param("id")

	var order model.Order
	result := config.DB.First(&order, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Data Berhasil Diambil",
		Data:    order,
	})
}

func (OrderController) CreateOrder(c *gin.Context) {
	var order model.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	order.ID = 0

	result := config.DB.Create(&order)

	if result.Error != nil {
		c.JSON(500, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		Status:  true,
		Message: "Order Baru Ditambahkan",
		Data:    order,
	})
}

func (OrderController) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var order model.Order

	var existing model.Order
	errExist := config.DB.First(&existing, id).Error
	if errExist != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Order Tidak Ditemukan",
		})
		return
	}

	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	result := config.DB.Model(&existing).Updates(order)

	if result.Error != nil {
		c.JSON(500, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  false,
		Message: "Order Berhasil Diupdate",
		Data:    existing,
	})
}

func (OrderController) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	var existing model.Order
	if err := config.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Order Tidak Ditemukan",
		})
		return
	}

	result := config.DB.Delete(&model.Order{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Order Gagal Dihapus",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Order Berhasil Dihapus",
	})
}

