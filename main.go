package main

import (
	"time"
	"tony/lib"
)

func main()  {
	today := time.Now().Day()

	loadDay := lib.LoadFile("day")
	loadAqi := lib.LoadFile("aqi")

	aqi, wgt := lib.XAqi()
	btc := lib.XBtc()
	say := lib.XSay()

	isNewday := lib.IsNewDay(loadDay, today)
	isSameaqi := lib.IsSameAqi(loadAqi, aqi)
	if isNewday || !isSameaqi {
		lib.WriteFile("aqi", aqi)
		lib.WriteFile("day", today)
		lib.Send(wgt, btc, say)
	}
}