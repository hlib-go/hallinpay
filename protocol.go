package hallinpay

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 表单POST请求
func PostForm(conf *Config, path string, bm map[string]string, result interface{}) (err error) {
	if bm == nil {
		bm = make(map[string]string)
	}
	bm["cusid"] = conf.Cusid
	bm["appid"] = conf.Appid
	bm["version"] = VERSION_11
	bm["randomstr"] = Rand32()
	bm["signtype"] = SIGN_TYPE_RSA

	bm["sign"] = ""
	data := url.Values{}
	for k, v := range bm {
		if v == "" {
			continue
		}
		data.Set(k, v)
	}
	resp, err := http.PostForm(conf.BaseUrl+path, data)
	if err != nil {
		return
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(ret, &result)
	if err != nil {
		return
	}
	return
}
