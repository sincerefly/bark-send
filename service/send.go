package service

import (
	"bark-send/utils"
	"bark-send/utils/config"
	"bark-send/utils/logger"
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
