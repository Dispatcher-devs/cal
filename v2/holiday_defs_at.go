// (c) 2019 Paul Zeinlinger. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in Austria
var (
	atNeujahr            = newYear.SetLabel("Neujahr")
	atHeiligeDreiKoenige = NewHoliday(time.January, 6).SetLabel("Heilige Drei Koenige")
	atOstermontag        = easterMonday.SetLabel("Ostermontag")
	atTagderArbeit       = NewHoliday(time.May, 1).SetLabel("Tag der Arbeit")
	atChristiHimmelfahrt = NewHolidayFunc(calculateHimmelfahrt).SetLabel("Christi Himmelfahrt")
	atPfingstmontag      = NewHolidayFunc(calculatePfingstMontag).SetLabel("Pgingstmontag")
	atFronleichnam       = NewHolidayFunc(calculateFronleichnam).SetLabel("Fronleichnam")
	atMariaHimmelfahrt   = NewHoliday(time.August, 15).SetLabel("Maria Himmelfahrt")
	atNationalfeiertag   = NewHoliday(time.October, 26).SetLabel("Nationalfeiertag")
	atAllerheiligen      = NewHoliday(time.November, 1).SetLabel("Allerheiligen")
	atMariaEmpfaengnis   = NewHoliday(time.December, 8).SetLabel("Maria Emfaengnis")
	atChristtag          = christmas.SetLabel("Christtag")
	atStefanitag         = christmas2.SetLabel("Stefanitag")
)

// addAustrianHolidays adds all Austrian holidays to the Calendar
func addAustrianHolidays(c *Calendar) {
	c.AddHoliday(
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
