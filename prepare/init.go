package prepare

import (
	"bark-send/config"
	"bark-send/logger"
)

func Init() {
	cfg := config.InitConfig()

	//_ = os.Setenv("BARK_SEND_DEBUG", "true")
	logger.InitLogger(&cfg.Log)
}
