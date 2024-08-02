package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) GetAllRegion(c *gin.Context) {
	regions, err := s.service.Region().GetAllRegion()
	if err != nil {
		s.logger.Error().Err(err).Msg("Fail")
		errorHandler(c, err)
		return
	}

	res := NewSuccessResponse(serializeRegion(regions))

	c.JSON(http.StatusOK, res)
}
