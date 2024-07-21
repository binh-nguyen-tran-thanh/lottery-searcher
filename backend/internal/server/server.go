package server

import (
	"fmt"
	"net/http"
	"time"

	"backend/internal/database"
	"backend/internal/global"
)

type Server struct {
	addr string

	db database.Service
}

func NewServer() *http.Server {
	NewServer := &Server{
		addr: fmt.Sprintf("%s:%s", global.Config.Server.Host, global.Config.Server.Port),
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         NewServer.addr,
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Server is running at: %s \n", NewServer.addr)

	return server
}
