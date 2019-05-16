package lib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type saycontent struct {
	Content string `json:"content"`
}

type SayResult struct{
	Data saycontent `json:"data"`
}

func XSay() string {
	resp, _ := http.Get("https://v2.jinrishici.com/one.json")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var XS SayResult
	json.Unmarshal(body, &XS)

	return XS.Data.Content
}