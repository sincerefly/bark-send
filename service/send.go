package service

import (
	"bark-send/config"
	"bark-send/utils"
)

func SendAll(message string) {
	cfg := config.GetConfig()
	for _, token := range cfg.Users {
		url := utils.MakeUrl(cfg.Service.BaseUrl, token, message)
		go utils.Get(url)
	}
}
