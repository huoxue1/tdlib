#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Author: huoxue1
Date: 2022/9/30 15:30
Description:: 用于一键安装tdlib️
"""
import io
import os
import platform
from functools import partial

import requests

if platform.system().lower() == "windows":
    import zipfile

print = partial(print, flush=True)


task = '''
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
'''

def updateDependent() -> str:
    """
    更新依赖的主函数
    """
    system = platform.system().lower()
    PyVersion_ = platform.python_version()
    PyVersion = ''.join(PyVersion_.split('.')[:2])
    fileName = ""
    if system == "windows":

        if platform.architecture()[0] == "64bit":
            fileName = f"tdlib_windows_amd64.zip"
            print(f"✅识别本机设备为Windows amd64,Py版本为{PyVersion_}\n")
        else:
            fileName = f"tdlib_windows_386.zip"
    elif system == "darwin":
        fileName = f"tlib_darwin_amd64.tar.gz"
        print(f"✅识别本机设备为MacOS x86_64,Py版本为{PyVersion_}\n")

    else:

        framework = os.uname().machine
        if framework == "x86_64":
            fileName = f"tdlib_linux_amd64.tar.gz"
            print(f"✅识别本机设备为Linux {framework},Py版本为{PyVersion_}\n")
        elif framework == "aarch64" or framework == "arm64":
            fileName = f"tdlib_linux_arm64.tar.gz"
            print(f"✅识别本机设备为Linux {framework},Py版本为{PyVersion_}\n")
        elif framework == "armv7l":
            fileName = f"tdlib_linux_386.tar.gz"
            print(f"✅识别本机设备为Linux {framework},Py版本为{PyVersion_}\n")
        else:
            fileName = f"tdlib_linux_amd64.tar.gz"
            print(f"⚠️无法识别本机设备操作系统,默认本机设备为Linux x86_64,Py版本为{PyVersion_}\n")
    return fileName


def last_version() -> str:
    return requests.get("https://api.github.com/repos/huoxue1/tdlib/releases/latest").json().get("tag_name")


def download(github: str, version: str, binaryName: str):
    print("正在下载文件中，请耐心等待！！！")
    content = requests.get(f"{github}/huoxue1/tdlib/releases/download/{version}/{binaryName}").content
    if platform.system().lower() == "windows":
        with zipfile.ZipFile(io.BytesIO(content)) as zf:
            data = zf.open("tdlib.exe")
            with open("tdlib.exe", "wb") as f:
                f.write(data.read())
    else:
        with open(binaryName, "wb") as f:
            f.write(content)
        os.system(f"tar xvf {binaryName}")
        os.remove(binaryName)


def checkYesOrNo() -> bool:
    data = input("请输入:").lower()
    if data == "y" or data == "yes":
        return True
    else:
        return False


def addSystemctl():
    with open("/etc/systemd/system/tdlib.service", "w", encoding="utf-8") as f:
        f.write(f'''[Unit]
    Description=tdlib
    Documentation=tdlib
    After=network-online.target
    Wants=network-online.target systemd-networkd-wait-online.service

    [Service]
    Restart=always

    ; User and group the process will run as.
    User=root
    Group=root

    WorkingDirectory={os.getcwd()}
    ExecStart={os.getcwd()}/tdlib

    ; Limit the number of file descriptors; see `man systemd.exec` for more limit settings.
    LimitNOFILE=1048576
    ; Unmodified caddy is not expected to use more than that.
    LimitNPROC=512

    [Install]
    WantedBy=multi-user.target''')


def main():
    github = "https://github.com"
    version = last_version()
    binaryName = updateDependent()
    if not os.path.isfile("tdlib.exe") or not os.path.isfile("tdlib"):
        print("检测到tdlib文件已经存在，是否跳过下载(y/n)")
        if not checkYesOrNo():
            download(github, version, binaryName)
        print("已跳过下载")
    else:
        download(github, version, binaryName)
    if not os.path.isfile("task.yaml"):
        global task
        print("检测到目录下未存在任务监听文件，是否生成默认监听配置文件(y/n)")
        if checkYesOrNo():
            with open("task.yaml", "w", encoding="utf-8") as f:
                f.write(task)

    if not platform.system().lower() == "windows":
        print("是否将tdlib加入系统启动system命令(y/n)")
        if checkYesOrNo():
            addSystemctl()
    print("是否启动tdlib? (y/n)")
    if checkYesOrNo():
        if platform.system().lower() == "windows":
            os.system("tdlib.exe")
        else:
            os.system("./tdlib")


if __name__ == '__main__':
    main()


