# tg机器人ql监控

## 如何使用

在release下载文件，直接运行


## 配置文件
兼容auto_spy配置文件，只是proxy改为了proxy_url

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