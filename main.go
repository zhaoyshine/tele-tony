package main

import (
	"time"
	"tony/lib"
)

func main()  {
	today := time.Now().Day()

	loadDay := 5
	loadAqi := 10

	aqi, wgt := lib.XAqi()
	btc := lib.XBtc()
	say := lib.XSay()

	isNewday := lib.IsNewDay(loadDay, today)
	isSameaqi := lib.IsSameAqi(loadAqi, aqi)

	if isNewday || !isSameaqi {
		lib.Send(wgt, btc, say)
	}
}