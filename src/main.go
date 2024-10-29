package main

import (
	"mysadapi/dataSource"
	"mysadapi/logs"
	"mysadapi/service"

	"github.com/gin-gonic/gin"
)

func main() {
	logs.InitializeLogs()
	logs.PostLog("INFO", "Logs Initialized!")
	router := gin.Default()
	logs.PostLog("INFO", "gin router initialized!")
	dataSource.InitDB()
	logs.PostLog("INFO", "Go ToDo API connected with PostgreSQL Database!")
	//Cerrar la db si termina la ejecución de main
	defer dataSource.CloseDatabase()
	//Cierra el fichero de logs (archivo Service)
	defer logs.CloseLogs()

	service.ConfigurarRutas(router)

	router.Run(":8080")
}
