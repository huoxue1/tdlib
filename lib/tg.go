package lib

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/updates"
	updhook "github.com/gotd/td/telegram/updates/hook"
	"github.com/gotd/td/tg"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/net/proxy"
)

// noSignUp can be embedded to prevent signing up.
type noSignUp struct{}

func (c noSignUp) SignUp(ctx context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, errors.New("not implemented")
}

func (c noSignUp) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return &auth.SignUpRequired{TermsOfService: tos}
}

// termAuth implements authentication via terminal.
type termAuth struct {
	noSignUp
}

func (a termAuth) Phone(_ context.Context) (string, error) {
	fmt.Print("Enter Phone: ")
	code, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(code), nil
}

func (a termAuth) Password(_ context.Context) (string, error) {
	fmt.Print("Enter 2FA password: ")
	bytePwd, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytePwd)), nil
}

func (a termAuth) Code(_ context.Context, _ *tg.AuthSentCode) (string, error) {
	fmt.Print("Enter code: ")
	code, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(code), nil
}

type Config struct {
	AppID   int
	AppHash string

	ProxyAddr string
}

func Init(ctx context.Context, appID int, appHash string, proxy string) error {

	handlerMap.Range(func(key, value any) bool {
		log.Infoln("已加载插件 ==》 " + key.(string))
		return true
	})

	l, _ := zap.NewDevelopment(zap.IncreaseLevel(zapcore.InfoLevel), zap.AddStacktrace(zapcore.FatalLevel))
	dispatcher := tg.NewUpdateDispatcher()
	gaps := updates.New(updates.Config{
		Handler: dispatcher,
		Logger:  l.Named("gaps"),
	})

	dispatcher.OnNewChannelMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {

		var ch *tg.Channel
		var user *tg.User

		for _, u := range e.Users {
			user = u
		}

		for _, channel := range e.Channels {
			ch = channel
		}

		msg, ok := update.Message.(*tg.Message)
		if !ok {
			return nil
		}

		client := ctx.Value("client").(*telegram.Client)
		self, _ := client.Self(ctx)
		c := new(Context)
		c.ctx = ctx
		c.Self = self
		c.Message = msg
		c.Client = client.API()
		c.Channel = ch
		c.Text = msg.Message
		c.User = user
		c.MessageType = MESSAGETYPECHANNEL
		c.MsgID = msg.ID

		channels[c.Channel.ID] = c.Channel.AccessHash
		handlerMap.Range(func(key, value any) bool {
			defer func() {
				_ = recover()
			}()
			handle := value.(*Matcher)
			for _, rule := range handle.Rules {
				handleRule := func(rule2 Rule) bool {
					defer func() {
						err := recover()
						if err != nil {
							log.Errorln("处理事件过程异常")
							log.Errorln(err)
						}
					}()
					return rule2(c)
				}
				if !handleRule(rule) {
					return true
				}
			}
			go func() {
				defer func() {
					err := recover()
					if err != nil {
						log.Errorln("处理事件过程异常")
						log.Errorln(err)
					}
				}()
				log.Infoln("handle the matcher " + key.(string))
				handle.Handler(c)
			}()
			return false
		})

		log.Infoln("收到消息" + msg.Message)
		return nil
	})
	dispatcher.OnNewMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
		client := ctx.Value("client").(*telegram.Client)

		msg, ok := update.Message.(*tg.Message)
		if !ok {
			return nil
		}
		self, _ := client.Self(ctx)
		if msg.FromID == nil {
			msg.FromID = &tg.PeerUser{UserID: self.ID}
		}
		c := new(Context)
		c.ctx = ctx
		c.Self = self
		c.Message = msg
		c.Client = client.API()
		c.Channel = &tg.Channel{ID: 0}
		c.Text = msg.Message
		c.User = &tg.User{ID: (msg.FromID.(*tg.PeerUser)).UserID}
		c.MessageType = MESSAGETYPEUSER
		c.MsgID = msg.ID
		handlerMap.Range(func(key, value any) bool {
			handle := value.(*Matcher)
			for _, rule := range handle.Rules {
				handleRule := func(rule2 Rule) bool {
					defer func() {
						err := recover()
						if err != nil {
							log.Errorln("处理事件过程异常")
							log.Errorln(err)
						}
					}()
					return rule2(c)
				}
				if !handleRule(rule) {
					return true
				}
			}
			go func() {
				defer func() {
					err := recover()
					if err != nil {
						log.Errorln("处理事件过程异常")
						log.Errorln(err)
					}
				}()
				log.Infoln("handle the matcher " + key.(string))
				handle.Handler(c)
			}()
			return false
		})

		log.Infoln("收到消息" + msg.Message)
		return nil
	})

	ConnectTelegram(ctx, &Config{
		AppID:     appID,
		AppHash:   appHash,
		ProxyAddr: proxy,
	}, gaps)
	return nil
}

func ConnectTelegram(ctx context.Context, config *Config, manager *updates.Manager) {

	db, err := InitDB()
	if err != nil {
		log.Errorln("打开db文件错误" + err.Error())
		return
	}
	defer db.Close()
	proxyUrl, err := url.Parse(config.ProxyAddr)
	if err != nil {
		log.Errorln("解析代理地址失败")
		return
	}
	sock5, _ := proxy.FromURL(proxyUrl, proxy.Direct)

	if sock5 == nil {
		log.Warningln("未配置代理，使用环境变量！")
		sock5 = proxy.FromEnvironmentUsing(proxy.Direct)
	}
	dc := sock5.(interface {
		DialContext(ctx context.Context, network, addr string) (net.Conn, error)
	})
	l, _ := zap.NewDevelopment(zap.IncreaseLevel(zapcore.InfoLevel), zap.AddStacktrace(zapcore.FatalLevel))
	client := telegram.NewClient(config.AppID, config.AppHash, telegram.Options{
		SessionStorage: &MyStore{Db: db},
		DC:             2,
		DialTimeout:    time.Minute * 5,
		Logger:         l,
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dc.DialContext,
		}),
		Middlewares: []telegram.Middleware{
			updhook.UpdateHook(manager.Handle),
		},
		UpdateHandler: manager,
	})
	err = client.Run(ctx, func(ctx context.Context) error {
		err := client.Auth().IfNecessary(ctx, auth.NewFlow(&termAuth{}, auth.SendCodeOptions{}))
		if err != nil {
			log.Errorln("鉴权失败" + err.Error())
			return err
		}
		user, err := client.Self(ctx)
		if err != nil {
			log.Errorln(err.Error())
			return err
		}
		_ = db.Store("self_id", strconv.FormatInt(user.ID, 10))
		log.Infoln(fmt.Sprintf("%v已登陆", user.Username))
		// Notify update manager about authentication.

		// 处理连接事件
		for _, handler := range connectHandlers {
			h := handler
			go func() {
				defer func() {
					_ = recover()
				}()
				h(&Context{
					ctx:    ctx,
					Client: client.API(),
					Text:   "",
				})
			}()
		}

		if err := manager.Auth(context.WithValue(ctx, "client", client), client.API(), user.ID, user.Bot, true); err != nil {
			log.Errorln(err.Error())
			return err
		}
		defer func() { _ = manager.Logout() }()

		<-ctx.Done()
		return ctx.Err()
	})
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}

type MyStore struct {
	Db *MyDB
}

func (m *MyStore) LoadSession(ctx context.Context) ([]byte, error) {
	data, err := m.Db.Load("tg_session")
	if err != nil {
		return []byte(""), nil
	}
	return []byte(data), err
}

func (m *MyStore) StoreSession(ctx context.Context, data []byte) error {
	return m.Db.Store("tg_session", regexp.MustCompile(`(,"ReactionsDefault":\{.*?})`).ReplaceAllString(string(data), ""))
	//return m.Db.Store("tg_session", strings.ReplaceAll(string(data), `,"ReactionsDefault":{"Emoticon":"👍"}`, ""))
}
