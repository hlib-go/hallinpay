package hallinpay

import "errors"

// 交易查询 ,reqsn = 商户订单号 trxid=平台交易流水  reqsn和trxid必填其一 建议:商户如果同时拥有trxid和reqsn,优先使用trxid
func ApiwebUnitorderQuery(conf *Config, reqsn, trxid string) (result *QueryResult, err error) {
	var bm = make(map[string]string)
	bm["trxid"] = trxid
	bm["reqsn"] = reqsn

	err = PostForm(conf, "/apiweb/unitorder/query", bm, &result)
	if err != nil {
		return
	}
	if result.RetCode != RET_SUCCESS {
		err = errors.New(string(result.RetCode) + ":" + result.RetMsg)
		return
	}
	return
}

// 文档： https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=93
type QueryResult struct {
	RetCode RetCode `json:"retcode"`
	RetMsg  string  `json:"retmsg"`
	//以下信息只有当retcode为SUCCESS时有返回
	Trxid     string `json:"trxid"`
	Chnltrxid string `json:"chnltrxid"`
	Reqsn     string `json:"reqsn"`
	Trxcode   string `json:"trxcode"`
	Trxamt    string `json:"trxamt"`
	Trxstatus string `json:"trxstatus"`
	Errmsg    string `json:"errmsg"`
	Acct      string `json:"acct"`
	Fintime   string `json:"fintime"`
	Cmid      string `json:"cmid"`
	Chnlid    string `json:"chnlid"`
	Initamt   string `json:"initamt"`
	Fee       int64  `json:"fee"`
	Chnldata  string `json:"chnldata"`
}
