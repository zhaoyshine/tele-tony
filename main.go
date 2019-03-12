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
	fileOperation.WriteFile("./record/time", strconv.Itoa(day))

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
	fmt.Println("------------------1")
	aqi, wgt := gadget.GetAqi()
	fmt.Println("------------------2")
	btc := gadget.GetBtc()
	fmt.Println("------------------3")
	isToday := IsToday("./record/time", today)
	fmt.Println("------------------4")
	same := gadget.Same("./record/data", aqi)
	fmt.Println("------------------5")

	if !isToday {
		fmt.Println("------------------6")
		saying := gadget.GetSay()
		fmt.Println("------------------7")
		say := "新的一天！现在的的pm2.5是" + strconv.Itoa(aqi) + "\n当前比特币价格为 " + btc + "。\n" + saying
		fmt.Println("------------------8")
		SendMessage(say)
		fmt.Println("------------------9")
	}
	if !same && isToday {
		fmt.Println("------------------10")
		SendPhoto(wgt)
		fmt.Println("------------------11")
		saying := gadget.GetSay()
		fmt.Println("------------------12")
		say := "当前比特币价格为 " + btc + "。\n" + saying
		fmt.Println("------------------13")
		SendMessage(say)
		fmt.Println("------------------14")
	}
}
