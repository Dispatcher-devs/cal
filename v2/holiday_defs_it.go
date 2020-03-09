package cal

import "time"

// Italian holidays
var (
	itCapodanno             = newYear.setLabel("Capodanno")
	itEpifania              = newHoliday(time.January, 6).setLabel("Epifania")
	itPasquetta             = easterMonday.setLabel("Pasquetta")
	itFestaDellaLiberazione = newHoliday(time.April, 25).setLabel("Festa della liberazione")
	itFestaDelLavoro        = newHoliday(time.May, 1).setLabel("Festa del lavoro")
	itFestaDellaRepubblica  = newHoliday(time.June, 2).setLabel("Festa della Repubblica")
	itFerragosto            = newHoliday(time.August, 15).setLabel("Ferragosto")
	itTuttiISanti           = newHoliday(time.November, 1).setLabel("Tutti i santi")
	itImmacolata            = newHoliday(time.December, 8).setLabel("Immacolata")
	itNatale                = christmas.setLabel("Natale")
	itSantoStefano          = christmas2.setLabel("Santo Stefano")
)

// addItalianHolidays adds all Italian holidays to the Calendar
func addItalianHolidays(c *Calendar) {
	c.addHoliday(
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
