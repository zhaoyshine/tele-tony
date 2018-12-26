package main

import (
	"io/ioutil"
	"strconv"
	"time"
	"net/http"
	"encoding/json"
	"strings"
	"tele-tony/fileOperation"
)


type AirQuality struct{
	Aqi       int  `json:"aqi"`
}

func IsToday(filename string, day int) bool {
	result := true

	record := fileOperation.ReadFile(filename)
	recordDate, _ := strconv.Atoi(record)
	if day != recordDate {
		result = false
		fileOperation.WriteFile("record/time", strconv.Itoa(day))
	}

	return result
}

func Same(filename string, aqi int) bool {
	result := true

	record := fileOperation.ReadFile(filename)
	recordDate, _ := strconv.Atoi(record)

	switch {
	case recordDate > 300 && aqi > 300 :

	case recordDate > 201 && aqi > 201 && recordDate < 300 && aqi < 300 :

	case recordDate > 151 && aqi > 151 && recordDate < 201 && aqi < 201 :

	case recordDate > 101 && aqi > 101 && recordDate < 151 && aqi < 151 :

	case recordDate > 51 && aqi > 51 && recordDate < 101 && aqi < 101 :

	case recordDate > 0 && aqi > 0 && recordDate < 51 && aqi < 51 :

	default:
		result = false
	}

	fileOperation.WriteFile("record/data", strconv.Itoa(aqi))
	return result
}

func main() {
	today := time.Now().Day()
	isToday := IsToday("/home/tele-tony/record/time", today)
	//获取pm2.5的值
	resp, _ := http.Get("http://aqicn.org/aqicn/json/android/shanghai/json")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var air AirQuality
	json.Unmarshal(body, &air)
	//名言
	saying, _ := http.Get("https://v1.hitokoto.cn/?encode=text")
	defer saying.Body.Close()
	sayingBody, _ := ioutil.ReadAll(saying.Body)

	same := Same("/home/tele-tony/record/data", air.Aqi)

	if !isToday || !same {
		http.Post("https://api.telegram.org/bot705617182:AAHyw5JrrlWCQf-D2l5X1fLtXJE8plJqtOU/sendMessage",
			"application/x-www-form-urlencoded",
			strings.NewReader("chat_id=-321414996&text=pm2.5区间变动，目前是 "+strconv.Itoa(air.Aqi)+"。 \n"+string(sayingBody)))
	}
}
