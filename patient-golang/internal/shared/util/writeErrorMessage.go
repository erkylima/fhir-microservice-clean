package util

import (
	"github.com/gin-gonic/gin"
)

func WriteError(c *gin.Context, status int, err error) {
	c.Writer.Header().Add("X-Error-Message", err.Error())
	c.Writer.WriteHeader(status)
	c.Writer.WriteString(err.Error())
	c.Status(status)
}

func WriteErrorMessage(c *gin.Context, status int, err error, message string) {
	c.Writer.Header().Add("X-Error-Message", message)
	c.Writer.WriteHeader(status)
	c.Status(status)
}
