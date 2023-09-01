package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/internal/patient/provider"
	"github.com/qbem-repos/patient-service/internal/shared/util"
	"github.com/qbem-repos/patient-service/internal/shared/validation"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func UpdatingPatientFields(c *gin.Context) {

	var slug = c.Param("slug")
	if slug == "" {
		util.WriteErrorMessage(c, http.StatusNotFound, nil, "Slug is required")
		return
	}
	var patientUpdateValue []provider.PatientChangeData
	errBind := c.Bind(&patientUpdateValue)

	if errBind != nil {
		util.WriteErrorMessage(c, http.StatusNotFound, errBind, "Json format is invalid")
		return
	}

	provider := provider.NewPatientProvider()
	for _, v := range patientUpdateValue {
		if slug == "" || v.Key == "" || v.Value == "" {
			util.WriteErrorMessage(c, http.StatusNotFound, errBind, "Slug, key and value are required")
			return
		}
		errUpdate := provider.UpdateOne(slug, v.Key, v.Value)
		if errUpdate != nil {
			util.WriteErrorMessage(c, http.StatusInternalServerError, errUpdate, fmt.Sprintf("Error updating field %s of patient", v.Key))
			return
		}
	}
	c.Status(http.StatusNoContent)
}

func UpdatingPatientBySlug(c *gin.Context) {

	var slug = c.Param("id")
	if slug == "" {
		util.WriteErrorMessage(c, http.StatusNotFound, nil, "Slug is required")
		return
	}

	var patientUpdateValue fhir.Patient
	errBind := c.Bind(&patientUpdateValue)
	if errValidation := validation.Validation(patientUpdateValue); errValidation != nil {
		util.WriteError(c, http.StatusBadRequest, errValidation)
		return
	}
	if errBind != nil {
		util.WriteErrorMessage(c, http.StatusNotFound, errBind, "Json format is invalid")
		return
	}

	provider := provider.NewPatientProvider()
	_, errFinding := provider.PullOneBySlug(slug)
	if errFinding != nil {
		util.WriteErrorMessage(c, http.StatusNotFound, errFinding, "patient not found")
		return
	}

	errUpdate := provider.UpdateOnePatientBySlug(slug, &patientUpdateValue)
	if errUpdate != nil {
		util.WriteErrorMessage(c, http.StatusInternalServerError, errUpdate, "error updating patient")
		return
	}

	c.Status(http.StatusNoContent)
}
