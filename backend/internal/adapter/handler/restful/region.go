package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegionResponse struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	IsActive string `json:"isActive"`
}

type AllRegionResponse struct {
	Regions []RegionResponse `json:"regions"`
}

func (s Server) GetAllRegion(c *gin.Context) {
	regions, err := s.service.Region().GetAllRegion()
	if err != nil {
		s.logger.Error().Err(err).Msg("Fail")
		errorHandler(c, err)
		return
	}

	res := NewSuccessResponse(regions)

	c.JSON(http.StatusOK, res)
}
