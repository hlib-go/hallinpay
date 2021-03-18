package hallinpay

import "errors"

// 统一支付接口 , 注意：根据trxstatus判断状态
func ApiwebUnitorderPay(conf *Config, p *PayParams) (result *PayResult, err error) {
	var bm = make(map[string]string)
	bm["trxamt"] = p.Trxamt
	bm["reqsn"] = p.Reqsn
	bm["paytype"] = string(p.Paytype)
	// 微信支付独有字段
	if p.Paytype == PAY_TYPE_W06 || p.Paytype == PAY_TYPE_W02 {
		bm["sub_appid"] = p.SubAppid
	}
	bm["acct"] = p.Acct
	bm["body"] = p.Body
	bm["remark"] = p.Remark
	bm["validtime"] = p.Validtime

	err = PostForm(conf, "/apiweb/unitorder/pay", bm, &result)
	if err != nil {
		return
	}
	if result.RetCode != RET_SUCCESS {
		err = errors.New(string(result.RetCode) + ":" + result.RetMsg)
		return
	}
	return
}

type PayParams struct {
	Trxamt        string  `json:"trxamt"`        //交易金额
	Reqsn         string  `json:"reqsn"`         //商户交易单号
	Paytype       PayType `json:"paytype"`       //交易方式
	Body          string  `json:"body"`          //订单商品名称，为空则以商户名作为商品名称
	Remark        string  `json:"remark"`        //
	Validtime     string  `json:"validtime"`     // 默认5分钟，最大1440分钟
	Acct          string  `json:"acct"`          //JS支付时使用 微信支付-用户的微信openid  支付宝支付-用户user_id  微信小程序-用户小程序的openid  云闪付JS-用户userId
	NotifyUrl     string  `json:"notify_url"`    //
	LimitPay      string  `json:"limit_pay"`     //
	SubAppid      string  `json:"sub_appid"`     //微信小程序/微信公众号/APP的appid
	GoodsTag      string  `json:"goods_tag"`     //
	Benefitdetail string  `json:"benefitdetail"` //
	FrontUrl      string  `json:"front_url"`     //只支持payType=U02云闪付JS支付  payType=W02微信JS支付
}

type PayResult struct {
	RetCode RetCode `json:"retcode"` //SUCCESS/FAIL
	RetMsg  string  `json:"retmsg"`
	//以下信息只有当retcode为SUCCESS时有返回
	Trxid     string    `json:"trxid"`     //收银宝平台的交易流水号
	Chnltrxid string    `json:"chnltrxid"` // 渠道平台交易单号 例如微信,支付宝平台的交易单号
	Reqsn     string    `json:"reqsn"`
	Trxstatus TrxStatus `json:"trxstatus"` //对于刷卡支付，该状态表示实际的支付结果，其他为下单状态
	Fintime   string    `json:"fintime"`   // 交易完成时间 yyyyMMddHHmmss 对于微信刷卡支付有效
	Errmsg    string    `json:"errmsg"`    //失败的原因说明
	Payinfo   string    `json:"payinfo"`   //扫码支付则返回二维码串，js支付则返回json字符串  QQ钱包及云闪付的JS支付返回支付的链接,商户只需跳转到此链接即可完成支付  支付宝App支付返回支付信息串
}

// 交易结果通知文档，表单POST请求  https://aipboss.allinpay.com/know/devhelp/main.php?pid=15#mid=94

/*
交易返回码trxstatus说明
0000：交易成功
1001：交易不存在
2008或者2000 : 交易处理中,请查询交易,如果是实时交易(例如刷卡支付,交易撤销,退货),建议每隔一段时间(10秒)查询交易
3开头的错误码代表交易失败
3888-流水号重复
3889-交易控制失败，具体原因看errmsg
3099-渠道商户错误
3014-交易金额小于应收手续费
3031-校验实名信息失败
3088-交易未支付(在查询时间区间内未成功支付,如已影响资金24小时内会做差错退款处理)
3089-撤销异常,如已影响资金24小时内会做差错退款处理
3045-其他错误，具体原因看errmsg
3050-交易已被撤销
3999-其他错误，具体原因看errmsg
其他3开头的错误码代表交易失败,具体原因请读取errmsg
*/
