package prepare

import (
	"bark-send/api_service/router"
	"bark-send/service"
	"github.com/jasonlvhit/gocron"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func Server() {
	Init()
	go ListenApi()
	SetCron()
	Start()
}

func SetCron() {

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("Unfortunately can't load a location")
		log.Println(err)
	} else {
		gocron.ChangeLoc(location)
	}

	gocron.Every(1).Day().At("07:00").Do(service.SendAll, "⏰喝水啦 07:00")
	gocron.Every(1).Day().At("09:00").Do(service.SendAll, "⏰喝水啦 09:00")
	gocron.Every(1).Day().At("11:30").Do(service.SendAll, "⏰喝水啦 11:30")
	gocron.Every(1).Day().At("13:00").Do(service.SendAll, "⏰喝水啦 13:00")
	gocron.Every(1).Day().At("15:30").Do(service.SendAll, "⏰喝水啦 15:30")
	gocron.Every(1).Day().At("17:30").Do(service.SendAll, "⏰喝水啦 17:30")
	gocron.Every(1).Day().At("19:00").Do(service.SendAll, "⏰喝水啦 19:00")
	gocron.Every(1).Day().At("21:00").Do(service.SendAll, "⏰喝水啦 21:00")
	gocron.Every(1).Day().At("22:40").Do(service.SendAll, "⏰睡觉啦 22:40")

	//gocron.Every(5).Seconds().Do(service.SendAll, "⏰测试")
}

// Start 启动定时任务
func Start() {
	<-gocron.Start()
}

// ListenApi 初始化 HTTP 服务
func ListenApi() {
	e := echo.New()
	router.InitRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
