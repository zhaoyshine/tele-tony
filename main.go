package main

import (
	"tony/lib"
	"strconv"
	"time"
)

func main() {
	aqi, _ := lib.XAqi()
	today := time.Now().Day()

	loadDay := lib.LoadFile("/home/tele-tony/day")
	loadAqi := lib.LoadFile("/home/tele-tony/aqi")

	isNewday := lib.IsNewDay(loadDay, today)
	isSameaqi := lib.IsSameAqi(loadAqi, aqi)

	if isNewday || !isSameaqi {

		btc := lib.XBtc()
		say := lib.XSay()
		wgt := lib.XBingpic()

		lib.WriteFile("/home/tele-tony/aqi", aqi)
		lib.WriteFile("/home/tele-tony/day", today)

		lib.Send(wgt, btc, say, strconv.Itoa(aqi))
	}
}
