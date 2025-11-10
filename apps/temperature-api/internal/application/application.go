package application

import (
	"temperature-api/internal/config"

	"go.uber.org/zap"
)

type App struct {
	Log  *zap.Logger
	Conf *config.Config
}
