package lib

import (
	"errors"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gotd/td/tg"
	log "github.com/sirupsen/logrus"
)

var (
	connectHandlers []func(ctx *Context)
	handlerMap      sync.Map
)

const (
	MESSAGETYPECHANNEL = "channel"
	MESSAGETYPEUSER    = "user"
)

type Rule func(ctx *Context) bool

type Handler func(ctx *Context)

type Matcher struct {
	Name    string
	Rules   []Rule
	Handler Handler
	Block   bool
	Disable bool
}

func NewPlugin(name string, rules ...Rule) *Matcher {
	m := new(Matcher)
	m.Name = name
	m.Rules = append(m.Rules, rules...)
	return m
}

func (m *Matcher) R(rule Rule) *Matcher {
	m.Rules = append(m.Rules, rule)
	return m
}

func (m *Matcher) Dis() *Matcher {
	m.Disable = true
	return m
}

func (m *Matcher) SetBlock(block bool) {
	m.Block = block
}

func (m *Matcher) Handle(h Handler) {
	m.Handler = h
	handlerMap.Store(m.Name, m)
}

func (m *Matcher) OnMessage() *Matcher {
	return m
}

func (m *Matcher) OnFullMatch(message string) *Matcher {
	m.Rules = append(m.Rules, func(ctx *Context) bool {
		if ctx.Text == message {
			return true
		} else {
			return false
		}
	})
	return m
}

func (m *Matcher) OnCommand(command ...string) *Matcher {
	m.Rules = append(m.Rules, func(ctx *Context) bool {
		text := ctx.Text
		s := strings.Split(text, " ")
		for _, s2 := range command {
			if s[0] == s2 {
				ctx.Args = s[1:]
				return true
			}
		}
		return false
	})
	return m
}

func (m *Matcher) OnRegex(regexStr string) *Matcher {
	m.Rules = append(m.Rules, func(ctx *Context) bool {
		text := ctx.Text
		r, err := regexp.Compile(regexStr)
		if err != nil {
			log.Errorln("错误的正则表达式" + err.Error())
		}
		s := r.FindAllStringSubmatch(text, -1)
		if len(s) < 1 {
			return false
		}
		ctx.RegexMatchers = s
		return true
	})
	return m
}

// OnConnect
/* @Description: 连接时的生命周期
*  @param f
 */
func OnConnect(f func(ctx *Context)) {
	connectHandlers = append(connectHandlers, f)
}

func (ctx *Context) newTempMatcher(matcher *Matcher) func() {
	s := uuid.New().String()
	matcher.Name = s
	handlerMap.Store(s, matcher)
	return func() {
		handlerMap.Delete(s)
	}
}

func (ctx *Context) GetEvent() (*tg.Message, error) {
	ch := make(chan *tg.Message, 1)
	m := &Matcher{Rules: []Rule{
		func(c *Context) bool {
			if c.Channel.ID == ctx.Channel.ID && c.User.ID == ctx.User.ID && c.MessageType == ctx.MessageType {
				return true
			} else {
				return false
			}
		},
	},
		Handler: func(c *Context) {
			ch <- c.Message
		},
	}
	cancel := ctx.newTempMatcher(m)
	after := time.After(time.Minute)
	select {
	case data := <-ch:
		cancel()
		return data, nil
	case <-after:
		log.Errorln("等待上下文超时")
		cancel()
		return nil, errors.New("wait message time out")
	}
}

func (ctx *Context) GetEventChan() (chan *tg.Message, func()) {
	ch := make(chan *tg.Message, 1)
	m := &Matcher{Rules: []Rule{
		func(c *Context) bool {
			if c.Channel.ID == ctx.Channel.ID && c.User.ID == ctx.User.ID && c.MessageType == ctx.MessageType {
				return true
			} else {
				return false
			}
		},
	},
		Handler: func(c *Context) {
			ch <- c.Message
		},
	}
	cancel := ctx.newTempMatcher(m)
	return ch, func() {
		cancel()
		close(ch)
	}
}
