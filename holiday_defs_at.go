// (c) 2019 Paul Zeinlinger. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in Austria
var (
	atNeujahr            = newYear
	atHeiligeDreiKoenige = NewHoliday(time.January, 6)
	atOstermontag        = easterMonday
	atTagderArbeit       = NewHoliday(time.May, 1)
	atChristiHimmelfahrt = NewHolidayFunc(calculateHimmelfahrt)
	atPfingstmontag      = NewHolidayFunc(calculatePfingstMontag)
	atFronleichnam       = NewHolidayFunc(calculateFronleichnam)
	atMariaHimmelfahrt   = NewHoliday(time.August, 15)
	atNationalfeiertag   = NewHoliday(time.October, 26)
	atAllerheiligen      = NewHoliday(time.November, 1)
	atMariaEmpfaengnis   = NewHoliday(time.December, 8)
	atChristtag          = christmas
	atStefanitag         = christmas2
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
