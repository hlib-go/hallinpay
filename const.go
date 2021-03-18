package hallinpay

const (
	VERSION_11 = "11"
)

const (
	SIGN_TYPE_RSA = "RSA"
	SIGN_TYPE_SM2 = "SM2"
)

type RetCode string

const (
	RET_SUCCESS RetCode = "SUCCESS"
	RET_FAIl    RetCode = "FAIL"
)

/*
接口返回码retcode说明
接口版本号(version)为空或者版本号等于11：

SUCCESS-请求成功

FAIL-请求或前端处理失败，具体看retmsg

接口版本号(version)大于11：

SUCCESS-请求成功

PARAMERR-请求参数错误

SIGNAUTHERR-签名或者api权限不足

FAIL-其他请求或前端处理失败，具体看retmsg

SYSTEMERR-系统异常，对于实时类交易（例如被扫交易），建议进行查询
*/

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
type TrxStatus string

const (
	TRX_0000 TrxStatus = "0000"
	TRX_1001 TrxStatus = "1001"
	TRX_3888 TrxStatus = "3888"
	TRX_3050 TrxStatus = "3050"
)

/*
交易类型	注释
VSP501	微信支付
VSP502	微信支付撤销
VSP503	微信支付退款
VSP505	手机QQ 支付
VSP506	手机QQ支付撤销
VSP507	手机QQ支付退款
VSP511	支付宝支付
VSP512	支付宝支付撤销
VSP513	支付宝支付退款
VSP541	扫码支付
VSP542	扫码撤销
VSP543	扫码退货
VSP551	银联扫码支付
VSP552	银联扫码撤销
VSP553	银联扫码退货
VSP907	差错借记调整
VSP908	差错贷记调整
*/
type Vsp string

const (
	VSP501 Vsp = "VSP501" // 微信支付
)

/*
交易方式	注释
W01	微信扫码支付
W02	微信JS支付
W06	微信小程序支付
A01	支付宝扫码支付
A02	支付宝JS支付
A03	支付宝APP支付
Q01	手机QQ扫码支付
Q02	手机QQ JS支付
U01	银联扫码支付(CSB)
U02	银联JS支付
*/
type PayType string

const (
	PAY_TYPE_W01 PayType = "W01"
	PAY_TYPE_W02 PayType = "W02"
	PAY_TYPE_W06 PayType = "W06"
)
