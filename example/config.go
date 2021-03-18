package example

import (
	"encoding/json"
	"github.com/hlib-go/hallinpay"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

var cfg *hallinpay.Config

func init() {
	b, err := ioutil.ReadFile("../.pwd/allinpay-55233207311VKJM.json")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		panic(err)
	}
}
