package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/internal/patient/handler"
)

func RoutesRegistry(r *gin.Engine) {
	r.POST("/patients", handler.PostingPatient)
	r.GET("/patients/:id", handler.GettingPatient)

	r.GET("/patients", handler.ListAllPatients)
	r.PUT("/patients/:id", handler.UpdatingPatientBySlug)
	r.DELETE("/patients/:id", handler.DeletingPatient)

}
