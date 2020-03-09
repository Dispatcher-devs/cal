package cal

import "time"

// US holidays
var (
	usNewYear      = newYear.setLabel("New year")
	usMLK          = newHolidayFloat(time.January, time.Monday, 3).setLabel("Martin Luther King Jr. day")
	usPresidents   = newHolidayFloat(time.February, time.Monday, 3).setLabel("Presidents day")
	usMemorial     = newHolidayFloat(time.May, time.Monday, -1).setLabel("Memorial day")
	usIndependence = newHoliday(time.July, 4).setLabel("Independence day")
	usLabor        = newHolidayFloat(time.September, time.Monday, 1).setLabel("Labor day")
	usColumbus     = newHolidayFloat(time.October, time.Monday, 2).setLabel("Columbus day")
	usVeterans     = newHoliday(time.November, 11).setLabel("Veteran's day")
	usThanksgiving = newHolidayFloat(time.November, time.Thursday, 4).setLabel("Thanksgiving")
	usChristmas    = christmas.setLabel("Christmas")
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
