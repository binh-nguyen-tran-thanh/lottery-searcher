package logger

import (
	"backend/internal/core/port"
	"backend/internal/core/util"
)

func NewLogger(config util.Config) port.Logger {
	return NewZeroLogLogger(config)
}
