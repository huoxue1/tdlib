package task

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/tdlib/lib"
	"github.com/huoxue1/tdlib/utils"
)

func init() {
	lib.NewPlugin("bean", lib.OnlySelf()).OnCommand("豆").Handle(func(ctx *lib.Context) {
		index := 1
		if len(ctx.Args) > 0 {
			i, _ := strconv.Atoi(ctx.Args[0])
			index = i
		}
		ql := qlMap[1].QL
		envs, err := ql.GetEnvs("JD_COOKIE")
		if err != nil {
			log.Errorln("获取cookie失败" + err.Error())
			return
		}

		detail, err := utils.GetJxBean(envs[index-1].Value)
		if err != nil {
			_ = ctx.EditMessage(err.Error())
			return
		}
		total := 0
		beans := make(map[string]int, 50)
		msg := "京豆总和：" + strconv.Itoa(detail.BeanNum) + ",今日总计： %d\n"
		for _, s := range detail.List {
			data, _ := time.Parse("2006-01-02 15:04:05", s.CreateDate)
			if data.Day() == time.Now().Day() {
				total += s.Amount
				if _, ok := beans[s.VisibleInfo]; ok {
					beans[s.VisibleInfo] += s.Amount
				} else {
					beans[s.VisibleInfo] = s.Amount
				}

			}
		}
		for s, i := range beans {
			msg += fmt.Sprintf("%v: %d\n", s, i)
		}
		_ = ctx.EditMessage(fmt.Sprintf(msg, total))
		time.Sleep(time.Second * 5)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
		_ = ctx.EditMessage("豆")
	})
}
