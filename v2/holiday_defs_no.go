// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in Norway
// Reference https://no.wikipedia.org/wiki/Helligdager_i_Norge
var (
	noFoersteNyttaarsdag   = newYear.setLabel("Foerste nyttaarsdag")
	noSkjaertorsdag        = newHolidayFunc(calculateSkaertorsdag).setLabel("Skjaertorsdag")
	noLangfredag           = goodFriday.setLabel("Langfredag")
	noFoerstePaaskedag     = newHolidayFunc(calculatePaskdagen).setLabel("Foerste paaskedag")
	noAndrePaaskedag       = easterMonday.setLabel("Andre Paaskedag")
	noArbeiderenesdag      = newHoliday(time.May, 1).setLabel("Arbeiderenesdag")
	noGrunnlovsdag         = newHoliday(time.May, 17).setLabel("Grunnlovsdag")
	noKristihimmelfartsdag = newHolidayFunc(calculateKristiHimmelfardsdag).setLabel("Kristihimmelfartsdag")
	noFoerstePinsedag      = newHolidayFunc(calculatePingstdagen).setLabel("Foerste pinsedag")
	noAndrePinsedag        = newHolidayFunc(calculateAndenPinsedag).setLabel("Andre pinsedag")
	noFoersteJuledag       = christmas.setLabel("Foerste juledag")
	noAndreJuledag         = christmas2.setLabel("Andre juledag")
	// Half days
	noJulaften      = newHoliday(time.December, 24).setLabel("Julaften")
	noNyttaarsaften = newHoliday(time.December, 31).setLabel("Nyttaarsaften")
)

// addNorwegianHolidays adds all Norwegian holidays to Calendar
func addNorwegianHolidays(c *Calendar) {
	c.addHoliday(
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
	c.addHoliday(
		noJulaften,
		noNyttaarsaften,
	)
}
