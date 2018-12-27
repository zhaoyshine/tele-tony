package gadget

import (
	"net/http"
	"io/ioutil"
)

func GetSay() string {
	saying, _ := http.Get("https://v1.hitokoto.cn/?encode=text")
	defer saying.Body.Close()
	sayingBody, _ := ioutil.ReadAll(saying.Body)

	return string(sayingBody)
}