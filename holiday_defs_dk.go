// (c) 2014, 2017 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays (official and traditional) in Denmark
// Reference https://da.wikipedia.org/wiki/Helligdag#Danske_helligdage
var (
	dkNytaarsdag           = newYear
	dkSkaertorsdag         = NewHolidayFunc(calculateSkaertorsdag)
	dkLangfredag           = goodFriday
	dkPaaskedag            = NewHolidayFunc(calculatePaaskedag)
	dkAndenPaaskedag       = easterMonday
	dkStoreBededag         = NewHolidayFunc(calculateStoreBededag)
	dkKristiHimmelfartsdag = NewHolidayFunc(calculateKristiHimmelfartsdag)
	dkPinsedag             = NewHolidayFunc(calculatePinsedag)
	dkAndenPinsedag        = NewHolidayFunc(calculateAndenPinsedag)
	dkGrundlovsdag         = NewHoliday(time.June, 5)
	dkJuleaften            = NewHoliday(time.December, 24)
	dkJuledag              = christmas
	dkAndenJuledag         = christmas2
	dkNytaarsaften         = NewHoliday(time.December, 31)
)

// addDanishHolidays adds all Danish holidays to the Calendar
func addDanishHolidays(c *Calendar) {
	c.AddHoliday(
		dkNytaarsdag,
		dkSkaertorsdag,
		dkLangfredag,
		dkPaaskedag,
		dkAndenPaaskedag,
		dkStoreBededag,
		dkKristiHimmelfartsdag,
		dkPinsedag,
		dkAndenPinsedag,
		dkJuledag,
		dkAndenJuledag,
	)
}

// AddDanishTraditions adds Grundlovsdag (Constitution Day), Christmas
// Eve, and New Years Eve which are not official holidays.
func AddDanishTraditions(c *Calendar) {
	c.AddHoliday(
		dkGrundlovsdag,
		dkJuleaften,
		dkNytaarsaften,
	)
}

func calculateSkaertorsdag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	em := easter.AddDate(0, 0, -3)
	// 3 days before Easter Sunday
	return em.Month(), em.Day()
}

func calculatePaaskedag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	return easter.Month(), easter.Day()
}

func calculateStoreBededag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 26 days after Easter Sunday
	em := easter.AddDate(0, 0, +26)
	return em.Month(), em.Day()
}

func calculateKristiHimmelfartsdag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 39 days after Easter Sunday
	em := easter.AddDate(0, 0, +39)
	return em.Month(), em.Day()
}

func calculatePinsedag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 49 days after Easter Sunday
	em := easter.AddDate(0, 0, +49)
	return em.Month(), em.Day()
}

func calculateAndenPinsedag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +50)
	return em.Month(), em.Day()
}
