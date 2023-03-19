package nolan

//
//import (
//	"errors"
//	"github.com/huoxue1/tdlib/utils/db"
//	"regexp"
//	"strconv"
//	"time"
//
//	"github.com/imroc/req/v3"
//	log "github.com/sirupsen/logrus"
//	"github.com/tidwall/gjson"
//
//	"github.com/huoxue1/tdlib/lib"
//)
//
///*
//	nolan_api  nolan的api
//	global_nolan_enable  1 启用  其他 禁用
//	nolan_enable_channel   nolan允许使用的频道
//	nolan_disable_user     禁用使用nolan的用户
//
//	nolan_target nolan登录的提示信息
//*/
//
//func init() {
//	lib.NewPlugin("nolan", checkGlobal, checkChannelAndUser).OnCommand("nark_login").Handle(nolan)
//	lib.NewPlugin("nolan_enable", lib.OnlySelf()).OnCommand("enable_nolan", "启用nark").Handle(enableNolan)
//	lib.NewPlugin("nolan_disable", lib.OnlyChannel(), lib.OnlySelf()).OnCommand("disable_nolan", "禁用nark").Handle(disableNolan)
//}
//
//func disableNolan(ctx *lib.Context) {
//	c := db.GetRedisClient()
//	_ = c.SRem("nolan_enable_channel", strconv.FormatInt(ctx.Channel.ID, 10))
//	_ = ctx.EditMessage("🈲本群已成功禁用nark登录")
//
//}
//
//func enableNolan(ctx *lib.Context) {
//	c := db.GetRedisClient()
//	_ = c.SAdd("nolan_enable_channel", strconv.FormatInt(ctx.Channel.ID, 10))
//	_ = ctx.EditMessage("🈲本群已成功禁用nark登录")
//	_ = ctx.EditMessage("本群已成功启用nark登录，发送nark_login即可登录")
//
//}
//
//func checkGlobal(ctx *lib.Context) bool {
//	c := db.GetRedisClient()
//	// 检查是否全局禁用
//	globalNolanEnable, _ := c.Get("global_nolan_enable").Result()
//	if globalNolanEnable != "1" {
//		return false
//	}
//	return true
//}
//
//func checkChannelAndUser(ctx *lib.Context) bool {
//	c := db.GetRedisClient()
//
//	// 检查是否在白名单群内
//	if ctx.MessageType == lib.MESSAGETYPECHANNEL {
//		result, err := c.SMembers("c := db.GetRedisClient()").Result()
//		if err != nil {
//			return false
//		} else {
//			for _, s := range result {
//				if s == strconv.FormatInt(ctx.Channel.ID, 10) {
//					return true
//				}
//			}
//		}
//	}
//
//	// 检查是否属于黑名单
//	//nolanDisableUser, _ := db.Load("nolan_disable_user")
//	//if nolanDisableUser == "" {
//	//	return true
//	//} else {
//	//	for _, user := range strings.Split(nolanDisableUser, ",") {
//	//		if user == strconv.FormatInt(ctx.User.ID, 10) {
//	//			return false
//	//		}
//	//	}
//	//}
//	return false
//}
//
//func nolan(ctx *lib.Context) {
//	c := db.GetRedisClient()
//	nolanApi, err := c.Get("nolan_api").Result()
//	if err != nil {
//		_ = ctx.SendText("nolan_api未设置！", 0)
//		return
//	}
//	s := c.Get("nolan_target").String()
//	if s == "" {
//		s = "nark为你服务，请输入11位手机号： "
//	}
//	err = ctx.SendText(s, ctx.MsgID)
//	if err != nil {
//		return
//	}
//	eventChan, cancel := ctx.GetEventChan()
//	defer cancel()
//	after := time.After(time.Minute)
//	var phone string
//	select {
//	case <-after:
//		return
//	case msg := <-eventChan:
//		if msg.Message == "q" {
//			return
//		}
//		if !checkMobile(msg.Message) {
//			_ = ctx.SendText("不合法的手机号", msg.ID)
//			return
//		}
//		err := sendSMS(nolanApi, msg.Message)
//		if err != nil {
//			_ = ctx.SendText("nolan出现异常"+err.Error(), msg.ID)
//			return
//		}
//		phone = msg.Message
//		_ = ctx.SendText("请输入六位验证码：", msg.ID)
//	}
//	after = time.After(time.Minute)
//	select {
//	case <-after:
//		return
//	case msg := <-eventChan:
//		if msg.Message == "q" {
//			return
//		}
//		if len(msg.Message) != 6 {
//			return
//		}
//		nick, err := veirfyCode(nolanApi, phone, msg.Message, strconv.FormatInt(ctx.User.ID, 10))
//		if err != nil {
//			_ = ctx.SendText("nolan出现异常"+err.Error(), msg.ID)
//			return
//		}
//		_ = ctx.SendText("登录成功！,备注信息为："+nick, msg.ID)
//	}
//
//}
//
//// checkMobile 检验手机号
//func checkMobile(phone string) bool {
//	// 匹配规则
//	// ^1第一位为一
//	// [345789]{1} 后接一位345789 的数字
//	// \\d \d的转义 表示数字 {9} 接9位
//	// $ 结束符
//	regRuler := "^1[345789]{1}\\d{9}$"
//
//	// 正则调用规则
//	reg := regexp.MustCompile(regRuler)
//
//	// 返回 MatchString 是否匹配
//	return reg.MatchString(phone)
//
//}
//
//// sendSMS
///* @Description: 发送验证码
//*  @param nolanApi
//*  @param phone
//*  @return error
// */
//func sendSMS(nolanApi string, phone string) error {
//	response, err := req.C().R().SetBodyJsonMarshal(map[string]any{
//		"Phone": phone,
//		"qlkey": 1,
//	}).SetHeaders(map[string]string{
//		"content-type": "application/json",
//		"accept":       "application/json",
//	}).Post(nolanApi + "/api/SendSMS")
//	if err != nil {
//		return err
//	}
//	log.Infoln(response.String())
//	if gjson.GetBytes(response.Bytes(), "success").Bool() {
//		return nil
//	} else {
//		return errors.New(response.String())
//	}
//}
//
//// veirfyCode
///* @Description: 验证码校验
//*  @param nolanApi
//*  @param phone
//*  @param code
//*  @param user_id
//*  @return string
//*  @return error
// */
//func veirfyCode(nolanApi string, phone string, code string, user_id string) (string, error) {
//	response, err := req.C().R().SetBodyJsonMarshal(map[string]any{
//		"Phone": phone,
//		"Code":  code,
//		"QQ":    user_id,
//		"qlkey": 1,
//	}).SetHeaders(map[string]string{
//		"content-type": "application/json",
//		"accept":       "application/json",
//	}).Post(nolanApi + "/api/VerifyCode")
//	if err != nil {
//		return "", err
//	}
//	log.Infoln(response.String())
//	if gjson.GetBytes(response.Bytes(), "success").Bool() {
//		return gjson.GetBytes(response.Bytes(), "data.nickname").String(), nil
//	} else {
//		return "", errors.New(response.String())
//	}
//}
