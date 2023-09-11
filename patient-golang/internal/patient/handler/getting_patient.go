package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/internal/patient/provider"
	"github.com/qbem-repos/patient-service/internal/shared/util"
)

func GettingPatient(c *gin.Context) {
	var id = c.Param("id")

	provider := provider.NewPatientProvider()
	patient, err := provider.PullOne(id)

	if err != nil {
		util.WriteError(c, http.StatusBadRequest, err)
	}

	if patient == nil {
		util.WriteErrorMessage(c, http.StatusNotFound, err, "patient not found")
		return
	}

	c.JSON(http.StatusOK, &patient)
}
