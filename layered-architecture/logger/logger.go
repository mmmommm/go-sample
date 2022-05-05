package logger

import (
	"github.com/mmmommm/go-sample/config"
	"go.uber.org/zap"
)

type Logger = zap.SugaredLogger

func ProvideZapLogger(config *config.Config) (*Logger, func(), error) {
	var log *zap.Logger
	var err error
	if config.Env == "local" {
		log, err = zap.NewDevelopment()
	} else {
		log, err = zap.NewProduction()
	}
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err := log.Sync(); err != nil {
			log.Error("logger sync error")
		}
	}

	return log.Sugar(), cleanup, nil
}
