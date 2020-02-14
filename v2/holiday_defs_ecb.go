package cal

import "time"

// European Central Bank Target2 holidays
var (
	ecbGoodFriday       = goodFriday
	ecbEasterMonday     = easterMonday
	ecbNewYearsDay      = newYear
	ecbLabourDay        = NewHoliday(time.May, 1)
	ecbChristmasDay     = christmas
	ecbChristmasHoliday = christmas2
)

// AddECBHolidays adds all ECB Target2 holidays to the calendar
func AddECBHolidays(c *Calendar) {
	c.AddHoliday(
		ecbGoodFriday,
		ecbEasterMonday,
		ecbNewYearsDay,
		ecbLabourDay,
		ecbChristmasDay,
		ecbChristmasHoliday,
	)
}
