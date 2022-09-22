package task

import (
	"fmt"
	"strconv"
	"time"

	"github.com/huoxue1/tdlib/lib"
)

func init() {
	lib.NewPlugin("search_cron", lib.OnlySelf()).OnCommand("search_cron", "定时任务搜索").Handle(func(ctx *lib.Context) {
		ql := qlMap[1]
		crons, err := ql.QL.GetCrons(ctx.Args[0])
		if err != nil {
			return
		}
		getEmoji := func(bool2 bool) string {
			if !bool2 {
				return "🉑"
			} else {
				return "🚫"
			}
		}
		getEmoji1 := func(bool2 bool) string {
			if !bool2 {
				return "🏃‍♀️"
			} else {
				return ""
			}
		}
		msg := "搜索结果：\n"
		for _, cron := range crons {
			msg += fmt.Sprintf("\nid：%v\n名称：%v%v%v\n命令：%v\n", cron.Id, getEmoji(cron.IsDisabled == 1), getEmoji1(cron.Status == 1), cron.Name, cron.Command)
		}
		_ = ctx.EditMessage(msg)
		time.Sleep(10 * time.Second)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
		_ = ctx.EditMessage("search_cron")
	})

	lib.NewPlugin("run_cron", lib.OnlySelf()).OnCommand("run_cron", "执行定时任务").Handle(func(ctx *lib.Context) {
		ql := qlMap[1]
		id, err := strconv.Atoi(ctx.Args[0])
		if err != nil {
			return
		}
		err = ql.QL.RunCrons(id)
		if err != nil {
			return
		}
		_ = ctx.EditMessage("执行任务成功")
		time.Sleep(10 * time.Second)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
		_ = ctx.EditMessage("run_cron")
	})
}
