package utils

import (
	"encoding/json"
	"time"

	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func GetJxBean(cookie string) (*JxBean, error) {
	response, err := req.C().R().SetHeaders(map[string]string{
		"Origin":       "https://wqs.jd.com",
		"User-Agent":   "jdapp;iPhone;9.4.4;14.3;network/4g;Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148;supportJDSHWK/1",
		"Host":         "api.m.jd.com",
		"Content-Type": "application/x-www-form-urlencoded",
		"Cookie":       cookie,
	}).SetQueryParams(map[string]string{
		"appid":        "jd-cphdeveloper-m",
		"functionId":   "myBean",
		"body":         `{"tenantCode":"jgm","bizModelCode":6,"bizModeClientType":"M","externalLoginType":1}`,
		"g_login_type": "0",
		"g_tk":         "398644289",
		"g_ty":         "ajax",
		"appCode":      "msc588d6d5",
	},
	).Get("https://api.m.jd.com")
	if err != nil {
		return nil, err
	}
	j := new(JxBean)
	err = json.Unmarshal(response.Bytes(), j)
	if err != nil {
		return nil, err
	}
	return j, err
}

func GetBeanDetail(cookie string) ([]*Detail, error) {
	var details []*Detail

	response, err := req.C().R().SetHeaders(map[string]string{
		"User-Agent":   "jdapp;iPhone;9.4.4;14.3;network/4g;Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148;supportJDSHWK/1",
		"Host":         "api.m.jd.com",
		"Content-Type": "application/x-www-form-urlencoded",
		"Cookie":       cookie,
	}).SetBodyString("body=%7B%22pageSize%22%3A%2220%22%2C%22page%22%3A%221%22%7D&appid=ld").Post("https://api.m.jd.com/client.action?functionId=getJingBeanBalanceDetail")
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	for _, result := range gjson.GetBytes(response.Bytes(), "detailList").Array() {
		d := new(Detail)
		err := json.Unmarshal([]byte(result.String()), d)
		if err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		data, err := time.Parse("2006-01-02 15:04:05", d.Date)
		if err != nil {
			return nil, err
		}
		if data.Day() == time.Now().Day() {
			details = append(details, d)
		}
	}
	return details, err

}

func TotalBean(cookie string) (*TotalBeanResp, error) {
	t := new(TotalBeanResp)
	response, err := req.C().R().SetHeaders(map[string]string{
		"Host":            "me-api.jd.com",
		"Accept":          "*/*",
		"Connection":      "keep-alive",
		"Cookie":          cookie,
		"User-Agent":      "jdapp;iPhone;9.4.4;14.3;network/4g;Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148;supportJDSHWK/1",
		"Accept-Language": "zh-cn",
		"Referer":         "https://home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&",
	}).Get("https://me-api.jd.com/user_new/info/GetJDUserInfoUnion")
	if err != nil {
		log.Errorln(err.Error())
		return t, err
	}
	err = json.Unmarshal(response.Bytes(), t)
	if err != nil {
		return nil, err
	}
	return t, err

}
