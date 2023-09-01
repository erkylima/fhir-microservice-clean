package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/adapter"
	"github.com/qbem-repos/patient-service/internal/patient/provider"
	"github.com/qbem-repos/patient-service/internal/shared/util"
	"github.com/qbem-repos/patient-service/internal/shared/validation"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func PostingPatient(c *gin.Context) {
	patientEtty, err := adapter.RequestToPatientDBDatamodel(c)
	if err != nil {
		util.WriteError(c, http.StatusUnprocessableEntity, err)
		return
	}

	if errValidation := validation.Validation(patientEtty); errValidation != nil {
		util.WriteError(c, http.StatusBadRequest, errValidation)

		return
	}
	slug, shouldReturn := saveInRepositoryOrReturnError(patientEtty, c)
	if shouldReturn {
		return
	}

	c.Writer.Header().Add("Content-Location", fmt.Sprintf("%s/%s", c.Request.URL.Path, *patientEtty.Id))
	c.Writer.Header().Add("X-Patient-Slug", slug)
	c.Writer.Header().Add("X-Patient-RegistryCode", *patientEtty.Id)
	c.Status(http.StatusCreated)
}

func saveInRepositoryOrReturnError(patient *fhir.Patient, c *gin.Context) (string, bool) {
	provider := provider.NewPatientProvider()
	slug, err := provider.Push(patient)
	if err != nil {

		util.WriteError(c, http.StatusBadRequest, err)
		return "", true
	}
	return slug, false
}
