package controllers

import (
	"net/http"

	"github.com/GeZaM8/laundry-be/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (cc *CategoryController) GetAll(c *gin.Context) {
	var categories []model.Category
	result := cc.DB.Find(&categories)
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

func (cc *CategoryController) GetByID(c *gin.Context) {
	id := c.Param("id")

	var category model.Category
	result := cc.DB.First(&category, id)

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

func (cc *CategoryController) Create(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	category.ID = 0

	result := cc.DB.Create(&category)
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

func (cc *CategoryController) Update(c *gin.Context) {
	id := c.Param("id")

	var existing model.Category
	if err := cc.DB.First(&existing, id).Error; err != nil {
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

	result := cc.DB.Model(&existing).Updates(body)
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

func (cc *CategoryController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := cc.DB.Delete(&model.Category{}, id).Error; err != nil {
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
