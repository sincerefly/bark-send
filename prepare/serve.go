package prepare

import (
	"bark-send/service"
	"github.com/jasonlvhit/gocron"
	"log"
	"time"
)

func Server() {
	Init()
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
}

func Start() {
	<-gocron.Start()
}
