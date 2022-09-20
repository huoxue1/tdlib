package lib

import (
	"context"
	"time"

	"github.com/gotd/td/tg"
)

var channels = make(map[int64]int64, 10)

type Context struct {
	ctx context.Context

	Client *tg.Client

	Self *tg.User

	Message *tg.Message

	// 消息类型 user channel
	MessageType string

	Channel *tg.Channel

	User *tg.User

	MsgID int

	Text string

	Args []string
	// 正则匹配结果
	RegexMatchers [][]string
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) Value(key any) any {
	return c.ctx.Value(key)
}
