package task

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/tdlib"
	conf2 "github.com/huoxue1/tdlib/test/conf"
	"github.com/huoxue1/tdlib/test/utils"
)

var (
	qlMap = make(map[int]*QingLong, 5)
	tasks []*Task
)

type QingLong struct {
	ClientID     string `json:"Client_ID" mapstructure:"Client_ID"`
	ClientSecret string `json:"Client_Secret" mapstructure:"Client_Secret"`
	Url          string `json:"url" mapstructure:"url"`
	QL           *utils.Ql
}

type Task struct {
	Env      string
	KeyWords []string
	Script   string
	Name     string
	TimeOut  int
	Disable  bool
	cronID   map[int]int
	// 运行过的变量
	oldExport []string
	// 运行的通道
	ch chan int
	// 总共运行
	total int
	// 等待中的变量
	wait int
}

func init() {
	tdlib.OnConnect(func(ctx *tdlib.Context) {
		log.Infoln("tdlib已连接成功！")
		conf := conf2.GetConfig()
		for i, s := range conf.QingLong {
			ql := utils.InitQl(s.Url, s.ClientID, s.ClientSecret)
			if ql == nil {
				log.Errorln("初始化ql容器" + s.Url + "失败！！")
				continue
			}
			qlMap[i+1] = &QingLong{
				ClientID:     s.ClientID,
				ClientSecret: s.ClientSecret,
				Url:          s.Url,
				QL:           ql,
			}
			log.Infoln(fmt.Sprintf("初始化青龙%v成功", s.Url))
		}
		for _, s := range conf.JsConfig {
			t := &Task{
				Env:       s.Env,
				KeyWords:  s.KeyWord,
				Script:    s.Script,
				Name:      s.Name,
				TimeOut:   s.TimeOut,
				Disable:   s.Disable,
				cronID:    make(map[int]int, 5),
				ch:        make(chan int, 20),
				total:     0,
				wait:      0,
				oldExport: []string{},
			}
			for _, i := range s.Container {
				ql, ok := qlMap[i]
				if !ok {
					log.Errorln(fmt.Sprintf("青龙%d不存在，已跳过", i))
					continue
				}
				crons, err := ql.QL.GetCrons(s.Script)
				if err != nil {
					log.Errorln("从青龙获取定时任务失败" + err.Error())
					continue
				}
				for _, cron := range crons {
					paths := strings.Split(strings.TrimSpace(strings.TrimPrefix(cron.Command, "task")), "/")
					if strings.Join(paths, "/") == s.Script {
						log.Infoln(fmt.Sprintf("已成功初始化%v,定时id%d", s.Name, cron.Id))
						t.cronID[i] = cron.Id
						break
					}
					if paths[len(paths)-1] == s.Script {
						log.Infoln(fmt.Sprintf("已成功初始化%v,定时id%d", s.Name, cron.Id))
						t.cronID[i] = cron.Id
						break
					}
				}

			}
			tasks = append(tasks, t)

			go func(ctx2 *tdlib.Context, task *Task) {
				for {
					_ = <-task.ch
					task.wait--
					log.Infoln("开始执行任务" + task.Name)

					for i, id := range task.cronID {
						err := ctx2.SendChannelMsg(conf.Telegram.LogId, "开始执行任务"+task.Name, 0)
						if err != nil {
							log.Errorln("发送通知失败" + err.Error())
							return
						}
						ql := qlMap[i]
						err = ql.QL.RunCrons(id)
						if err != nil {
							log.Errorln("执行定时任务异常" + err.Error())
							continue
						}
						start := time.Now()
						c := make(chan int, 1)
						id := id
						go func() {
							for true {
								cron, err := ql.QL.GetCron(id)
								if err != nil {
									log.Errorln("获取定时任务状态异常" + err.Error())
									c <- 0
									return
								}
								if cron.Status == 1 {
									log.Infoln("任务执行结束")
									c <- 1
									return
								}
								time.Sleep(time.Second * 5)
							}
						}()

						after := time.After(time.Minute * 5)
						select {
						case <-after:
							log.Errorln("任务执行超时")
						case <-c:
							log.Infoln("任务执行完成")
						}
						err = ctx2.SendChannelMsg(conf.Telegram.LogId, fmt.Sprintf("%v任务执行完成，用时%.2f秒", task.Name, time.Now().Sub(start).Seconds()), 0)
						if err != nil {
							log.Errorln("发送通知失败" + err.Error())
							return
						}
					}
					time.Sleep(1 * time.Minute)
				}
			}(ctx, t)
		}
	})

	tdlib.NewPlugin("export", tdlib.OnlyChannels(conf2.GetConfig().Telegram.ListenCH...)).OnRegex(`export\s(.*?)="(.*?)"`).Handle(exportHandler)
	tdlib.NewPlugin("check_task", tdlib.OnlySelf()).OnCommand("check_task").Handle(func(ctx *tdlib.Context) {
		msg := ""
		for _, task := range tasks {
			msg += fmt.Sprintf("\n%v,%d/%d\n", task.Name, task.wait, task.total)
		}
		_ = ctx.EditMessage(msg)
		time.Sleep(time.Second * 7)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
	})
}

func exportHandler(ctx *tdlib.Context) {
	var exports []map[string]string
	msg := "检测到变量"
	for _, matcher := range ctx.RegexMatchers {
		msg += fmt.Sprintf("\n键: %v\n值：%v\n\n", matcher[1], matcher[2])
		exports = append(exports, map[string]string{
			"name":  matcher[1],
			"value": matcher[2],
		})
	}
	var matchTask *Task
	for _, task := range tasks {
		for _, export := range exports {
			for _, word := range task.KeyWords {
				if word == export["name"] {
					log.Infoln("匹配到任务" + task.Name)
					matchTask = task
				}
			}
		}
	}
	config := conf2.GetConfig()
	if matchTask == nil {
		if ctx.Self.ID == ctx.User.ID {
			msg += "\n自身发送，已替换\n"
			for _, ql := range qlMap {
				err := AddEnv(ql.QL, exports)
				if err != nil {
					log.Errorln("添加环境变量失败" + err.Error())
					return
				}
				log.Debugln("检测到自身发送，已替换变量")
			}
		}
		log.Debugln("未匹配到任务，已退出")
		msg += "未匹配到任务，已忽略"
		err := ctx.SendChannelMsg(config.Telegram.LogId, msg, 0)
		if err != nil {
			log.Errorln("发送通知失败" + err.Error())
			return
		}
		return
	} else {
		for i := range matchTask.cronID {
			err := AddEnv(qlMap[i].QL, exports)
			if err != nil {
				return
			}
		}
	}

	for _, s := range matchTask.oldExport {
		if s == exports[0]["name"] {
			msg += "旧的变量，已忽略"
			err := ctx.SendChannelMsg(config.Telegram.LogId, msg, 0)
			if err != nil {
				log.Errorln("发送通知失败" + err.Error())
				return
			}
			log.Debugln("旧的变量，已忽略")
			return
		}
	}

	matchTask.total++
	matchTask.wait++
	matchTask.oldExport = append(matchTask.oldExport, exports[0]["name"])
	matchTask.ch <- 0
	msg += "检测到任务" + matchTask.Name
	msg += fmt.Sprintf("\n等待中：%d,总共运行：%d", matchTask.wait, matchTask.total)
	err := ctx.SendChannelMsg(config.Telegram.LogId, msg, 0)
	if err != nil {
		log.Errorln("发送通知失败" + err.Error())
		return
	}
}

func AddEnv(ql *utils.Ql, data []map[string]string) error {
	file, err := ql.GetConfigFile("config.sh")
	if err != nil {
		return err
	}
	if strings.Contains(file, fmt.Sprintf(`export %v=`, data[0]["name"])) {
		log.Debugln("已存在该变量，开始替换")
		for _, d := range data {
			file = string(regexp.MustCompile(fmt.Sprintf(`export\s%s="(.*?)"`, d["name"])).ReplaceAll([]byte(file), []byte(fmt.Sprintf(`export %v="%v"`, d["name"], d["value"]))))
		}
	} else {
		log.Debugln("未检测到变量，开始添加变量")
		for _, d := range data {
			file += fmt.Sprintf(`%vexport %s="%s"`, "\n\n", d["name"], d["value"])
		}
	}
	err = ql.SaveConfigFile("config.sh", file)
	if err != nil {
		log.Errorln("保存变量失败" + err.Error())
		return err
	}
	return err
}
