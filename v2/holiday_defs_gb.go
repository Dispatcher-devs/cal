package cal

import "time"

// British holidays
var (
	gbNewYear       = NewHolidayFunc(calculateNewYearsHoliday).SetLabel("New year")
	gbGoodFriday    = goodFriday.SetLabel("Good friday")
	gbEasterMonday  = easterMonday.SetLabel("Easter monday")
	gbEarlyMay      = makeGBEarlyMay().SetLabel("Early may")
	gbSpringHoliday = NewHolidayFloat(time.May, time.Monday, -1).SetLabel("Spring holiday")
	gbSummerHoliday = NewHolidayFloat(time.August, time.Monday, -1).SetLabel("Summer holiday")
	gbChristmasDay  = christmas.SetLabel("Christmas day")
	gbBoxingDay     = christmas2.SetLabel("Boxing day")
)

// addBritishHolidays adds all British holidays to the Calender
func addBritishHolidays(c *Calendar) {
	c.AddHoliday(
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
func makeGBEarlyMay() Holiday {
	hol := NewHolidayFloat(time.May, time.Monday, 1)
	hol.Func = func(year int, loc *time.Location) (month time.Month, day int) {
		if year == 2020 {
			// In 2020 force VE day
			return time.May, 8
		}
		// Else defer to floating holiday calculation
		return time.May, 0
	}
	return hol
}
