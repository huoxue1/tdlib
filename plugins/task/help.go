package task

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	conf2 "github.com/huoxue1/tdlib/conf"
	"github.com/huoxue1/tdlib/lib"
)

func init() {
	lib.NewPlugin("help").OnCommand("help").Handle(func(ctx *lib.Context) {
		_ = ctx.EditMessage(fmt.Sprintf("check_task\nenable_task\ndisable_task\nsearch_cron\nrun_cron\n豆\nid\nhelp"))
		time.Sleep(5 * time.Second)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
		_ = ctx.EditMessage("help")
	})
	lib.NewPlugin("restart").OnCommand("restart", "tdlib重启").Handle(func(ctx *lib.Context) {
		_ = ctx.EditMessage("开始重启")
		Restart()
	})

}

// Restart
/* @Description:
 */
func Restart() {
	once := sync.Once{}
	once.Do(func() {
		log.Infoln("程序启动命令： " + conf2.GetConfig().RestartCmd)
		cmd := exec.Command(conf2.GetConfig().RestartCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Start()
		os.Exit(3)
	})

}
