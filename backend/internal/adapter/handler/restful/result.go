package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) GetTodayResult(c *gin.Context) {
	results, err := s.service.Result().GetTodayResult()

	if err != nil {
		s.logger.Error().Err(err).Msg("Fail")
		errorHandler(c, err)
		return
	}

	res := NewSuccessResponse(serializeResults(results))

	c.JSON(http.StatusOK, res)
}

func (s Server) FilterResult(c *gin.Context) {
	var payload FilterResultPayload

	if err := c.ShouldBind(&payload); err != nil {
		s.logger.Error().Err(err).Msg("Fail")
		errorHandler(c, err)
		return
	}

	results, err := s.service.Result().FilterResult(AsPortFilterResultPayload(payload))

	if err != nil {
		s.logger.Error().Err(err).Msg("Fail")
		errorHandler(c, err)
		return
	}

	res := NewSuccessResponse(serializeResults(results))

	c.JSON(http.StatusOK, res)
}

func (s Server) FilterOpenNumb(c *gin.Context) {
	var payload FilterOpenNumbPayload

	if err := c.ShouldBind(&payload); err != nil {
		s.logger.Error().Err(err).Msg("Fail")
		errorHandler(c, err)
		return
	}

	results, err := s.service.Result().FilterOpenNumb(AsPortFilterOpenNumbPayload(payload))

	if err != nil {
		s.logger.Error().Err(err).Msg("Fail")
		errorHandler(c, err)
		return
	}

	res := NewSuccessResponse(results)

	c.JSON(http.StatusOK, res)
}
