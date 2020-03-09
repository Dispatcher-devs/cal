package cal

import "time"

// Holidays in the Netherlands
var (
	nlNieuwjaar       = newYear.SetLabel("Nieuwjaar")
	nlGoedeVrijdag    = goodFriday.SetLabel("GoedeVrijdag")
	nlPaasMaandag     = easterMonday.SetLabel("Paas maandag")
	nlKoningsDag      = newHolidayFunc(calculateKoningsDag).SetLabel("Konings dag")
	nlBevrijdingsDag  = newHoliday(time.May, 5).SetLabel("Bevrijdings Dag")
	nlHemelvaart      = deChristiHimmelfahrt.SetLabel("Hemelvaart")
	nlPinksterMaandag = dePfingstmontag.SetLabel("Pinkster maandag")
	nlEersteKerstdag  = christmas.SetLabel("Eerste kerstdag")
	nlTweedeKerstdag  = christmas2.SetLabel("Tweede kerstdag")
)

// addDutchHolidays adds all Dutch holidays to the Calendar
func addDutchHolidays(c *Calendar) {
	c.addHoliday(
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
