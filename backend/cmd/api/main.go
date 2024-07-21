package main

import (
	"backend/config"
	"backend/internal/global"
	"backend/internal/server"
	"fmt"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Fail to load the config: %s", err))
	}

	global.New(config)

	server := server.NewServer()

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
