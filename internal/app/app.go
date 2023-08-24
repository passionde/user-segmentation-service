package app

import (
	"github.com/passionde/user-segmentation-service/config"
	log "github.com/sirupsen/logrus"
)

func Run(configPath string) {
	// Config
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Log.Level)
}
