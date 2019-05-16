package main

import (
	"io/ioutil"
	"strconv"
	"net/http"
	"strings"
	"tele-tony/fileOperation"
	"fmt"
	"time"
	"tele-tony/gadget"
	"net/url"
)

func IsToday(filename string, day int) bool {
	result := false

	record := fileOperation.ReadFile(filename)
	recordDate, _ := strconv.Atoi(record)
	if day == recordDate {
		result = true
	}
	fileOperation.WriteFile("./record/time", strconv.Itoa(day))

	return result
}

func SendMessage(say string)  {
	resp, err := http.PostForm("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=15e669e8-0ba6-4269-8e25-9e4483256ed3",
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	resBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(resBody))
}

func SendPhoto(wgt string)  {
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=15e669e8-0ba6-4269-8e25-9e4483256ed3",
		"Content-Type: application/json",
		strings.NewReader("chat_id=-1001122390151&photo="+ wgt))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	resBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(resBody))
}

func main() {
	today := time.Now().Day()
	aqi, wgt := gadget.GetAqi()
	btc := gadget.GetBtc()
	isToday := IsToday("./record/time", today)
	same := gadget.Same("./record/data", aqi)

	if !isToday {
		saying := gadget.GetSay()
		say := "新的一天！现在的的pm2.5是" + strconv.Itoa(aqi) + "\n当前比特币价格为 " + btc + "。\n" + saying
		SendMessage(say)
	}
	if !same && isToday {
		SendPhoto(wgt)
		saying := gadget.GetSay()
		say := "当前比特币价格为 " + btc + "。\n" + saying
		SendMessage(say)
	}
}
