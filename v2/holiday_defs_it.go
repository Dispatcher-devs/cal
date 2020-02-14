package cal

import "time"

// Italian holidays
var (
	itCapodanno             = newYear
	itEpifania              = NewHoliday(time.January, 6)
	itPasquetta             = easterMonday
	itFestaDellaLiberazione = NewHoliday(time.April, 25)
	itFestaDelLavoro        = NewHoliday(time.May, 1)
	itFestaDellaRepubblica  = NewHoliday(time.June, 2)
	itFerragosto            = NewHoliday(time.August, 15)
	itTuttiISanti           = NewHoliday(time.November, 1)
	itImmacolata            = NewHoliday(time.December, 8)
	itNatale                = christmas
	itSantoStefano          = christmas2
)

// addItalianHolidays adds all Italian holidays to the Calendar
func addItalianHolidays(c *Calendar) {
	c.AddHoliday(
		itCapodanno,
		itEpifania,
		itPasquetta,
		itFestaDellaLiberazione,
		itFestaDelLavoro,
		itFestaDellaRepubblica,
		itFerragosto,
		itTuttiISanti,
		itImmacolata,
		itNatale,
		itSantoStefano,
	)
}
