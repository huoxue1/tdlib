package main

import (
	"context"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	nested "github.com/Lyrics-you/sail-logrus-formatter/sailor"

	"github.com/huoxue1/tdlib"
	"github.com/huoxue1/tdlib/test/conf"

	_ "github.com/huoxue1/tdlib/test/plugins/task"
)

func init() {
	conf.InitConfig()
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
	log.SetLevel(log.DebugLevel)
}

func init() {
	tdlib.NewPlugin("test", tdlib.OnlySelf()).OnCommand("test1").Handle(func(ctx *tdlib.Context) {
		log.Infoln("测试插件成功")
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
	})
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	config := conf.GetConfig()
	if err := tdlib.Init(ctx, config.Telegram.ApiId, config.Telegram.ApiHash, config.Telegram.ProxyURL); err != nil {
		panic(err)
	}
}
