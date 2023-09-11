package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/internal/patient/router"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	store := persistence.NewInMemoryStore(time.Second)

	r := gin.Default()
	router.RoutesRegistry(r, store)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("error on start server:", err)
	}
}
