package main

import (
	"context"
	"os"
	"os/signal"

	nested "github.com/Lyrics-you/sail-logrus-formatter/sailor"
	log "github.com/sirupsen/logrus"

	_ "github.com/huoxue1/tdlib/plugins/task"

	"github.com/huoxue1/tdlib/conf"
	"github.com/huoxue1/tdlib/lib"
)

func init() {
	log.SetFormatter(&nested.Formatter{
		FieldsOrder:           nil,
		TimeStampFormat:       "2006-01-02 15:04:05",
		CharStampFormat:       "",
		HideKeys:              false,
		Position:              true,
		Colors:                true,
		FieldsColors:          true,
		FieldsSpace:           true,
		ShowFullLevel:         false,
		LowerCaseLevel:        true,
		TrimMessages:          true,
		CallerFirst:           false,
		CustomCallerFormatter: nil,
	})
	conf.InitConfig()
	if conf.GetConfig().LogLevel == "info" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

}

func init() {
	lib.NewPlugin("echo", lib.OnlySelf()).OnCommand("echo").Handle(func(ctx *lib.Context) {
		event, err := ctx.GetEvent()
		if err != nil {
			log.Errorln("等待超时")
			return
		}
		log.Infoln("收到连续消息" + event.Message)
	})
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	config := conf.GetConfig()
	if err := lib.Init(ctx, config.Telegram.ApiId, config.Telegram.ApiHash, config.Telegram.ProxyURL); err != nil {
		panic(err)
	}
}
