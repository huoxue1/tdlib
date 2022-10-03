package main

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	nested "github.com/Lyrics-you/sail-logrus-formatter/sailor"
	"github.com/huoxue1/xdaemon"
	log "github.com/sirupsen/logrus"

	_ "github.com/huoxue1/tdlib/plugins/task"

	"github.com/huoxue1/tdlib/conf"
	"github.com/huoxue1/tdlib/lib"
)

func init() {
	runBack()
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

func runBack() {
	cmd, err := xdaemon.Background(os.Stdout, false)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if xdaemon.IsParent() {
		go onKill(cmd)
		for true {
			_ = cmd.Wait()
			if cmd.ProcessState.Exited() {
				if cmd.ProcessState.ExitCode() != 1001 {
					break
				}
			}
			cmd, err = xdaemon.Background(os.Stdout, false)
			if err != nil {
				return
			}
		}
		os.Exit(0)
	}
}

func onKill(cmd *exec.Cmd) {
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	if cmd.Process != nil {
		cmd.Process.Kill()
	}
	os.Exit(1)
}
