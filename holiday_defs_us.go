package cal

import "time"

// US holidays
var (
	usNewYear      = newYear
	usMLK          = NewHolidayFloat(time.January, time.Monday, 3)
	usPresidents   = NewHolidayFloat(time.February, time.Monday, 3)
	usMemorial     = NewHolidayFloat(time.May, time.Monday, -1)
	usIndependence = NewHoliday(time.July, 4)
	usLabor        = NewHolidayFloat(time.September, time.Monday, 1)
	usColumbus     = NewHolidayFloat(time.October, time.Monday, 2)
	usVeterans     = NewHoliday(time.November, 11)
	usThanksgiving = NewHolidayFloat(time.November, time.Thursday, 4)
	usChristmas    = christmas
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
