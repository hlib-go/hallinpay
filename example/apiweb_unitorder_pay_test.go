package example

import (
	"github.com/hlib-go/hallinpay"
	"testing"
)

// 下单测试
func TestApiwebUnitorderPay(t *testing.T) {
	r, err := hallinpay.Pay(cfg, &hallinpay.PayParams{
		Trxamt:        "1",
		Reqsn:         "111111111122",
		Paytype:       hallinpay.PAY_TYPE_W06,
		Body:          "",
		Remark:        "",
		Validtime:     "",
		Acct:          "oSxaj4qF_DkjV2QU4GmS0FAaa6TU",
		NotifyUrl:     "",
		LimitPay:      "",
		SubAppid:      "wx4cf01a042bcc7599",
		GoodsTag:      "",
		Benefitdetail: "",
		FrontUrl:      "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)

}
