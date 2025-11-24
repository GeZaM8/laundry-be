package controllers

import (
	"net/http"

	"github.com/GeZaM8/laundry-be/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (o *OrderController) GetAll(c *gin.Context) {
	var orders []model.Order
	result := o.DB.Preload("Customer").Preload("Items.Category").Find(&orders)
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

func (o *OrderController) GetByID(c *gin.Context) {
	id := c.Param("id")

	var order model.Order
	result := o.DB.Preload("Customer").Preload("Items.Category").First(&order, id)

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

func (o *OrderController) Create(c *gin.Context) {
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
		o.DB.First(&cat, body.Items[i].CategoryID)

		var harga float64 = cat.PricePerUnit

		if body.Items[i].Unit == "kg" {
			body.Items[i].Price = harga * body.Items[i].WeightKg
		} else {
			body.Items[i].Price = harga * float64(body.Items[i].Qty)
		}

		total += body.Items[i].Price
	}

	body.TotalPrice = total

	var customer model.User

	if body.Customer.Phone != "" {
		o.DB.Where("phone = ?", body.Customer.Phone).First(&customer)

		if (customer.ID) == 0 {

			customer.Name = body.Customer.Name
			if customer.Name == "" {
				customer.Name = "Pelanggan " + body.Customer.Phone
			}
			customer.Phone = body.Customer.Phone
			o.DB.Create(&customer)
		}

		body.CustomerID = customer.ID
	}

	result := o.DB.Create(&body)
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

func (o *OrderController) Update(c *gin.Context) {
	id := c.Param("id")

	var order model.Order
	if err := o.DB.First(&order, id).Error; err != nil {
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

	o.DB.Model(&order).Updates(body)

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Order Updated",
		Data:    order,
	})
}

func (o *OrderController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := o.DB.Delete(&model.Order{}, id).Error; err != nil {
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
