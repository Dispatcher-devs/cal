// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in Norway
// Reference https://no.wikipedia.org/wiki/Helligdager_i_Norge
var (
	noFoersteNyttaarsdag   = newYear
	noSkjaertorsdag        = NewHolidayFunc(calculateSkaertorsdag)
	noLangfredag           = goodFriday
	noFoerstePaaskedag     = NewHolidayFunc(calculatePaskdagen)
	noAndrePaaskedag       = easterMonday
	noArbeiderenesdag      = NewHoliday(time.May, 1)
	noGrunnlovsdag         = NewHoliday(time.May, 17)
	noKristihimmelfartsdag = NewHolidayFunc(calculateKristiHimmelfardsdag)
	noFoerstePinsedag      = NewHolidayFunc(calculatePingstdagen)
	noAndrePinsedag        = NewHolidayFunc(calculateAndenPinsedag)
	noFoersteJuledag       = christmas
	noAndreJuledag         = christmas2
	// Half days
	noJulaften      = NewHoliday(time.December, 24)
	noNyttaarsaften = NewHoliday(time.December, 31)
)

// addNorwegianHolidays adds all Norwegian holidays to Calendar
func addNorwegianHolidays(c *Calendar) {
	c.AddHoliday(
		noFoersteNyttaarsdag,
		noSkjaertorsdag,
		noLangfredag,
		noFoerstePaaskedag,
		noAndrePaaskedag,
		noArbeiderenesdag,
		noGrunnlovsdag,
		noKristihimmelfartsdag,
		noFoerstePinsedag,
		noAndrePinsedag,
		noFoersteJuledag,
		noAndreJuledag,
	)
}

// AddNorwegianHalfDays are note holidays, but often practiced as a half-business day.
func AddNorwegianHalfDays(c *Calendar) {
	c.AddHoliday(
		noJulaften,
		noNyttaarsaften,
	)
}
