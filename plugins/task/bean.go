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
		detail, err := utils.GetBeanDetail(envs[index-1].Value)
		if err != nil {
			log.Errorln("获取京豆详情失败" + err.Error())
			return
		}
		total := 0
		for _, d := range detail {
			dou, _ := strconv.Atoi(d.Amount)
			total += dou
		}
		bean, err := utils.TotalBean(envs[index-1].Value)
		if err != nil {
			log.Errorln("获取京豆总和失败")
			return
		}
		log.Debugln(*bean)
		msg := bean.Data.UserInfo.BaseInfo.Nickname + " 京豆详情：\n"
		msg += fmt.Sprintf("京东总和：%v,今日收入：%d\n", bean.Data.AssetInfo.BeanNum, total)
		for _, d := range detail {
			msg += fmt.Sprintf("\n%s: %v豆", d.EventMassage, d.Amount)
		}
		_ = ctx.EditMessage(msg)
		time.Sleep(time.Second * 5)
		ctx.DeleteMsg(ctx.Message.Flags, ctx.Channel.ID, ctx.MsgID)
	})
}
