package cal

import "time"

// US holidays
var (
	usNewYear      = newYear.SetLabel("New year")
	usMLK          = NewHolidayFloat(time.January, time.Monday, 3).SetLabel("Martin Luther King Jr. day")
	usPresidents   = NewHolidayFloat(time.February, time.Monday, 3).SetLabel("Presidents day")
	usMemorial     = NewHolidayFloat(time.May, time.Monday, -1).SetLabel("Memorial day")
	usIndependence = NewHoliday(time.July, 4).SetLabel("Independence day")
	usLabor        = NewHolidayFloat(time.September, time.Monday, 1).SetLabel("Labor day")
	usColumbus     = NewHolidayFloat(time.October, time.Monday, 2).SetLabel("Columbus day")
	usVeterans     = NewHoliday(time.November, 11).SetLabel("Veteran's day")
	usThanksgiving = NewHolidayFloat(time.November, time.Thursday, 4).SetLabel("Thanksgiving")
	usChristmas    = christmas.SetLabel("Christmas")
)

// addUSHolidays adds all US holidays to the Calendar
func addUSHolidays(cal *Calendar) {
	cal.AddHoliday(
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
