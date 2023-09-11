package router

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/internal/patient/handler"
)

func RoutesRegistry(r *gin.Engine, store *persistence.InMemoryStore) {
	r.POST("/patients", cache.CachePage(store, time.Minute, handler.PostingPatient))
	r.GET("/patients/:id", cache.CachePage(store, time.Minute, handler.GettingPatient))

	r.GET("/patients", cache.CachePage(store, time.Minute, handler.ListAllPatients))
	r.PUT("/patients/:id", cache.CachePage(store, time.Minute, handler.UpdatingPatientBySlug))
	r.DELETE("/patients/:id", cache.CachePage(store, time.Minute, handler.DeletingPatient))

}
