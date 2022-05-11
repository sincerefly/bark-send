package prepare

import (
	"bark-send/utils/config"
	"bark-send/utils/logger"
	"bark-send/utils/sqlite3"
)

func Init() {
	cfg := config.InitConfig()

	//_ = os.Setenv("BARK_SEND_DEBUG", "true")
	logger.InitLogger(&cfg.Log)

	sqlite3.InitStore()
}
