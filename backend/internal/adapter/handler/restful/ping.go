package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) Pong(c *gin.Context) {
	res, err := s.service.Ping().Ping()
	if err != nil {
		s.logger.Fatal().Err(err).Msg("Fail")
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) Heathy(c *gin.Context) {
	res, err := s.service.Ping().Healthy()
	if err != nil {
		s.logger.Fatal().Err(err).Msg("Fail")
	}
	c.JSON(http.StatusOK, res)
}
