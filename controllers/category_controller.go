package controllers

import (
	"net/http"

	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/model"
	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

func (CategoryController) GetAll(c *gin.Context) {
	var categories []model.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Categories Retrieved",
		Data:    categories,
	})
}

func (CategoryController) GetByID(c *gin.Context) {
	id := c.Param("id")

	var category model.Category
	result := config.DB.First(&category, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Category Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Category Retrieved",
		Data:    category,
	})
}

func (CategoryController) Create(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	category.ID = 0

	result := config.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		Status:  true,
		Message: "Category Created",
		Data:    category,
	})
}

func (CategoryController) Update(c *gin.Context) {
	id := c.Param("id")

	var existing model.Category
	if err := config.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Category Not Found",
		})
		return
	}

	var body model.Category
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	result := config.DB.Model(&existing).Updates(body)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Category Updated",
		Data:    existing,
	})
}

func (CategoryController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&model.Category{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Category Delete Failed",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Category Deleted",
	})
}
