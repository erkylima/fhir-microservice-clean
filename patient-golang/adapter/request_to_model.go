package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func RequestToPatientDBDatamodel(c *gin.Context) (*fhir.Patient, error) {
	patientEtty := new(fhir.Patient)

	if err := c.ShouldBindJSON(&patientEtty); err != nil {
		c.Writer.Header().Add("X-Error-Message", err.Error())
		c.JSON(http.StatusBadRequest, "Error in payload. Verify your payload.")
		return &fhir.Patient{}, err
	}

	return patientEtty, nil
}
