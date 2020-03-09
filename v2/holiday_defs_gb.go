package cal

import "time"

// British holidays
var (
	gbNewYear       = newHolidayFunc(calculateNewYearsHoliday).setLabel("New year")
	gbGoodFriday    = goodFriday.setLabel("Good friday")
	gbEasterMonday  = easterMonday.setLabel("Easter monday")
	gbEarlyMay      = makeGBEarlyMay().setLabel("Early may")
	gbSpringHoliday = newHolidayFloat(time.May, time.Monday, -1).setLabel("Spring holiday")
	gbSummerHoliday = newHolidayFloat(time.August, time.Monday, -1).setLabel("Summer holiday")
	gbChristmasDay  = christmas.setLabel("Christmas day")
	gbBoxingDay     = christmas2.setLabel("Boxing day")
)

// addBritishHolidays adds all British holidays to the Calender
func addBritishHolidays(c *Calendar) {
	c.addHoliday(
		gbNewYear,
		gbGoodFriday,
		gbEasterMonday,
		gbEarlyMay,
		gbSpringHoliday,
		gbSummerHoliday,
		gbChristmasDay,
		gbBoxingDay,
	)
}

// NewYearsDay is the 1st of January unless the 1st is a Saturday or Sunday
// in which case it occurs on the following Monday.
func calculateNewYearsHoliday(year int, loc *time.Location) (time.Month, int) {
	day := time.Date(year, time.January, 1, 0, 0, 0, 0, loc)
	switch day.Weekday() {
	case time.Saturday:
		day = day.AddDate(0, 0, 2)
	case time.Sunday:
		day = day.AddDate(0, 0, 1)
	}
	return time.January, day.Day()
}

// Early May holiday in UK is usually 1st Monday in May, but in 2020 this is
// replaced by VE day on Fri 8th May
func makeGBEarlyMay() holiday {
	hol := newHolidayFloat(time.May, time.Monday, 1)
	hol.fn = func(year int, loc *time.Location) (month time.Month, day int) {
		if year == 2020 {
			// In 2020 force VE day
			return time.May, 8
		}
		// Else defer to floating holiday calculation
		return time.May, 0
	}
	return hol
}
