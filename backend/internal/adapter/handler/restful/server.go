package restful

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/core/port"
	"backend/internal/core/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	service port.Service
	config  util.Config
	logger  port.Logger
	router  *gin.Engine
}

func NewServer(config util.Config, service port.Service, logger port.Logger) port.Server {
	server := &Server{
		service: service,
		config:  config,
		logger:  logger,
	}

	server.setUpRouter()
	return server
}

func (s *Server) setUpRouter() {
	r := gin.New()

	r.Use(gin.Recovery(), cors.Default())

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	})
	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"message": "no method provided"})
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	r.GET("/ping", s.Pong)
	r.GET("/health", s.Heathy)

	regionGroup := r.Group("/region")

	regionGroup.GET("/", s.GetAllRegion)

	s.router = r
}

func (s Server) Start() error {
	addr := fmt.Sprintf("%s:%s", s.config.Server.Host, s.config.Server.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	s.logger.Info().Msgf("Server is up and running on: %s", addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Fatal().Err(err).Msg("failed listen")
	}

	go func() {
		s.logger.Info().Msgf("Server is up and running on: %s", addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal().Err(err).Msg("failed listen")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.logger.Info().Msg("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Fatal().Err(err).Msg("failed shutdown server")
	}

	<-ctx.Done()
	s.logger.Info().Msg("timeout in 5 seconds")

	s.logger.Info().Msg("server exiting")

	return nil
}
