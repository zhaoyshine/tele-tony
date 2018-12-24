package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
	"net/http"
	"encoding/json"
	"strings"
)


type AirQuality struct{
	Aqi       int  `json:"aqi"`
}

func IsOverproof(filename string) bool {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.ParseBool(string(content))
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func writeFile(filename string)  {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("true")
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

	isOverproof := IsOverproof("data.txt")

	if isOverproof {
		writeFile("data.txt")
		http.Post("https://api.telegram.org/bot705617182:AAHyw5JrrlWCQf-D2l5X1fLtXJE8plJqtOU/sendMessage",
			"application/x-www-form-urlencoded",
			strings.NewReader("chat_id=-321414996&text=pm2.5超标了！已经达到 "+strconv.Itoa(air.Aqi)))
	}
}
