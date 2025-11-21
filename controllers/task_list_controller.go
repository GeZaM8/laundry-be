package controllers

import (
	"net/http"

	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/model"
	"github.com/gin-gonic/gin"
)

type TaskListController struct{}

func (TaskListController) GetAllTaskList(c *gin.Context) {
	var taskLists []model.TaskList

	result := config.DB.Find(&taskLists)
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
		Data:    taskLists,
	})
}

func (TaskListController) GetTaskList(c *gin.Context) {
	id := c.Param("id")

	var taskList model.TaskList

	result := config.DB.First(&taskList, id)

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
		Data:    taskList,
	})
}

func (TaskListController) CreateTaskList(c *gin.Context) {
	var taskList model.TaskList

	err := c.ShouldBindJSON(&taskList)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	taskList.ID = 0

	result := config.DB.Create(&taskList)

	if result.Error != nil {
		c.JSON(500, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		Status:  true,
		Message: "Task List Baru Ditambahkan",
		Data:    taskList,
	})
}

func (TaskListController) UpdateTaskList(c *gin.Context) {
	id := c.Param("id")

	var taskList model.TaskList

	var existing model.TaskList
	errExist := config.DB.First(&existing, id).Error
	if errExist != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Task List Tidak Ditemukan",
		})
		return
	}

	err := c.ShouldBindJSON(&taskList)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	result := config.DB.Model(&existing).Updates(taskList)

	if result.Error != nil {
		c.JSON(500, model.Response{
			Status:  false,
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Task List Berhasil Diubah",
		Data:    taskList,
	})
}
