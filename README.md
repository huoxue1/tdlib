# tg机器人ql监控

## 如何使用

在release下载文件，直接运行

## 命令列表
[点我查看命令列表](./doc/cmd.md)

## 配置文件
+ config.yaml

兼容auto_spy配置文件，只是proxy改为了proxy_url，同时将js_config独立未task.yaml文件

```yaml
QingLong:
    - Client_ID: 123
      Client_Secret: 123
      url: http://127.0.0.1:5700
Telegram:
    api_hash: 123
    api_id: 123
    listen_CH:
        - -1001765547510
        - -1001415461569
        - -1001276799295
        - -1001591969398
        - -1001728533280
        - -1001765547510
        - -1001718319262
        - -1001533334185
        - -1001720740578
        - -1001798634590
    log_id: -123456
    master_id:
        - 12345
    # 配置格式 proto://username:password@ip:port
    # socks5 socks5：//user:pass@127.0.0.1:7890
    # http   http://user:pass@127.0.0.1:7890
    proxy_url: "socks5://127.0.0.1:7891"
js_config:
    - Container:
          - 1
      Env: jd_wxCartKoi_activityId
      KeyWord:
          - jd_wxCartKoi_activityId
      Name: 【云上】购物车锦鲤
      Script: jd_wxCartKoi.js
      TimeOut: 0
```

+ task.yaml
```yaml
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
  Script: jd_jinggengInvite.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_inv_authorCode
  KeyWord:
    - jd_inv_authorCode
    - yhyauthorCode
  Name: 【船长】邀请赢大礼
  Script: jd_inviteFriendsGift.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxShopGiftId
  KeyWord:
    - jd_wxShopGiftId
  Name: 【船长】特效关注有礼
  Script: jd_wxShopGift.py
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
  Script: jd_shopLeague_opencard.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_shopCollectGiftId
  KeyWord:
    - jd_shopCollectGiftId
  Name: 【船长】店铺会员礼包
  Script: jd_shopCollectGift.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wdz_activityId
  KeyWord:
    - jd_wdz_activityId
  Name: 【船长】微定制组队
  Script: jd_wdz.py
  TimeOut: 0
  Wait: 2
- Container:
    - 1
  Env: jd_wxCompleteInfoId
  KeyWord:
    - jd_wxCompleteInfoId
  Name: 【船长】完善信息有礼
  Script: jd_wxCompleteInfo.py
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
  Script: jd_wxBirthGifts.py
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
  Script: jd_wxCollectionActivity.py
  TimeOut: 0
  Wait: 0

```