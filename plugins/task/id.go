package task

import (
	"fmt"
	"time"

	"github.com/gotd/td/tg"

	"github.com/huoxue1/tdlib/lib"
)

func init() {
	lib.NewPlugin("id", lib.OnlySelf()).OnCommand("id").Handle(func(ctx *lib.Context) {
		if ctx.Message.ReplyTo.ReplyToMsgID != 0 {
			msg := ""
			message := ctx.GetMsg(ctx.Channel.ID, ctx.Message.ReplyTo.ReplyToMsgID).(*tg.MessagesChannelMessages)
			msg += fmt.Sprintf("userId: %v\n", message.Users[0].(*tg.User).ID)
			msg += fmt.Sprintf("channelId: %v\n", lib.IDConveryChannelID(message.Chats[0].(*tg.Channel).ID))
			if len(message.Users) > 1 {
				msg += fmt.Sprintf("forwardFrom：\n   userId: %v\n", message.Users[len(message.Users)-1].(*tg.User).ID)
			}
			if len(message.Chats) > 1 {
				msg += fmt.Sprintf("   channelId: %v\n", lib.IDConveryChannelID(message.Chats[len(message.Users)-1].(*tg.Channel).ID))
			}
			_ = ctx.EditMessage(msg)
			time.Sleep(10 * time.Second)
			ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
			_ = ctx.EditMessage("id")
		}
	})
}
