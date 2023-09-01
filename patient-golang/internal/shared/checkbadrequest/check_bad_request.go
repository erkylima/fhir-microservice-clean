package checkbadrequest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IfErroReturnBadRequest(err error, c *gin.Context) {
	if err != nil {
		c.Writer.Header().Add("X-Error-Message", err.Error())
		c.Status(http.StatusBadRequest)
	}
}
