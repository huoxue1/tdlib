package task

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	conf2 "github.com/huoxue1/tdlib/conf"
	"github.com/huoxue1/tdlib/lib"
	"github.com/huoxue1/tdlib/utils"
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
	lib.OnConnect(connectHandler)
	lib.NewPlugin("export", lib.OnlyChannels(conf2.GetConfig().Telegram.ListenCH...)).OnRegex(`export\s(.*?)="(.*?)"`).Handle(exportHandler)
	lib.NewPlugin("check_task", lib.OnlySelf()).OnCommand("check_task", "检查任务").Handle(checkTaskHandler)
	lib.NewPlugin("disable_task", lib.OnlySelf()).OnCommand("disable_task", "禁用任务").Handle(disableTaskHandler)
	lib.NewPlugin("enable_task", lib.OnlySelf()).OnCommand("disable_task", "启用任务").Handle(enableTaskHandler)
	lib.NewPlugin("add_task", lib.OnlySelf()).OnCommand("add_task", "添加任务").Handle(addTaskHandler)
}

func addTaskHandler(ctx *lib.Context) {
	propomts := []string{
		"请输入任务名称", "请输入变量名称", "请输入关键字", "请输入脚本名称",
	}
	var values []string
	for _, propomt := range propomts {
		_ = ctx.EditMessage(propomt)
		message, err := ctx.GetEvent()
		if err != nil {
			_ = ctx.EditMessage("等待消息超时" + err.Error())
			return
		}
		values = append(values, message.Message)
	}
	task := new(conf2.TaskConfig)
	task.Name = values[0]
	task.Env = values[1]
	task.KeyWord = strings.Split(values[2], " ")
	task.Script = values[3]
	task.Container = []int{1}
	taskConfigs := conf2.GetConfig().JsConfig
	taskConfigs = append(taskConfigs, task)
	err := conf2.SaveTask(taskConfigs)
	if err != nil {
		log.Errorln("保存数据失败")
		return
	}
	t := &Task{
		Env:       task.Env,
		KeyWords:  task.KeyWord,
		Script:    task.Script,
		Name:      task.Name,
		TimeOut:   0,
		Disable:   false,
		cronID:    make(map[int]int, 5),
		ch:        make(chan int, 20),
		total:     0,
		wait:      0,
		oldExport: []string{},
	}
	tasks = append(tasks, t)
	go runTask(ctx, t)
	_ = ctx.EditMessage("成功添加任务" + t.Name)
}

func enableTaskHandler(ctx *lib.Context) {
	if len(ctx.Args) < 1 {
		log.Errorln("参数不足")
		return
	}
	id, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		log.Errorln("转化参数错误")
		return
	}
	task := tasks[id-1]

	log.Infoln("即将启用task" + task.Name)
	task.Disable = false
	db, _ := lib.InitDB()
	_ = db.Delete(task.Name + "_disable")
	_ = ctx.EditMessage("已成功启用任务" + task.Name)
	time.Sleep(10 * time.Second)
	ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
	_ = ctx.EditMessage("enable")

}

func disableTaskHandler(ctx *lib.Context) {
	if len(ctx.Args) < 1 {
		return
	}
	id, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return
	}
	task := tasks[id-1]

	log.Infoln("即将禁用task" + task.Name)
	task.Disable = true
	db, _ := lib.InitDB()
	err = db.Store(task.Name+"_disable", "true")
	if err != nil {
		return
	}
	_ = ctx.EditMessage("已成功禁用任务" + task.Name)
	time.Sleep(10 * time.Second)
	ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
	_ = ctx.EditMessage("disable")

}

func checkTaskHandler(ctx *lib.Context) {
	msg := ""
	getEmoji := func(bool2 bool) string {
		if !bool2 {
			return "🉑"
		} else {
			return "🚫"
		}
	}
	for i, task := range tasks {
		msg += fmt.Sprintf("\n%v %d：%v,%d/%d\n", getEmoji(task.Disable), i+1, task.Name, task.wait, task.total)
	}
	_ = ctx.EditMessage(msg)
	time.Sleep(time.Second * 7)
	ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
}

// connectHandler
/* @Description: 初始化task
*  @param ctx
 */
func connectHandler(ctx *lib.Context) {
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
	db, _ := lib.InitDB()
	for _, s := range conf.JsConfig {
		t := &Task{
			Env:       s.Env,
			KeyWords:  s.KeyWord,
			Script:    s.Script,
			Name:      s.Name,
			TimeOut:   s.TimeOut,
			Disable:   s.Disable == 1,
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

		load, _ := db.Load(t.Name + "_disable")
		if load != "" {
			t.Disable = true
			log.Warningln("已禁用任务 " + t.Name)
		}
		go runTask(ctx, t)
	}
}

// runTask
/* @Description: 循环等待任务
*  @param ctx2
*  @param task
 */
func runTask(ctx2 *lib.Context, task *Task) {
	conf := conf2.GetConfig()
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
}

func exportHandler(ctx *lib.Context) {
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
	if matchTask.Disable {
		msg += "检测到任务" + matchTask.Name + "\n任务已被禁用"
		err := ctx.SendChannelMsg(config.Telegram.LogId, msg, 0)
		if err != nil {
			log.Errorln("发送通知失败" + err.Error())
			return
		}
		return
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
