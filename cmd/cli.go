package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	nested "github.com/Lyrics-you/sail-logrus-formatter/sailor"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/huoxue1/tdlib/cmd/update"
	"github.com/huoxue1/tdlib/conf"
)

func main() {
	version, err := update.LastVersion()
	if err != nil {
		log.Errorln("检查最新版本失败" + err.Error())
	}
	if checkFile("tdlib") || checkFile("tdlib.exe") {
		log.Infoln("检测到tdlib文件已经存在，是否跳过下载?(y or n)")
		if !checkYesOrNo(getCmdLine()) {
			err := update.SelfUpdate("", version)
			if err != nil {
				return
			}
		}
		log.Infoln("已跳过下载！")
	} else {
		err := update.SelfUpdate("", version)
		if err != nil {
			return
		}
	}
	log.Infoln("是否生成配置文件？ (y or n)")
	if checkYesOrNo(getCmdLine()) {
		generateConfig()
	}

	log.Infoln("是否生成默认的监控变量文件？ (y or n)")
	if checkYesOrNo(getCmdLine()) {
		err := os.WriteFile("task.yaml", []byte(task), 0666)
		if err != nil {
			log.Errorln("写入监控变量文件失败" + err.Error())
			return
		}
	}

}

func generateConfig() {
	c := new(conf.Config)
	c.LogLevel = "info"
	log.Infoln("请输入你的青龙访问地址,例如： http://127.0.0.1:5701")
	qlAdd := getCmdLine()
	log.Infoln("请输入你的青龙应用的client_id")
	qlClientID := getCmdLine()
	log.Infoln("请输入你的青龙应用的client_secret")
	qlClientSecret := getCmdLine()
	c.QingLong = append(c.QingLong, conf.QingLongConf{
		ClientID:     qlClientID,
		ClientSecret: qlClientSecret,
		Url:          qlAdd,
	})
	log.Infoln("请输入你的telegram应用的app_id: ")
	tgApiID, _ := strconv.Atoi(getCmdLine())
	log.Infoln("请输入你的telegram应用的app_hash: ")
	tgApiHash := getCmdLine()
	log.Infoln("请输入你的telegram日志发送的群组id: ")
	logChanID, _ := strconv.ParseInt(getCmdLine(), 10, 64)
	log.Infoln("请输入你的telegram代理地址,例如： socks5://username:password@127.0.0.1:7890 ")
	tgProxy := getCmdLine()
	c.Telegram.ApiId = tgApiID
	c.Telegram.ApiHash = tgApiHash
	c.Telegram.LogId = logChanID
	c.Telegram.ProxyURL = tgProxy
	log.Infoln("已添加默认监控频道！")
	c.Telegram.ListenCH = append(c.Telegram.ListenCH, -1001765547510, -100141546156, -100127679929, -100159196939, -100172853328, -100176554751, -100171831926, -100153333418, -100172074057, -100179863459, -1001661204967)
	data, _ := yaml.Marshal(c)
	err := os.WriteFile("config.yaml", data, 0666)
	if err != nil {
		log.Errorln("写入配置到配置文件失败" + err.Error())
		return
	}
}

func checkYesOrNo(data string) bool {
	data = strings.ToUpper(data)
	if data == "Y" || data == "YES" {
		return true
	} else {
		return false
	}
}

func getCmdLine() string {
	log.Info("请输入: ")
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(data)
}

func checkFile(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func init() {
	log.SetFormatter(&nested.Formatter{
		FieldsOrder:           nil,
		TimeStampFormat:       "2006-01-02 15:04:05",
		CharStampFormat:       "",
		HideKeys:              false,
		Position:              false,
		Colors:                true,
		FieldsColors:          true,
		FieldsSpace:           true,
		ShowFullLevel:         false,
		LowerCaseLevel:        true,
		TrimMessages:          true,
		CallerFirst:           false,
		CustomCallerFormatter: nil,
	})
}

var (
	task = `
- Container:
    - 1
  Env: jd_wxCartKoi_activityId
  KeyWord:
    - jd_wxCartKoi_activityId
  Name: 【云上】购物车锦鲤
  Script: jd_wxCartKoi.js
  TimeOut: 0
  Wait: 5
- Container:
    - 1
  Env: WXGAME_ACT_ID
  KeyWord:
    - WXGAME_ACT_ID
  Name: 【云上】通用游戏任务
  Script: jd_wxgame.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxFansInterActionActivity_activityId
  KeyWord:
    - jd_wxFansInterActionActivity_activityId
  Name: 【云上】粉丝互动
  Script: jd_wxFansInterActionActivity.js
  TimeOut: 0
  Wait: 10
- Container:
    - 1
  Env: jd_wxSecond_activityId
  KeyWord:
    - jd_wxSecond_activityId
  Name: 【云上】读秒拼手速
  Script: jd_wxSecond.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: computer_activityId
  KeyWord:
    - computer_activityId
  Name: 【云上】电脑配件
  Script: jd_computer.js
  TimeOut: 0
  Wait: 5
  OverdueTime: 600
- Container:
    - 1
  Env: jd_fxyl_activityId
  KeyWord:
    - jd_fxyl_activityId
  Name: 【云上】分享有礼
  Script: jd_share.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_cjhy_wxKnowledgeActivity_activityId
  KeyWord:
    - jd_cjwxKnowledgeActivity_activityId
    - jd_cjhy_wxKnowledgeActivity_activityId
  Name: 【云上】CJ知识超人
  Script: jd_cjhy_wxKnowledgeActivity.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_lzkj_wxKnowledgeActivity_activityId
  KeyWord:
    - jd_lzkj_wxKnowledgeActivity_activityId
    - jd_wxKnowledgeActivity_activityId
  Name: 【云上】LJ知识超人
  Script: jd_lzkj_wxKnowledgeActivity.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_cjhy_activityId
  KeyWord:
    - jd_cjhy_activityId
  Name: 【组队】CJ瓜分京豆
  Script: jd_cjzdgf.js
  TimeOut: 0
  Wait: 5
- Container:
    - 1
  Env: jd_zdjr_activityId
  KeyWord:
    - jd_zdjr_activityId
  Name: 【组队】ZJ瓜分京豆
  Script: jd_zdjr.js
  TimeOut: 0
  Wait: 5
- Container:
    - 1
  Env: jd_cjhy_wxCollectionActivityId
  KeyWord:
    - jd_cjhy_wxCollectionActivityId
  Name: 【抽奖】cjhy加购物车
  Script: jd_cjhy_wxCollectionActivity.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_cjhy_wxDrawActivity_Id
  KeyWord:
    - jd_cjhy_wxDrawActivity_Id
  Name: 【抽奖】cjhy幸运抽大奖
  Script: jd_cjhy_wxCollectionActivity.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_lzkj_wxCollectionActivityId
  KeyWord:
    - jd_lzkj_wxCollectionActivityId
  Name: 【抽奖】lzkj加购物车
  Script: jd_lzkj_wxCollectionActivity.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxCollectCard_activityId
  KeyWord:
    - jd_wxCollectCard_activityId
  Name: 【抽奖】集卡抽奖通用
  Script: jd_wxCollectCard.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: JD_Lottery
  KeyWord:
    - JD_Lottery
  Name: 【抽奖】joy抽奖机通用
  Script: jd_lotterys.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxShopFollowActivity_activityId
  KeyWord:
    - jd_wxShopFollowActivity_activityId
  Name: 【抽奖】关注店铺抽奖
  Script: jd_wxShopFollowActivity.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_mhurlLis
  KeyWord:
    - jd_mhurlLis
  Name: 【抽奖】盲盒抽京豆
  Script: jd_mhtask.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_nzmhurl
  KeyWord:
    - jd_nzmhurl
  Name: 【抽奖】女装盲盒抽京豆
  Script: jd_nzmh.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: DPLHTY
  KeyWord:
    - DPLHTY
  Name: 【开卡】大牌联合
  Script: jd_opencardLH.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: VENDER_ID
  KeyWord:
    - VENDER_ID
  Name: 【开卡】入会开卡领取礼包
  Script: jd_card_force.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wdz_activityId
  KeyWord:
    - jd_wdz_activityId
  Name: 【开卡】微定制
  Script: jd_wdz.js
  TimeOut: 0
  Wait: 10
- Container:
    - 1
  Env: JD_JOYOPEN
  KeyWord:
    - JD_JOYOPEN
  Name: 【开卡】JoyJd任务脚本
  Script: jd_opencard_joyopen.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wdz_openLuckBag_activityId
  KeyWord:
    - jd_wdz_openLuckBag_activityId
  Name: 【开卡】微定制-开福袋
  Script: jd_wdz_openLuckBag.js
  TimeOut: 0
  Wait: 5
- Container:
    - 1
  Env: DPQDTK
  KeyWord:
    - DPQDTK
  Name: 【签到】店铺签到
  Script: jd_dpqd.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: M_WX_LUCK_DRAW_URL
  KeyWord:
    - M_WX_LUCK_DRAW_URL
  Name: 【M系】幸运抽奖
  Script: m_jd_wx_luckDraw.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: M_WX_ADD_CART_URL
  KeyWord:
    - M_WX_ADD_CART_URL
  Name: 【M系】加购有礼
  Script: m_jd_wx_addCart.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: M_WX_COLLECT_CARD_URL
  KeyWord:
    - M_WX_COLLECT_CARD_URL
  Name: 【M系】集卡抽奖
  Script: m_jd_wx_collectCard.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: M_WX_CENTER_DRAW_URL
  KeyWord:
    - M_WX_CENTER_DRAW_URL
  Name: 【M系】老虎机抽奖
  Script: m_jd_wx_centerDraw.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: M_FAV_SHOP_ARGV
  KeyWord:
    - M_FAV_SHOP_ARGV
  Name: 【M系】收藏有礼
  Script: m_jd_fav_shop_gift.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: M_FOLLOW_SHOP_ARGV
  KeyWord:
    - M_FOLLOW_SHOP_ARGV
  Name: 【M系】关注有礼
  Script: m_jd_follow_shop.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: M_WX_SHOP_GIFT_URL
  KeyWord:
    - M_WX_SHOP_GIFT_URL
  Name: 【M系】关注有礼无线
  Script: m_jd_wx_shopGift.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: M_WX_FOLLOW_DRAW_URL
  KeyWord:
    - M_WX_FOLLOW_DRAW_URL
  Name: 【M系】关注抽奖
  Script: m_jd_wx_followDraw.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: M_WX_BUILD_DRAW_URL
  KeyWord:
    - M_WX_BUILD_DRAW_URL
  Name: 【M系】盖楼领奖
  Script: m_jd_wx_buildDraw.js
  TimeOut: 0
  Wait: 2
  OverdueTime: 1800
- Container:
    - 1
  Env: jinggengInviteJoin
  KeyWord:
    - jinggengInviteJoin
  Name: 【船长】邀请入会有礼
  Script: HarbourJ_HarbourToulu_main/jd_jinggengInvite.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_inv_authorCode
  KeyWord:
    - jd_inv_authorCode
    - yhyauthorCode
  Name: 【船长】邀请赢大礼
  Script: HarbourJ_HarbourToulu_main/jd_inviteFriendsGift.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxShopGiftId
  KeyWord:
    - jd_wxShopGiftId
  Name: 【船长】特效关注有礼
  Script: HarbourJ_HarbourToulu_main/jd_wxShopGift.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_joinCommonId
  KeyWord:
    - jd_joinCommonId
  Name: 【船长】通用开卡
  Script: HarbourJ_HarbourToulu_main/jd_joinCommon_opencard.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_shopLeagueId
  KeyWord:
    - jd_shopLeagueId
  Name: 【船长】开卡-shopLeague系列
  Script: HarbourJ_HarbourToulu_main/jd_shopLeague_opencard.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_shopCollectGiftId
  KeyWord:
    - jd_shopCollectGiftId
  Name: 【船长】店铺会员礼包
  Script: HarbourJ_HarbourToulu_main/jd_shopCollectGift.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wdz_activityId
  KeyWord:
    - jd_wdz_activityId
  Name: 【船长】微定制组队
  Script: HarbourJ_HarbourToulu_main/jd_wdz.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxCompleteInfoId
  KeyWord:
    - jd_wxCompleteInfoId
  Name: 【船长】完善信息有礼
  Script: HarbourJ_HarbourToulu_main/jd_wxCompleteInfo.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: M_WX_WDZ_ID
  KeyWord:
    - M_WX_WDZ_ID
  Name: 【M系列】微定制
  Script: m_jd_wx_microDz.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: M_WX_SECOND_DRAW_URL
  KeyWord:
    - M_WX_SECOND_DRAW_URL
  Name: 【M系列】读秒拼手速
  Script: m_jd_wx_secondDraw.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxBirthGiftsId
  KeyWord:
    - jd_wxBirthGiftsId
  Name: 【船长】生日礼包
  Script: HarbourJ_HarbourToulu_main/jd_wxBirthGifts.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_lzkj_wxBuildActivity_activityId
  KeyWord:
    - jd_lzkj_wxBuildActivity_activityId
    - jd_wxBuildActivity_activityId
  Name: 【Faker库】盖楼有礼
  Script: jd_lzkj_wxBuildActivity.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: JD_Lottery
  KeyWord:
    - JD_Lottery
  Name: 【Faker库】joy抽奖机通用
  Script: jd_lotterys.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: VENDER_ID
  KeyWord:
    - VENDER_ID
  Name: 【Faker库】入会开卡领取礼包通用
  Script: jd_card_force.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: whx_drawShopGift
  KeyWord:
    - whx_drawShopGift
  Name: 【小埋】关注有礼-自动解析通用
  Script: jd_whx_drawShopGift.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: prodevactCode
  KeyWord:
    - prodevactCode
  Name: 【小埋】邀请好友入会赢好礼
  Script: jd_prodev.js
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxCollectionActivityUrl
  KeyWord:
    - jd_wxCollectionActivityUrl
  Name: 加购有礼
  Script: HarbourJ_HarbourToulu_main/jd_wxCollectionActivity.py
  TimeOut: 0
  Wait: 0

`
)
