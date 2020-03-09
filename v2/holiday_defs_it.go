package cal

import "time"

// Italian holidays
var (
	itCapodanno             = newYear.SetLabel("Capodanno")
	itEpifania              = NewHoliday(time.January, 6).SetLabel("Epifania")
	itPasquetta             = easterMonday.SetLabel("Pasquetta")
	itFestaDellaLiberazione = NewHoliday(time.April, 25).SetLabel("Festa della liberazione")
	itFestaDelLavoro        = NewHoliday(time.May, 1).SetLabel("Festa del lavoro")
	itFestaDellaRepubblica  = NewHoliday(time.June, 2).SetLabel("Festa della Repubblica")
	itFerragosto            = NewHoliday(time.August, 15).SetLabel("Ferragosto")
	itTuttiISanti           = NewHoliday(time.November, 1).SetLabel("Tutti i santi")
	itImmacolata            = NewHoliday(time.December, 8).SetLabel("Immacolata")
	itNatale                = christmas.SetLabel("Natale")
	itSantoStefano          = christmas2.SetLabel("Santo Stefano")
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
