package gadget

import (
	"net/http"
	"io/ioutil"
	"github.com/henrylee2cn/pholcus/common/simplejson"
)

func GetSay() string {
	saying, _ := http.Get("https://v2.jinrishici.com/one.json")
	defer saying.Body.Close()
	sayingBody, _ := ioutil.ReadAll(saying.Body)
	js, err := simplejson.NewJson(sayingBody)
	if err != nil {
		panic(err.Error())
	}
	say := js.Get("data").Get("content").MustString()

	return say
}