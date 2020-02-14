package cal

import (
	"time"
)

// Poland holidays
var (
	plNewYear                     = newYear
	plThreeKings                  = NewHoliday(time.January, 6)
	plEasterMonday                = easterMonday
	plEasterSunday                = NewHolidayFunc(calculateEasterSunday)
	plLabourDay                   = NewHoliday(time.May, 1)
	plNationalDay                 = NewHoliday(time.May, 3)
	plPentecostDay                = NewHolidayFunc(calculatePentecostDay)
	plCorpusChristi               = NewHolidayFunc(calculateCorpusChristiDay)
	plAssumptionBlessedVirginMary = NewHoliday(time.August, 15)
	plAllSaints                   = NewHoliday(time.November, 1)
	plNationalIndependenceDay     = NewHoliday(time.November, 11)
	plChristmasDayOne             = christmas
	plChristmasDayTwo             = christmas2
)

// addPolandHolidays adds all Poland holidays to the Calendar
func addPolandHolidays(c *Calendar) {
	c.AddHoliday(
		plNewYear,
		plThreeKings,
		plEasterMonday,
		plEasterSunday,
		plLabourDay,
		plNationalDay,
		plPentecostDay,
		plCorpusChristi,
		plAssumptionBlessedVirginMary,
		plAllSaints,
		plNationalIndependenceDay,
		plChristmasDayOne,
		plChristmasDayTwo,
	)
}

func calculateEasterSunday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	return easter.Month(), easter.Day()
}

func calculatePentecostDay(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +49)
	return em.Month(), em.Day()
}

func calculateCorpusChristiDay(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 60 days after Easter Sunday
	em := easter.AddDate(0, 0, +60)
	return em.Month(), em.Day()
}
