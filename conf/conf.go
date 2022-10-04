package conf

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type QingLongConf struct {
	ClientID     string `json:"Client_ID" mapstructure:"Client_ID" yaml:"ClientID"`
	ClientSecret string `json:"Client_Secret" mapstructure:"Client_Secret" yaml:"ClientSecret"`
	Url          string `json:"url" mapstructure:"url" yaml:"url"`
}

type Config struct {
	LogLevel string         `json:"LogLevel" mapstructure:"LogLevel" yaml:"LogLevel"`
	QingLong []QingLongConf `json:"QingLong" mapstructure:"QingLong" yaml:"QingLong"`
	Telegram struct {
		LogEn            bool     `json:"Log_En" mapstructure:"Log_En" yaml:"Log_En"`
		TgLost           string   `json:"TgLost" mapstructure:"TgLost" yaml:"TgLost"`
		Welcome          string   `json:"Welcome" mapstructure:"Welcome" yaml:"Welcome"`
		ApiHash          string   `json:"api_hash" mapstructure:"api_hash" yaml:"api_hash"`
		ApiId            int      `json:"api_id" mapstructure:"api_id" yaml:"api_id"`
		ForwardChannelId string   `json:"forward_channel_id" mapstructure:"forward_channel_id" yaml:"forward_channel_id"`
		ForwardKeyword   []string `json:"forward_keyword" mapstructure:"forward_keyword" yaml:"forward_keyword"`
		ForwardUserId    []int64  `json:"forward_user_id" mapstructure:"forward_user_id" yaml:"forward_user_id"`
		ForwardUserName  []string `json:"forward_user_name" mapstructure:"forward_user_name" yaml:"forward_user_name"`
		ListenCH         []int64  `json:"listen_CH" mapstructure:"listen_CH" yaml:"listen_CH"`
		LogId            int64    `json:"log_id" mapstructure:"log_id" yaml:"log_id"`
		MasterId         []int64  `json:"master_id" mapstructure:"master_id" yaml:"master_id"`
		ProxyURL         string   `json:"proxy_url" mapstructure:"proxy_url" yaml:"proxy_url"`
	} `json:"Telegram" mapstructure:"Telegram" yaml:"Telegram"`
	WaitTime  int           `json:"WaitTime" mapstructure:"WaitTime" yaml:"WaitTime"`
	DownProxy string        `json:"down_proxy" mapstructure:"down_proxy" yaml:"down_proxy"`
	JsConfig  []*TaskConfig `json:"js_config" mapstructure:"js_config" yaml:"js_config"`
}

type TaskConfig struct {
	Container   []int    `json:"Container"   mapstructure:"Container" yaml:"Container"`
	Env         string   `json:"Env" mapstructure:"Env" yaml:"Env"`
	KeyWord     []string `json:"KeyWord" mapstructure:"KeyWord" yaml:"KeyWord"`
	Name        string   `json:"Name" mapstructure:"Name" yaml:"Name"`
	Script      string   `json:"Script" mapstructure:"Script" yaml:"Script"`
	TimeOut     int      `json:"TimeOut" mapstructure:"TimeOut" yaml:"TimeOut"`
	Wait        int      `json:"Wait" mapstructure:"Wait" yaml:"Wait"`
	OverdueTime int      `json:"OverdueTime,omitempty" mapstructure:"OverdueTime,omitempty" yaml:"OverdueTime,omitempty"`
	Disable     int      `json:"Disable,omitempty" mapstructure:"Disable,omitempty" yaml:"Disable,omitempty"`
}

var (
	config  *Config
	Version = "UNKNOWN"
)

func GetVersion() string {
	return Version
}

func InitConfig() {
	once := sync.Once{}
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		c := new(Config)
		err := viper.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				c = generateConfig()
				log.Infoln("配置文件config.yaml生成成，请重启程序！")
				os.Exit(1)
			} else {
				log.Errorln(err.Error())
			}
		}
		viper.SetDefault("LogLevel", "info")
		viper.SetDefault("WaitTime", 3)

		err = viper.Unmarshal(c)
		if err != nil {
			log.Errorln("配置解析失败" + err.Error())
			return
		}
		config = c
		file, err := os.ReadFile("task.yaml")
		if err != nil {
			log.Errorln("任务配置文件task.json不存在" + err.Error())
			return
		}
		err = yaml.Unmarshal(file, &c.JsConfig)
		if err != nil {
			log.Errorln("解析task.yaml失败" + err.Error())
			return
		}
		config = c

	})

}

func SaveTask(tasks []*TaskConfig) error {
	data, err := yaml.Marshal(&tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("task.yaml", data, 0666)
	if err != nil {
		return err
	}
	return err
}

func GetConfig() *Config {
	if config == nil {
		InitConfig()
	}
	return config
}

func generateConfig() *Config {
	log.Infoln("检测到配置文件不存在，是否生成配置文件(y or n)")
	if !checkYesOrNo(getCmdLine()) {
		os.Exit(3)
	}

	c := new(Config)
	c.LogLevel = "info"
	log.Infoln("请输入你的青龙访问地址,例如： http://127.0.0.1:5701")
	qlAdd := getCmdLine()
	log.Infoln("请输入你的青龙应用的client_id")
	qlClientID := getCmdLine()
	log.Infoln("请输入你的青龙应用的client_secret")
	qlClientSecret := getCmdLine()
	c.QingLong = append(c.QingLong, QingLongConf{
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
		os.Exit(3)
	}
	log.Infoln("配置文件生成成功")
	return c
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
