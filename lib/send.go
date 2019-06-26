package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

type news struct {
	Articles []Article `json:"articles"`
}

type Msg struct {
	Msgtype string `json:"msgtype"`
	News    news   `json:"news"`
}

func Send(wgt, btc, say, aqi string, isNewday bool) {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=15e669e8-0ba6-4269-8e25-9e4483256ed3"
	var msg Msg
	var article Article

	msg.Msgtype = "news"
	if isNewday {
		article.Title = "线上puma服务请升级到4.0解决部署问题"
	} else {
		article.Title = say
	}
	article.Description = "当前比特币价格为: " + btc + "\n当前空气质量为: " + aqi
	article.Url = "https://github.com/trending"
	article.Picurl = wgt
	msg.News.Articles = append(msg.News.Articles, article)

	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
