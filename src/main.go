package main

import (
	"mysadapi/dataSource"
	"mysadapi/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	logs.InitializeLogs()
	logs.PostLog("INFO", "Logs Initialized!")
	router := gin.Default()
	logs.PostLog("INFO", "gin router initialized!")
	InitDB()
	logs.PostLog("INFO", "Go ToDo API connected with PostgreSQL Database!")
	//Cerrar la db si termina la ejecución de main
	defer db.Close()

	//Crea la tabla ToDo si no existe
	CreateToDoTable()

	dataSource.SetDB(db)
	ConfigurarRutas(router)
	//Cierra el fichero de logs (archivo Service)
	defer logs.CloseLogs()
	router.Run(":8080")
}
