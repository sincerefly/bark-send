package service

import (
	"bark-send/config"
	"bark-send/logger"
	"bark-send/utils"
	"go.uber.org/zap"
)

func SendAll(message string) {
	cfg := config.GetConfig()
	for _, token := range cfg.Users {
		url := utils.MakeUrl(cfg.Service.BaseUrl, token, message)
		go utils.Get(url)
		logger.Info("send", zap.String("url", url))
	}
}
