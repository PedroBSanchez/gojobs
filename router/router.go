package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {

	//Initialize Router
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	router.Run(":" + port)
}
