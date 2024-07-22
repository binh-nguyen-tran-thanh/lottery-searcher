package restful

import (
	"backend/internal/core/util/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorHandler(c *gin.Context, err error) {
	if err == nil {
		c.AbortWithStatusJSON(http.StatusOK, nil)
		return
	}
	fail, ok := err.(*exception.Exception)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(fail))
		return
	}
	if !fail.HasError() {
		fail.AddError("exception", fail.Message)
	}
	var statusCode int
	switch fail.Type {
	case exception.TypeNotFound:
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}
	c.AbortWithStatusJSON(statusCode, NewErrorResponse(fail))
}
