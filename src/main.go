package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	ConfigurarRutas(router)

	router.Run(":8080")
}
