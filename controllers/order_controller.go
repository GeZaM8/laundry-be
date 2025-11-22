package controllers

import (
	"net/http"

	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/model"
	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func (OrderController) GetAll(c *gin.Context) {
	var orders []model.Order
	result := config.DB.Preload("Customer").Preload("Items.Category").Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Orders Retrieved",
		Data:    orders,
	})
}

func (OrderController) GetByID(c *gin.Context) {
	id := c.Param("id")

	var order model.Order
	result := config.DB.Preload("Customer").Preload("Items.Category").First(&order, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status: false,
			// Message: "Order Not Found",
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Order Retrieved",
		Data:    order,
	})
}

func (OrderController) Create(c *gin.Context) {
	var body model.Order
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	var total float64 = 0

	for i := range body.Items {
		var cat model.Category
		config.DB.First(&cat, body.Items[i].CategoryID)

		var harga float64 = cat.PricePerUnit

		if body.Items[i].Unit == "kg" {
			body.Items[i].Price = harga * body.Items[i].WeightKg
		} else {
			body.Items[i].Price = harga * float64(body.Items[i].Qty)
		}

		total += body.Items[i].Price
	}

	body.TotalPrice = total

	var customer model.Customer

	if body.Customer.Phone != "" {
		config.DB.Where("phone = ?", body.Customer.Phone).First(&customer)

		if (customer.ID) == 0 {

			customer.Name = body.Customer.Name
			if customer.Name == "" {
				customer.Name = "Pelanggan " + body.Customer.Phone
			}
			customer.Phone = body.Customer.Phone
			config.DB.Create(&customer)
		}

		body.CustomerID = customer.ID
	}

	result := config.DB.Create(&body)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		Status:  true,
		Message: "Order Created",
		Data:    body,
	})
}

func (OrderController) Update(c *gin.Context) {
	id := c.Param("id")

	var order model.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Order Not Found",
		})
		return
	}

	var body model.Order
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	config.DB.Model(&order).Updates(body)

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Order Updated",
		Data:    order,
	})
}

func (OrderController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&model.Order{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Delete Failed",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Order Deleted",
	})
}
