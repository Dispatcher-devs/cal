package cal

import "time"

// Holidays in the Netherlands
var (
	nlNieuwjaar       = newYear
	nlGoedeVrijdag    = goodFriday
	nlPaasMaandag     = easterMonday
	nlKoningsDag      = NewHolidayFunc(calculateKoningsDag)
	nlBevrijdingsDag  = NewHoliday(time.May, 5)
	nlHemelvaart      = deChristiHimmelfahrt
	nlPinksterMaandag = dePfingstmontag
	nlEersteKerstdag  = christmas
	nlTweedeKerstdag  = christmas2
)

// addDutchHolidays adds all Dutch holidays to the Calendar
func addDutchHolidays(c *Calendar) {
	c.AddHoliday(
		nlNieuwjaar,
		nlGoedeVrijdag,
		nlPaasMaandag,
		nlKoningsDag,
		nlBevrijdingsDag,
		nlHemelvaart,
		nlPinksterMaandag,
		nlEersteKerstdag,
		nlTweedeKerstdag,
	)
}

// KoningsDag (kingsday) is April 27th, 26th if the 27th is a Sunday
func calculateKoningsDag(year int, loc *time.Location) (time.Month, int) {
	koningsDag := time.Date(year, time.April, 27, 0, 0, 0, 0, loc)
	if koningsDag.Weekday() == time.Sunday {
		koningsDag = koningsDag.AddDate(0, 0, -1)
	}
	return koningsDag.Month(), koningsDag.Day()
}
