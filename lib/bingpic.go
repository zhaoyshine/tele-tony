package lib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type bingimage struct {
	Url string `json:"url"`
}

type BingResult struct{
	Image []bingimage `json:"images"`
}


func XBingpic() string {
	var b BingResult
	resp, _ := http.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &b)

	wgtq := "http://s.cn.bing.net" + b.Image[0].Url

	return wgtq
}
