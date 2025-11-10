package application

import (
	"device-service/internal/config"

	"go.uber.org/zap"
)

type App struct {
	Log  *zap.Logger
	Conf *config.Config
}
