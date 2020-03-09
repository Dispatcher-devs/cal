// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in Norway
// Reference https://no.wikipedia.org/wiki/Helligdager_i_Norge
var (
	noFoersteNyttaarsdag   = newYear.SetLabel("Foerste nyttaarsdag")
	noSkjaertorsdag        = newHolidayFunc(calculateSkaertorsdag).SetLabel("Skjaertorsdag")
	noLangfredag           = goodFriday.SetLabel("Langfredag")
	noFoerstePaaskedag     = newHolidayFunc(calculatePaskdagen).SetLabel("Foerste paaskedag")
	noAndrePaaskedag       = easterMonday.SetLabel("Andre Paaskedag")
	noArbeiderenesdag      = newHoliday(time.May, 1).SetLabel("Arbeiderenesdag")
	noGrunnlovsdag         = newHoliday(time.May, 17).SetLabel("Grunnlovsdag")
	noKristihimmelfartsdag = newHolidayFunc(calculateKristiHimmelfardsdag).SetLabel("Kristihimmelfartsdag")
	noFoerstePinsedag      = newHolidayFunc(calculatePingstdagen).SetLabel("Foerste pinsedag")
	noAndrePinsedag        = newHolidayFunc(calculateAndenPinsedag).SetLabel("Andre pinsedag")
	noFoersteJuledag       = christmas.SetLabel("Foerste juledag")
	noAndreJuledag         = christmas2.SetLabel("Andre juledag")
	// Half days
	noJulaften      = newHoliday(time.December, 24).SetLabel("Julaften")
	noNyttaarsaften = newHoliday(time.December, 31).SetLabel("Nyttaarsaften")
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
