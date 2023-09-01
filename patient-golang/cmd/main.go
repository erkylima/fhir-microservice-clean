package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/internal/patient/router"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	r := gin.Default()
	router.RoutesRegistry(r)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("error on start server:", err)
	}
}
