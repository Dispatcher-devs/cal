// (c) 2019 Paul Zeinlinger. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in Austria
var (
	atNeujahr            = newYear.setLabel("Neujahr")
	atHeiligeDreiKoenige = newHoliday(time.January, 6).setLabel("Heilige Drei Koenige")
	atOstermontag        = easterMonday.setLabel("Ostermontag")
	atTagderArbeit       = newHoliday(time.May, 1).setLabel("Tag der Arbeit")
	atChristiHimmelfahrt = newHolidayFunc(calculateHimmelfahrt).setLabel("Christi Himmelfahrt")
	atPfingstmontag      = newHolidayFunc(calculatePfingstMontag).setLabel("Pgingstmontag")
	atFronleichnam       = newHolidayFunc(calculateFronleichnam).setLabel("Fronleichnam")
	atMariaHimmelfahrt   = newHoliday(time.August, 15).setLabel("Maria Himmelfahrt")
	atNationalfeiertag   = newHoliday(time.October, 26).setLabel("Nationalfeiertag")
	atAllerheiligen      = newHoliday(time.November, 1).setLabel("Allerheiligen")
	atMariaEmpfaengnis   = newHoliday(time.December, 8).setLabel("Maria Emfaengnis")
	atChristtag          = christmas.setLabel("Christtag")
	atStefanitag         = christmas2.setLabel("Stefanitag")
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
