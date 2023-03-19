package param

import (
	"fmt"
	"github.com/huoxue1/tdlib/utils/db"
	"time"

	"github.com/huoxue1/tdlib/lib"
)

func init() {
	lib.NewPlugin("set_param", lib.OnlySelf()).OnCommand("set_tdlib_param", "设置运行参数").Handle(func(ctx *lib.Context) {
		if len(ctx.Args) < 2 {
			_ = ctx.EditMessage("缺少参数！！")
			return
		}
		c := db.GetCache()
		err := c.Set(ctx.Args[0], ctx.Args[1])
		if err != nil {
			_ = ctx.EditMessage("缺少参数！！")
			return
		}
		_ = ctx.EditMessage("设置参数" + ctx.Args[0] + "成功！")
		time.Sleep(5 * time.Second)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
	})

	lib.NewPlugin("get_param", lib.OnlySelf()).OnCommand("get_tdlib_param", "获取运行参数").Handle(func(ctx *lib.Context) {
		if len(ctx.Args) < 1 {
			_ = ctx.EditMessage("缺少参数！！")
			return
		}
		c := db.GetCache()
		data := c.Get(ctx.Args[0])
		if data == "" {
			ctx.EditMessage(fmt.Sprintf("键%v不存在", ctx.Args[0]))
		} else {
			_ = ctx.EditMessage(fmt.Sprintf("键： %v\n\n值：%v", ctx.Args[0], data))
		}
		time.Sleep(5 * time.Second)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
	})

}
