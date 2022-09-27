package conf

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Aauthentication string `json:"Aauthentication"   mapstructure:"Aauthentication"`
	AutoRestart     string `json:"AutoRestart" mapstructure:"AutoRestart"`
	ContainerWait   int    `json:"Container_Wait" mapstructure:"Container_Wait"`
	LogEnable       bool   `json:"LogEnable" mapstructure:"LogEnable"`
	LogLevel        string `json:"LogLevel" mapstructure:"LogLevel"`
	QingLong        []struct {
		ClientID     string `json:"Client_ID" mapstructure:"Client_ID"`
		ClientSecret string `json:"Client_Secret" mapstructure:"Client_Secret"`
		Url          string `json:"url" mapstructure:"url"`
	} `json:"QingLong" mapstructure:"QingLong"`
	SpyEnable string `json:"SpyEnable" mapstructure:"SpyEnable"`
	Telegram  struct {
		LogEn            bool     `json:"Log_En" mapstructure:"Log_En"`
		TgLost           string   `json:"TgLost" mapstructure:"TgLost"`
		Welcome          string   `json:"Welcome" mapstructure:"Welcome"`
		ApiHash          string   `json:"api_hash" mapstructure:"api_hash"`
		ApiId            int      `json:"api_id" mapstructure:"api_id"`
		ForwardChannelId string   `json:"forward_channel_id" mapstructure:"forward_channel_id"`
		ForwardKeyword   []string `json:"forward_keyword" mapstructure:"forward_keyword"`
		ForwardUserId    []int64  `json:"forward_user_id" mapstructure:"forward_user_id"`
		ForwardUserName  []string `json:"forward_user_name" mapstructure:"forward_user_name"`
		ListenCH         []int64  `json:"listen_CH" mapstructure:"listen_CH"`
		LogId            int64    `json:"log_id" mapstructure:"log_id"`
		MasterId         []int64  `json:"master_id" mapstructure:"master_id"`
		ProxyURL         string   `json:"proxy_url" mapstructure:"proxy_url"`
	} `json:"Telegram" mapstructure:"Telegram"`
	UseSillyGirl bool          `json:"Use_SillyGirl" mapstructure:"Use_SillyGirl"`
	WaitTime     int           `json:"WaitTime" mapstructure:"WaitTime"`
	Branch       string        `json:"branch" mapstructure:"branch"`
	DownProxy    string        `json:"down_proxy" mapstructure:"down_proxy"`
	JsConfig     []*TaskConfig `json:"js_config" mapstructure:"js_config"`
	UpdateUrl    string        `json:"update_url" mapstructure:"update_url"`
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
	config *Config
)

func InitConfig() {
	once := sync.Once{}
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Errorln("配置文件不存在")
				data, _ := yaml.Marshal(new(Config))
				_ = os.WriteFile("config.yaml", data, 0666)
				os.Exit(3)
			} else {
				log.Errorln(err.Error())
			}
		}
		viper.SetDefault("LogLevel", "info")
		c := new(Config)
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
