package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"strconv"
)


type AirQuality struct{
	Aqi       int  `json:"aqi"`
}

func main() {
	resp, err := http.Get("http://aqicn.org/aqicn/json/android/shanghai/json")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var air AirQuality
	json.Unmarshal(body, &air)

	http.Post("https://api.telegram.org/bot705617182:AAHyw5JrrlWCQf-D2l5X1fLtXJE8plJqtOU/sendMessage",
		"application/x-www-form-urlencoded",
		strings.NewReader("chat_id=-321414996&text=pm2.5超标了！已经达到 "+strconv.Itoa(air.Aqi)))
}
