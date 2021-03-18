package hallinpay

import "errors"

// 支持部分金额退款，隔天交易退款     注：含单品优惠交易只能整单退款，不支持部分退款
func ApiwebUnitorderRefund(conf *Config, p *RefundParams) (result *RefundResult, err error) {
	var bm = make(map[string]string)
	bm["trxamt"] = p.Trxamt
	bm["reqsn"] = p.Reqsn
	bm["oldreqsn"] = p.Oldreqsn
	bm["oldtrxid"] = p.Oldtrxid
	bm["remark"] = p.Remark
	bm["benefitdetail"] = p.Benefitdetail

	err = PostForm(conf, "/apiweb/unitorder/refund", bm, &result)
	if err != nil {
		return
	}
	if result.RetCode != RET_SUCCESS {
		err = errors.New(string(result.RetCode) + ":" + result.RetMsg)
		return
	}
	return
}

type RefundParams struct {
	Trxamt        string `json:"trxamt"`   //退款金额
	Reqsn         string `json:"reqsn"`    // 退款订单号
	Oldreqsn      string `json:"oldreqsn"` //原交易的商户订单号
	Oldtrxid      string `json:"oldtrxid"` // 原交易的收银宝平台流水
	Remark        string `json:"remark"`
	Benefitdetail string `json:"benefitdetail"` // 优惠信息，只适用于银联单品优惠交易的退货
}

type RefundResult struct {
	RetCode RetCode `json:"retcode"`
	RetMsg  string  `json:"retmsg"`
	//以下信息只有当retcode为SUCCESS时有返回
	Trxid     string `json:"trxid"`     //收银宝平台的退款交易流水号
	Reqsn     string `json:"reqsn"`     //商户的退款交易订单号
	Trxcode   string `json:"trxcode"`   //交易类型
	Trxstatus string `json:"trxstatus"` //交易的状态
	Errmsg    string `json:"errmsg"`    //失败的原因说明
	Fintime   string `json:"fintime"`
	Fee       int64  `json:"fee"` //手续费
}
