package gadget

import (
	"net/http"
	"io/ioutil"
	"github.com/henrylee2cn/pholcus/common/simplejson"
	"strconv"
)

func GetBtc() string {
	resp, _ := http.Get("https://blockchain.info/ticker")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	js, err := simplejson.NewJson(body)
	if err != nil {
		panic(err.Error())
	}
	price := js.Get("CNY").Get("15m").MustFloat64()
	symbol := js.Get("CNY").Get("symbol").MustString()

	return symbol + strconv.FormatFloat(price, 'f', -1, 64)
}