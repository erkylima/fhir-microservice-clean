package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qbem-repos/patient-service/internal/patient/provider"
	"github.com/qbem-repos/patient-service/internal/shared/util"
)

func ListAllPatients(c *gin.Context) {
	provider := provider.NewPatientProvider()
	var filter = prepare(c)
	var err error

	patients, err := provider.Pull(*filter)
	if err != nil {
		util.WriteError(c, http.StatusBadRequest, err)
	}
	c.Writer.Header().Add("X-Total", fmt.Sprintf("%d", len(*patients)))

	c.JSON(http.StatusOK, gin.H{
		"patients": patients,
		"metadata": gin.H{
			"total":  provider.Count(filter),
			"offset": filter.Offset,
			"limit":  filter.Limit,
		},
	})
}

func prepare(c *gin.Context) *provider.PatientPullFilter {
	var filter = new(provider.PatientPullFilter)

	if c.Query("limit") != "" {
		limit, _ := strconv.Atoi(c.Query("limit"))
		filter.Limit = int64(limit)

	}

	if c.Query("offset") != "" {
		offset, _ := strconv.Atoi(c.Query("offset"))
		filter.Offset = int64(offset)
	}

	return filter
}
