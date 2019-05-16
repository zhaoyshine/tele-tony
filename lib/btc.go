package lib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type btccontent struct {
	Price float64 `json:"15m"`
	Symbol string `json:"symbol"`
}

type BtcResult struct{
	CNY btccontent `json:"CNY"`
}

func XBtc() string {
	resp, _ := http.Get("https://blockchain.info/ticker")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var XB BtcResult
	json.Unmarshal(body, &XB)

	return XB.CNY.Symbol + strconv.FormatFloat(float64(XB.CNY.Price), 'f', 6, 64)

}