package cal

import "time"

// Italian holidays
var (
	itCapodanno             = newYear.SetLabel("Capodanno")
	itEpifania              = newHoliday(time.January, 6).SetLabel("Epifania")
	itPasquetta             = easterMonday.SetLabel("Pasquetta")
	itFestaDellaLiberazione = newHoliday(time.April, 25).SetLabel("Festa della liberazione")
	itFestaDelLavoro        = newHoliday(time.May, 1).SetLabel("Festa del lavoro")
	itFestaDellaRepubblica  = newHoliday(time.June, 2).SetLabel("Festa della Repubblica")
	itFerragosto            = newHoliday(time.August, 15).SetLabel("Ferragosto")
	itTuttiISanti           = newHoliday(time.November, 1).SetLabel("Tutti i santi")
	itImmacolata            = newHoliday(time.December, 8).SetLabel("Immacolata")
	itNatale                = christmas.SetLabel("Natale")
	itSantoStefano          = christmas2.SetLabel("Santo Stefano")
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
