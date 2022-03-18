package time

import (
	"fmt"
	"time"
)

//"Asia/Jakarta"
func Get_date(timezone string) (dmy string) {
	loc, _ := time.LoadLocation(timezone)
	timeutcplus := time.Now().In(loc)

	dmy = fmt.Sprintf("%02d-%02d-%02d", timeutcplus.Day(), timeutcplus.Month(), timeutcplus.Year())

	return dmy
}

func Get_hms(timezone string) (hms string) {
	loc, _ := time.LoadLocation(timezone)
	timeutcplus := time.Now().In(loc)

	hms = fmt.Sprintf("%02d:%02d:%02d", timeutcplus.Hour(), timeutcplus.Minute(), timeutcplus.Second())

	return hms
}
