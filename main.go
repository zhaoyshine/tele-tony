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
)

func IsToday(filename string, day int) bool {
	result := false

	record := fileOperation.ReadFile(filename)
	recordDate, _ := strconv.Atoi(record)
	if day == recordDate {
		result = true
	}
	fileOperation.WriteFile("/home/tele-tony/record/time", strconv.Itoa(day))

	return result
}

func SendMessage(say string)  {
	resp, err := http.Post("https://api.telegram.org/bot705617182:AAHyw5JrrlWCQf-D2l5X1fLtXJE8plJqtOU/sendMessage",
		"application/x-www-form-urlencoded",
		strings.NewReader("chat_id=-1001122390151&text="+ say))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	resBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(resBody))
}

func SendPhoto(wgt string)  {
	resp, err := http.Post("https://api.telegram.org/bot705617182:AAHyw5JrrlWCQf-D2l5X1fLtXJE8plJqtOU/sendPhoto",
		"application/x-www-form-urlencoded",
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
	isToday := IsToday("/home/tele-tony/record/time", today)
	same := gadget.Same("/home/tele-tony/record/data", aqi)

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
