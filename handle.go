package tdlib

import (
	"regexp"
	"strings"
	"sync"

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

func (m *Matcher) OnCommand(command string) *Matcher {
	m.Rules = append(m.Rules, func(ctx *Context) bool {
		text := ctx.Text
		s := strings.Split(text, " ")
		if s[0] == command {
			ctx.Args = s[1:]
			return true
		} else {
			return false
		}
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
