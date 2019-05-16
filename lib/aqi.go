package lib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type AirQuality struct{
	Aqi       int  `json:"aqi"`
	Wgt		  string `json:"wgt"`
}

func IsSameAqi(record int, aqi int) bool {
	result := true

	switch {
	case record > 300 && aqi > 300 :
	case record > 201 && aqi > 201 && record <= 300 && aqi <= 300 :
	case record > 151 && aqi > 151 && record <= 201 && aqi <= 201 :
	case record > 101 && aqi > 101 && record <= 151 && aqi <= 151 :
	case record > 51 && aqi > 51 && record <= 101 && aqi <= 101 :
	case record > 0 && aqi > 0 && record <= 51 && aqi <= 51 :
	default:
		result = false
	}
	return result
}

func XAqi() (aqi int, wgt string ) {
	var aq map[string]interface{}
	resp, _ := http.Get("http://aqicn.org/aqicn/json/android/shanghai/json")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &aq)

	aqiq := int(aq["aqi"].(float64))
	wgtq := aq["wgt"].(string)

	return aqiq, wgtq
}
