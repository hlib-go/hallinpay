package example

import (
	"github.com/hlib-go/hallinpay"
	"testing"
)

/*
time="2021-03-18T22:53:15+08:00" level=info msg="allinpay 请求URL:https://vsp.allinpay.com/apiweb/unitorder/pay" requestId=075faf70154aa7571232d6be6ce9542b
time="2021-03-18T22:53:15+08:00" level=info msg="allinpay 请求报文:{\"acct\":\"oSxaj4qF_DkjV2QU4GmS0FAaa6TU\",\"appid\":\"00209139\",\"body\":\"\",\"cusid\":\"55233207311VKJM\",\"paytype\":\"W06\",\"randomstr\":\"17d104e0fa708bb2b527276b6615a02c\",\"remark\":\"\",\"reqsn\":\"111111111122\",\"sign\":\"KYfUtnNDaTM5blu/Na85SGt2Wn70MZZ6OWs/pkfOUetKv7iHCJqy9iJVuU8IC2pSbrWg3440vF52BQgP8NcNzKqSzyuLOAQ7YsCq2ADtkvnYYJjUVmjHS6qZl5q3vamBYqAGsWQYCs/3bhOqL4HCDug/vDmBnBFK0LEli2bwOBI=\",\"signtype\":\"RSA\",\"sub_appid\":\"wx4cf01a042bcc7599\",\"trxamt\":\"1\",\"validtime\":\"\",\"version\":\"11\"}" requestId=075faf70154aa7571232d6be6ce9542b
time="2021-03-18T22:53:15+08:00" level=info msg="allinpay 响应报文:{\"appid\":\"00209139\",\"cusid\":\"55233207311VKJM\",\"payinfo\":\"{\\\"appId\\\":\\\"wx4cf01a042bcc7599\\\",\\\"timeStamp\\\":\\\"1616079195\\\",\\\"nonceStr\\\":\\\"f177ae5ca59f4cf2a97db7e5afd9a7ea\\\",\\\"package\\\":\\\"prepay_id=wx182253149385697a88bc6ecc03c3390000\\\",\\\"signType\\\":\\\"RSA\\\",\\\"paySign\\\":\\\"Tii6dVPDMTSCQLpbMM8RsEKdbUu6T+Pahk5ACHgt1u8ug0xrMevZNtR1oM6y3yZMnCQ1sfoUjAdU/jAV/IBb9JrneIXtozFA7LUvjJDdWKNIGJIAU1abvk+EBzCzpYGj4DkFrVuxVDr+egWEvPoWf2A9ZmXVISBi1p7djRU68Ttr0QfBrqHjvKLjcG3+KdHu4WbSzq01hBFvpnVLwcFHLP2I8MovovQPdcU6fmWfpHukegFiX6YT6rolFgwLiLvRPkB/F1bm4/OAJdN6fA1v9dph2FF6dWIqD32LKu1bYp8lo7ebakiTlnq2YiZ7NewDnRXY5u6FLJFHBuAASsKhGQ==\\\"}\",\"randomstr\":\"880807794642\",\"reqsn\":\"111111111122\",\"retcode\":\"SUCCESS\",\"sign\":\"eU1ZFJBmR3bqRkdGwcwYiE7HTbF6GQrmatFkrQbShKcvsOt+h9Jj+QyehosePm3r9aLAoELhHfkc+WcJmFfbRnDWRwnahvWWH1hDub3xVoKjBPXy9xteDAhL8VZ5o/x2Pi2H/DTsgPcp/0ObKD/IH/yAQzioak+vYea3iDmMgYc=\",\"trxid\":\"112181560000967098\",\"trxstatus\":\"0000\"}" ms=-5 requestId=075faf70154aa7571232d6be6ce9542b
*/

// 交易查询测试
func TestApiwebUnitorderQuery(t *testing.T) {
	r, err := hallinpay.Query(cfg, "111111111122", "112181560000967098")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
}
