package example

import (
	"github.com/hlib-go/hallinpay"
	"testing"
)

// 退款交易测试
func TestRefund(t *testing.T) {
	r, err := hallinpay.Refund(cfg, &hallinpay.RefundParams{
		Trxamt:        "1",
		Reqsn:         "12345678",
		Oldreqsn:      "111111111122",
		Oldtrxid:      "112181560000967098",
		Remark:        "",
		Benefitdetail: "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
}
