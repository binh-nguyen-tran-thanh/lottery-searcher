package global

import (
	"backend/config"
)

var (
	Config config.Config
)

func New(config config.Config) {
	Config = config
}
