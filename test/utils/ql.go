package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	"github.com/huoxue1/tdlib"
)

type Ql struct {
	Api          string
	ClientSecret string
	ClientID     string

	token   string
	c       *req.Client
	storage *tdlib.MyDB
	header  map[string]string
}

type Cron struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Command   string `json:"command"`
	Schedule  string `json:"schedule"`
	Timestamp string `json:"timestamp"`
	Saved     bool   `json:"saved"`
	// 1代表未运行，0代表运行中
	Status            int           `json:"status"`
	IsSystem          int           `json:"isSystem"`
	Pid               interface{}   `json:"pid"`
	IsDisabled        int           `json:"isDisabled"`
	IsPinned          int           `json:"isPinned"`
	LogPath           string        `json:"log_path"`
	Labels            []interface{} `json:"labels"`
	LastRunningTime   int           `json:"last_running_time"`
	LastExecutionTime int           `json:"last_execution_time"`
	CreatedAt         time.Time     `json:"createdAt"`
	UpdatedAt         time.Time     `json:"updatedAt"`
}

type Env struct {
	Id        int       `json:"id"`
	Value     string    `json:"value"`
	Timestamp string    `json:"timestamp"`
	Status    int       `json:"status"`
	Position  float64   `json:"position"`
	Name      string    `json:"name"`
	Remarks   string    `json:"remarks"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func InitQl(api, clientID, clientSecret string) *Ql {
	q := new(Ql)
	db, _ := tdlib.InitDB()
	q.storage = db
	q.c = req.C()
	q.Api = api
	q.ClientID = clientID
	q.ClientSecret = clientSecret

	token, err := q.storage.Load("ql_" + api)
	if err != nil {
		log.Errorln("从磁盘读取token失败" + err.Error())
		q.token = q.getToken()
		err = nil
	} else {
		if string(token) == "" {
			q.token = q.getToken()
		} else {
			q.token = token
		}
	}
	err = q.storage.Store("ql_"+api, q.token)
	if err != nil {
		log.Errorln("store token faild," + err.Error())
		return q
	}
	q.c.SetCommonHeader("Authorization", "Bearer "+q.token)
	return q
}

func (q *Ql) checkToken() bool {

	return true
}

func (q *Ql) getToken() string {
	response, err := q.c.R().SetQueryParams(map[string]string{
		"client_id":     q.ClientID,
		"client_secret": q.ClientSecret,
	}).Get(fmt.Sprintf("%v/open/auth/token", q.Api))
	if err != nil {
		log.Errorln("请求青龙token出现错误" + err.Error())
		return ""
	}
	if gjson.GetBytes(response.Bytes(), "code").Int() != 200 {
		log.Errorln("请求青龙token出现错误" + gjson.GetBytes(response.Bytes(), "data").String())
		return ""
	}
	return gjson.GetBytes(response.Bytes(), "data.token").String()
}

// SaveConfigFile
/* @Description: 保存配置文件
*  @receiver q
*  @param name
*  @param content
*  @return error
 */
func (q *Ql) SaveConfigFile(name, content string) error {
	response, err := q.c.R().SetBodyJsonMarshal(map[string]string{
		"name":    name,
		"content": content,
	}).Post(q.Api + "/open/configs/save")
	if err != nil {
		return err
	}
	res := response.Bytes()
	if gjson.GetBytes(res, "code").Int() != 200 {
		log.Errorln("请求api错误" + gjson.GetBytes(res, "data").String())
		return err
	}
	return err
}

func (q *Ql) GetConfigFile(name string) (string, error) {
	response, err := q.c.R().Get(q.Api + "/open/configs/" + name)
	if err != nil {
		return "", err
	}
	res := response.Bytes()
	fmt.Println(string(res))
	if gjson.GetBytes(res, "code").Int() != 200 {
		log.Errorln("请求api错误" + gjson.GetBytes(res, "data").String())
		return "", err
	}
	return gjson.GetBytes(res, "data").String(), err
}

func (q *Ql) GetCrons(search string) ([]*Cron, error) {
	var cron []*Cron
	response, err := q.c.R().SetQueryParam("searchValue", search).Get(q.Api + "/open/crons/")
	if err != nil {
		return nil, err
	}
	res := response.Bytes()
	if gjson.GetBytes(res, "code").Int() != 200 {
		log.Errorln("请求api错误" + gjson.GetBytes(res, "data").String())
		return nil, err
	}
	for _, result := range gjson.GetBytes(res, "data").Array() {
		c := new(Cron)
		err := json.Unmarshal([]byte(result.String()), c)
		if err != nil {
			return nil, err
		}
		cron = append(cron, c)
	}
	return cron, err
}

// GetCron
/* @Description: 获取定时任务信息
*  @receiver q
*  @param id
*  @return *Cron
*  @return error
 */
func (q *Ql) GetCron(id int) (*Cron, error) {
	response, err := q.c.R().Get(q.Api + "/open/crons/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	res := response.Bytes()
	if gjson.GetBytes(res, "code").Int() != 200 {
		log.Errorln("请求api错误" + gjson.GetBytes(res, "data").String())
		return nil, err
	}

	c := new(Cron)
	err = json.Unmarshal([]byte(gjson.GetBytes(res, "data").String()), c)
	if err != nil {
		return nil, err
	}

	return c, err
}

// RunCrons
/* @Description: 运行定时任务
*  @receiver q
*  @param cronIDs
*  @return error
 */
func (q *Ql) RunCrons(cronIDs ...int) error {
	response, err := q.c.R().SetBodyJsonMarshal(cronIDs).Put(q.Api + "/open/crons/run")
	if err != nil {
		return err
	}
	res := response.Bytes()
	log.Debugln(string(res))
	if gjson.GetBytes(res, "code").Int() != 200 {
		log.Errorln("请求api错误" + gjson.GetBytes(res, "data").String())
		return err
	}
	return err
}

// StopCrons
/* @Description: 停止定时任务
*  @receiver q
*  @param cronIDs
*  @return error
 */
func (q *Ql) StopCrons(cronIDs ...int) error {
	response, err := q.c.R().SetBodyJsonMarshal(cronIDs).Put(q.Api + "/open/crons/stop")
	if err != nil {
		return err
	}
	res := response.Bytes()
	log.Debugln(string(res))
	if gjson.GetBytes(res, "code").Int() != 200 {
		log.Errorln("请求api错误" + gjson.GetBytes(res, "data").String())
		return err
	}
	return err
}

// LogCrons
/* @Description: 获取日志
*  @receiver q
*  @param cronID
*  @return string
*  @return error
 */
func (q *Ql) LogCrons(cronID int) (string, error) {
	response, err := q.c.R().Get(q.Api + fmt.Sprintf("/open/crons/%d/log", cronID))
	if err != nil {
		return "", err
	}
	res := response.Bytes()
	log.Debugln(string(res))
	if gjson.GetBytes(res, "code").Int() != 200 {
		log.Errorln("请求api错误" + gjson.GetBytes(res, "data").String())
		return "", err
	}
	return gjson.GetBytes(res, "data").String(), err
}

func (q *Ql) GetEnvs(search string) ([]*Env, error) {
	var envs []*Env
	response, err := q.c.R().SetQueryParam("searchValue", search).Get(q.Api + "/open/envs/")
	if err != nil {
		return nil, err
	}
	res := response.Bytes()
	if gjson.GetBytes(res, "code").Int() != 200 {
		log.Errorln("请求api错误" + gjson.GetBytes(res, "data").String())
		return nil, err
	}
	for _, result := range gjson.GetBytes(res, "data").Array() {
		e := new(Env)
		err := json.Unmarshal([]byte(result.String()), e)
		if err != nil {
			return nil, err
		}
		envs = append(envs, e)
	}
	return envs, err
}
