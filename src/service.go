package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var ToDoTest = ToDo{ID: 1, Title: "Tarea de prueba", Description: "Descripción de prueba", DateTime: time.Now(), Completed: false}
var nextID = 2
var ToDos []ToDo = []ToDo{ToDoTest}

func ConfigurarRutas(r *gin.Engine) {
	r.GET("/", getToDos)
	r.POST("/", insertToDo)
	r.GET("/:id", getToDoByID)
	r.PUT("/:id", updateToDo)
	r.DELETE("/:id", deleteToDo)
	r.GET("/complete/:id", completeByID)
}

func getToDos(c *gin.Context) {
	c.JSON(http.StatusOK, ToDos)
}

func insertToDo(c *gin.Context) {
	var newToDo ToDo
	if err := c.ShouldBindJSON(&newToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newToDo.ID = nextID
	newToDo.DateTime = time.Now()
	nextID++
	ToDos = append(ToDos, newToDo)

	c.JSON(http.StatusCreated, newToDo)
}

func getToDoByID(c *gin.Context) {
	id := getIDFromQuery(c)
	if id == -1 {
		return
	}
	for _, toDo := range ToDos {
		if toDo.ID == id {
			c.JSON(http.StatusOK, toDo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with ID " + strconv.Itoa(id) + " not found!"})
}

func updateToDo(c *gin.Context) {
	id := getIDFromQuery(c)
	if id == -1 {
		return
	}
	var updatedToDo ToDo
	if err := c.ShouldBindJSON(&updatedToDo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, toDo := range ToDos {
		if toDo.ID == id {
			//Si se pasa el campo title vacío, no se debe actualizar como ""
			if updatedToDo.Title == "" {
				updatedToDo.Title = toDo.Title
			}
			if updatedToDo.Description == "" {
				updatedToDo.Description = toDo.Description
			}

			updatedToDo.ID = toDo.ID
			updatedToDo.DateTime = toDo.DateTime
			ToDos[i] = updatedToDo
			c.JSON(http.StatusOK, updatedToDo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with ID " + strconv.Itoa(id) + " not found!"})
}

func deleteToDo(c *gin.Context) {
	id := getIDFromQuery(c)
	if id == -1 {
		return
	}
	for i, toDo := range ToDos {
		if toDo.ID == id {
			ToDos = append(ToDos[:i], ToDos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"operation": "ToDo with ID " + strconv.Itoa(id) + " successfully deleted!"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with ID " + strconv.Itoa(id) + " not found!"})
}

func completeByID(c *gin.Context) {
	id := getIDFromQuery(c)
	if id == -1 {
		return
	}
	for _, toDo := range ToDos {
		if toDo.ID == id {
			toDo.Completed = true
			c.JSON(http.StatusOK, toDo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ToDo with ID " + strconv.Itoa(id) + " not found!"})
}

func getIDFromQuery(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": c.Param("id") + " is not a valid ID"})
		return -1
	}
	return id
}
