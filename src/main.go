package main

import (
	"mysadapi/dataSource"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	InitDB()
	//Cerrar la db si termina la ejecución de main
	defer db.Close()

	//Crea la tabla ToDo si no existe
	CreateToDoTable()

	dataSource.SetDB(db)
	ConfigurarRutas(router)

	router.Run(":8080")
}
