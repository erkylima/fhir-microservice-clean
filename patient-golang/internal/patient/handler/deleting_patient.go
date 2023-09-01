package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/internal/patient/provider"
	"github.com/qbem-repos/patient-service/internal/shared/util"
)

func DeletingPatient(c *gin.Context) {

	var slug = c.Param("slug")

	provider := provider.NewPatientProvider()
	if slug == "" {
		util.WriteError(c, http.StatusNotFound, errors.New("slug is required"))
		return
	}
	deleted, errUpdate := provider.DeleteOneBySlug(slug)
	if errUpdate != nil || deleted == 0 {
		util.WriteError(c, http.StatusNotFound, errUpdate)
		return
	}
	c.Status(http.StatusNoContent)
}
