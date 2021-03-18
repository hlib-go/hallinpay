package hallinpay

// 支持部分金额退款，隔天交易退款     注：含单品优惠交易只能整单退款，不支持部分退款
func Refund(conf *Config, p *RefundParams) (result *RefundResult, err error) {
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
	Trxid     string `json:"trxid"`     //收银宝平台的退款交易流水号
	Reqsn     string `json:"reqsn"`     //商户的退款交易订单号
	Trxcode   string `json:"trxcode"`   //交易类型
	Trxstatus string `json:"trxstatus"` //交易的状态
	Errmsg    string `json:"errmsg"`    //失败的原因说明
	Fintime   string `json:"fintime"`
	Fee       int64  `json:"fee"` //手续费
}
