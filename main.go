package main

import (
	"fmt"
	"time"
	"tony/lib"
)

func main() {
	today := time.Now().Day()

	loadDay := lib.LoadFile("/home/tele-tony/day")
	loadAqi := lib.LoadFile("/home/tele-tony/aqi")

	aqi, wgt := lib.XAqi()
	btc := lib.XBtc()
	say := lib.XSay()

	isNewday := lib.IsNewDay(loadDay, today)
	isSameaqi := lib.IsSameAqi(loadAqi, aqi)
	fmt.Println(isNewday)
	fmt.Println(isSameaqi)
	if isNewday || !isSameaqi {
		lib.WriteFile("/home/tele-tony/aqi", aqi)
		lib.WriteFile("/home/tele-tony/day", today)
		lib.Send(wgt, btc, say)
	}
}
