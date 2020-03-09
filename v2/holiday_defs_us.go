package cal

import "time"

// US holidays
var (
	usNewYear      = newYear.SetLabel("New year")
	usMLK          = newHolidayFloat(time.January, time.Monday, 3).SetLabel("Martin Luther King Jr. day")
	usPresidents   = newHolidayFloat(time.February, time.Monday, 3).SetLabel("Presidents day")
	usMemorial     = newHolidayFloat(time.May, time.Monday, -1).SetLabel("Memorial day")
	usIndependence = newHoliday(time.July, 4).SetLabel("Independence day")
	usLabor        = newHolidayFloat(time.September, time.Monday, 1).SetLabel("Labor day")
	usColumbus     = newHolidayFloat(time.October, time.Monday, 2).SetLabel("Columbus day")
	usVeterans     = newHoliday(time.November, 11).SetLabel("Veteran's day")
	usThanksgiving = newHolidayFloat(time.November, time.Thursday, 4).SetLabel("Thanksgiving")
	usChristmas    = christmas.SetLabel("Christmas")
)

// addUSHolidays adds all US holidays to the Calendar
func addUSHolidays(cal *Calendar) {
	cal.addHoliday(
		usNewYear,
		usMLK,
		usPresidents,
		usMemorial,
		usIndependence,
		usLabor,
		usColumbus,
		usVeterans,
		usThanksgiving,
		usChristmas,
	)
}
