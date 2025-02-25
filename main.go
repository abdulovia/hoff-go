package main

import (
	"github.com/abdulovia/hoff-go/config"
	"github.com/abdulovia/hoff-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	config.LoadEnv() // !before LoadDatabase
	config.LoadDatabase()

	r.Run() // listen and serve on 0.0.0.0:8080
}
