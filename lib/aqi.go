package lib

import (
	"tele-tony/fileOperation"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type AirQuality struct{
	Aqi       int  `json:"aqi"`
	Wgt		  string `json:"wgt"`
}

func Same(filename string, aqi int) bool {
	result := true
	record := fileOperation.ReadFile(filename)
	recordDate, _ := strconv.Atoi(record)

	switch {
	case recordDate > 300 && aqi > 300 :
	case recordDate > 201 && aqi > 201 && recordDate <= 300 && aqi <= 300 :
	case recordDate > 151 && aqi > 151 && recordDate <= 201 && aqi <= 201 :
	case recordDate > 101 && aqi > 101 && recordDate <= 151 && aqi <= 151 :
	case recordDate > 51 && aqi > 51 && recordDate <= 101 && aqi <= 101 :
	case recordDate > 0 && aqi > 0 && recordDate <= 51 && aqi <= 51 :
	default:
		result = false
	}

	fileOperation.WriteFile("./record/data", strconv.Itoa(aqi))
	return result
}

func GetAqi() (aqi int, wgt string ) {
	var aq map[string]interface{}
	resp, _ := http.Get("http://aqicn.org/aqicn/json/android/shanghai/json")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &aq)

	aqiq := int(aq["aqi"].(float64))
	wgtq := aq["wgt"].(string)

	return aqiq, wgtq
}
