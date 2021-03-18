package hallinpay

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// 表单POST请求
func PostForm(conf *Config, path string, bm map[string]string, result interface{}) (err error) {
	if conf == nil {
		err = errors.New("收银宝接口配置不能为空")
		return
	}

	var (
		begTime   = time.Now().Nanosecond()
		requestId = Rand32()
		reqUrl    = conf.BaseUrl + path
		reqJson   string
		resJson   string
		plog      = log.WithField("requestId", requestId)
	)
	defer func() {
		plog.Info("allinpay 请求URL:" + reqUrl)
		plog.Info("allinpay 请求报文:" + reqJson)
		plog.WithField("ms", (time.Now().Nanosecond()-begTime)/1e6).Info("allinpay 响应报文:" + resJson)
		if err != nil {
			plog.Error("allinpay 请求异常:" + err.Error())
			return
		}

	}()

	if bm == nil {
		bm = make(map[string]string)
	}

	bm["cusid"] = conf.Cusid
	bm["appid"] = conf.Appid
	bm["version"] = VERSION_11
	bm["randomstr"] = Rand32()
	bm["signtype"] = SIGN_TYPE_RSA

	sign, err := RsaSign(conf, bm)
	if err != nil {
		return
	}
	bm["sign"] = sign

	reqJsonBytes, err := json.Marshal(bm)
	reqJson = string(reqJsonBytes)

	data := url.Values{}
	for k, v := range bm {
		if v == "" {
			continue
		}
		data.Set(k, v)
	}
	resp, err := http.PostForm(reqUrl, data)
	if err != nil {
		return
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resJson = string(ret)

	var retBm = make(map[string]string)
	err = json.Unmarshal(ret, &retBm)
	if err != nil {
		return
	}

	// 非成功状态，不验证签名
	if retBm["retcode"] == string(RET_FAIl) {
		err = errors.New(retBm["retcode"] + ":" + retBm["retmsg"])
		return
	}

	err = RsaVerify(conf, retBm)
	if err != nil {
		return
	}

	err = json.Unmarshal(ret, &result)
	if err != nil {
		return
	}
	return
}
