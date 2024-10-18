package main

import (
	"mysadapi/dataSource"
	"mysadapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConfigurarRutas(r *gin.Engine) {
	r.GET("/", getToDos)
	r.POST("/", insertToDo)
	r.GET("/:id", getToDoByID)
	r.PUT("/:id", updateToDo)
	r.DELETE("/:id", deleteToDo)
	r.GET("/complete/:id", completeByID)
	r.GET("/title/:title", getByTitle)
}

func getToDos(c *gin.Context) {
	res, err := dataSource.GetToDos()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func insertToDo(c *gin.Context) {
	var newToDo models.ToDo
	if err := c.ShouldBindJSON(&newToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := dataSource.CreateToDo(newToDo.Title, newToDo.Description, newToDo.Completed)
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{"error": "error trying to create the toDo"})
		return
	}
	newToDo = res
	c.JSON(http.StatusCreated, newToDo)
}

func getToDoByID(c *gin.Context) {
	id := getIDFromQuery(c)
	if id == -1 {
		return
	}
	res, err := dataSource.GetToDosWhere("id =", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res[0])
}

func updateToDo(c *gin.Context) {
	id := getIDFromQuery(c)
	if id == -1 {
		return
	}
	var updatedToDo models.ToDo
	if err := c.ShouldBindJSON(&updatedToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedToDo.Title != "" {
		err := dataSource.UpdateToDo(id, "title", updatedToDo.Title)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with ID " + strconv.Itoa(id) + " not found!"})
		}
	}
	if updatedToDo.Description != "" {
		err := dataSource.UpdateToDo(id, "description", updatedToDo.Description)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with ID " + strconv.Itoa(id) + " not found!"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"operation": "ToDo with ID " + strconv.Itoa(id) + " updated correctly!"})
}

func deleteToDo(c *gin.Context) {
	id := getIDFromQuery(c)
	if id == -1 {
		return
	}
	err := dataSource.DeleteToDo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with ID " + strconv.Itoa(id) + " not deleted!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"operation": "ToDo with ID " + strconv.Itoa(id) + " deleted correctly!"})
}

func completeByID(c *gin.Context) {
	id := getIDFromQuery(c)
	if id == -1 {
		return
	}

	err := dataSource.UpdateToDo(id, "completed", true)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with ID " + strconv.Itoa(id) + " not found!"})
	}
	c.JSON(http.StatusOK, gin.H{"operation": "ToDo with ID " + strconv.Itoa(id) + " completed!"})
}

func getByTitle(c *gin.Context) {
	title := "%" + c.Param("title") + "%"

	res, err := dataSource.GetToDosWhere("title LIKE ", title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with title " + title + " not found!"})
		return
	}

	c.JSON(http.StatusOK, res)

}

func getIDFromQuery(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": c.Param("id") + " is not a valid ID"})
		return -1
	}
	return id
}
