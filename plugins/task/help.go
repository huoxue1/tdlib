package task

import (
	"fmt"
	"time"

	"github.com/huoxue1/tdlib/lib"
)

func init() {
	lib.NewPlugin("help").OnCommand("help").Handle(func(ctx *lib.Context) {
		_ = ctx.EditMessage(fmt.Sprintf("check_task\nenable_task\ndisable_task\nsearch_cron\nrun_cron\n豆\nid\nhelp"))
		time.Sleep(5 * time.Second)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
		_ = ctx.EditMessage("help")
	})
}
