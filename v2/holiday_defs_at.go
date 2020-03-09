// (c) 2019 Paul Zeinlinger. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in Austria
var (
	atNeujahr            = newYear.SetLabel("Neujahr")
	atHeiligeDreiKoenige = newHoliday(time.January, 6).SetLabel("Heilige Drei Koenige")
	atOstermontag        = easterMonday.SetLabel("Ostermontag")
	atTagderArbeit       = newHoliday(time.May, 1).SetLabel("Tag der Arbeit")
	atChristiHimmelfahrt = newHolidayFunc(calculateHimmelfahrt).SetLabel("Christi Himmelfahrt")
	atPfingstmontag      = newHolidayFunc(calculatePfingstMontag).SetLabel("Pgingstmontag")
	atFronleichnam       = newHolidayFunc(calculateFronleichnam).SetLabel("Fronleichnam")
	atMariaHimmelfahrt   = newHoliday(time.August, 15).SetLabel("Maria Himmelfahrt")
	atNationalfeiertag   = newHoliday(time.October, 26).SetLabel("Nationalfeiertag")
	atAllerheiligen      = newHoliday(time.November, 1).SetLabel("Allerheiligen")
	atMariaEmpfaengnis   = newHoliday(time.December, 8).SetLabel("Maria Emfaengnis")
	atChristtag          = christmas.SetLabel("Christtag")
	atStefanitag         = christmas2.SetLabel("Stefanitag")
)

// addAustrianHolidays adds all Austrian holidays to the Calendar
func addAustrianHolidays(c *Calendar) {
	c.addHoliday(
		atNeujahr,
		atHeiligeDreiKoenige,
		atOstermontag,
		atTagderArbeit,
		atChristiHimmelfahrt,
		atPfingstmontag,
		atFronleichnam,
		atMariaHimmelfahrt,
		atNationalfeiertag,
		atAllerheiligen,
		atMariaEmpfaengnis,
		atChristtag,
		atStefanitag,
	)
}
