package cal

import (
	"fmt"
	"time"
)

// Canadian holidays
// Source: https://en.wikipedia.org/wiki/Public_holidays_in_Canada
// Other source with conflicting data: https://www.timeanddate.com/holidays/canada/
// Wikipedia was chosen over timeanddate, an actual Canadian should go over the data.
var (
	// National holidays
	caNewYear      = newYear.SetLabel("New year")
	caGoodFriday   = goodFriday.SetLabel("Good friday")
	caCanadaDay    = NewHoliday(time.July, 1).SetLabel("Canada day")
	caLabourDay    = NewHolidayFloat(time.September, time.Monday, 1).SetLabel("Labour day")
	caChristmasDay = christmas.SetLabel("Christmas day")

	// Per-province holidays, some are on the same day but are not the same
	// holidays, entries are duplicated for future-proofing.

	// AB, BC (special case below), NB, ON, SK
	caFamilyDay = NewHolidayFloat(time.February, time.Monday, 3).SetLabel("Family day")

	// Acts as QC National Patriot's Day and Victoria Day for all the rest except NL.
	caVictoriaDay = NewHolidayFloat(time.May, time.Monday, 3).SetLabel("Victoria day")

	// NT, YT
	caAboriginalDay = NewHoliday(time.June, 21).SetLabel("Aboriginal day")

	// NL, QC, YT
	caDiscoveryDay = NewHoliday(time.June, 24).SetLabel("Discovery day")

	// Everyone except MB, ON, QC
	caRemembranceDay = NewHoliday(time.November, 11).SetLabel("Remembrance day")

	// AB, NB, NS, ON, PE
	caBoxingDay = christmas2.SetLabel("Boxing day")

	// aka. Civic Holiday, Heritage Day, New Brunswick Day, Natal Day
	// Everyone except QC
	caCivicHoliday = NewHolidayFloat(time.August, time.Monday, 1).SetLabel("Civic holiday")

	// Everyone except NB, NL, PE
	caThanksgiving = NewHolidayFloat(time.October, time.Monday, 2).SetLabel("Thanksgiving")

	caBCFamilyDay         = NewHolidayFactory(calculateBritishColumbiaFamilyDay).SetLabel("Family day")
	caMBLouisRielDay      = NewHolidayFloat(time.February, time.Monday, 3).SetLabel("Louis Riel day")
	caNLOrangemensDay     = NewHoliday(time.July, 12).SetLabel("Orangemens Day")
	caNLSaintGeorgesDay   = NewHoliday(time.April, 23).SetLabel("Saint George's day")
	caNLSaintPatricksDay  = NewHoliday(time.March, 17).SetLabel("Saint Patrick's day")
	caNSHeritageDay       = NewHolidayFloat(time.February, time.Monday, 3).SetLabel("Heritage day")
	caPEEasterMonday      = easterMonday.SetLabel("Easter monday")
	caPEGoldCupParadeDay  = NewHolidayFloat(time.August, time.Friday, 3).SetLabel("Gold Cup parade day")
	caPEIslanderDay       = NewHolidayFloat(time.February, time.Monday, 3).SetLabel("Islander day")
	caQCNationalHoliday   = NewHoliday(time.June, 24).SetLabel("National holiday")
	caQCPatriotsDay       = caVictoriaDay.SetLabel("Patriot's day")
	caYTDiscoveryDay      = NewHolidayFloat(time.August, time.Monday, 3).SetLabel("Discovery day")
	caYTSaintJeanBaptiste = caDiscoveryDay.SetLabel("Saint John the Baptist")
)

// addCanadianHolidays adds all Canadian holidays to the Calender
func addCanadianHolidays(c *Calendar) {
	c.AddHoliday(
		caBoxingDay,
		caNewYear,
		caGoodFriday,
		caCanadaDay,
		caLabourDay,
		caChristmasDay,
	)
}

// addCanadaProvinceHolidays adds Canadian state holidays to the calendar
func addCanadaProvinceHolidays(c *Calendar, province string) error { // nolint:funlen
	switch province {
	case "CA-AB": // Alberta
		c.AddHoliday(
			caCivicHoliday,
			caFamilyDay,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-BC": // British Columbia
		c.AddHoliday(
			caBCFamilyDay,
			caCivicHoliday,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-MB": // Manitoba
		c.AddHoliday(
			caCivicHoliday,
			caMBLouisRielDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-NB": // New Brunswick
		c.AddHoliday(
			caBoxingDay,
			caCivicHoliday,
			caFamilyDay,
			caRemembranceDay,
			caVictoriaDay,
		)
	case "CA-NL": // Newfoundland and Labrador
		c.AddHoliday(
			caCivicHoliday,
			caDiscoveryDay,
			caNLOrangemensDay,
			caNLSaintGeorgesDay,
			caNLSaintPatricksDay,
			caRemembranceDay,
		)
	case "CA-NT": // Northwest Territories
		c.AddHoliday(
			caAboriginalDay,
			caCivicHoliday,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-NS": // Nova Scotia
		c.AddHoliday(
			caBoxingDay,
			caCivicHoliday,
			caNSHeritageDay,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-NU": // Nunavut
		c.AddHoliday(
			caCivicHoliday,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-ON": // Ontario
		c.AddHoliday(
			caBoxingDay,
			caCivicHoliday,
			caFamilyDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-PE": // Prince Edward Island
		c.AddHoliday(
			caBoxingDay,
			caCivicHoliday,
			caPEEasterMonday,
			caPEGoldCupParadeDay,
			caPEIslanderDay,
			caRemembranceDay,
			caVictoriaDay,
		)
	case "CA-QC": // Quebec
		c.AddHoliday(
			caDiscoveryDay,
			caQCNationalHoliday,
			caQCPatriotsDay,
			caThanksgiving,
		)
	case "CA-SK": // Saskatchewan
		c.AddHoliday(
			caCivicHoliday,
			caFamilyDay,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-YT": // Yukon
		c.AddHoliday(
			caAboriginalDay,
			caCivicHoliday,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
			caYTDiscoveryDay,
			caYTSaintJeanBaptiste, // Saint Jean-Baptiste
		)
	default:
		return fmt.Errorf("unknown province %s", province)
	}

	return nil
}

// https://www.cbc.ca/news/canada/british-columbia/b-c-s-first-family-day-set-for-feb-11-2013-1.1240359
// https://www.cbc.ca/news/canada/british-columbia/b-c-family-day-moving-one-week-later-starting-in-2019-1.4528735
func calculateBritishColumbiaFamilyDay(year int, loc *time.Location) Holiday {
	switch {
	case year >= 2013 && year <= 2018:
		return NewHolidayFloat(time.February, time.Monday, 2)
	case year >= 2019:
		return NewHolidayFloat(time.February, time.Monday, 3)
	}

	return Holiday{}
}
