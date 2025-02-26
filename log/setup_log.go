package log

import (
	"backend/config"

	"go.uber.org/zap"
)

func Setup(cfg *config.Log) {
	logger := newLogger(cfg)
	defer func() {
		_ = logger.Sync()
	}()
	zap.ReplaceGlobals(logger)
}
