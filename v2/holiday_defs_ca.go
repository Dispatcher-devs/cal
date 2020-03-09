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
	caNewYear      = newYear.setLabel("New year")
	caGoodFriday   = goodFriday.setLabel("Good friday")
	caCanadaDay    = newHoliday(time.July, 1).setLabel("Canada day")
	caLabourDay    = newHolidayFloat(time.September, time.Monday, 1).setLabel("Labour day")
	caChristmasDay = christmas.setLabel("Christmas day")

	// Per-province holidays, some are on the same day but are not the same
	// holidays, entries are duplicated for future-proofing.

	// AB, BC (special case below), NB, ON, SK
	caFamilyDay = newHolidayFloat(time.February, time.Monday, 3).setLabel("Family day")

	// Acts as QC National Patriot's Day and Victoria Day for all the rest except NL.
	caVictoriaDay = newHolidayFloat(time.May, time.Monday, 3).setLabel("Victoria day")

	// NT, YT
	caAboriginalDay = newHoliday(time.June, 21).setLabel("Aboriginal day")

	// NL, QC, YT
	caDiscoveryDay = newHoliday(time.June, 24).setLabel("Discovery day")

	// Everyone except MB, ON, QC
	caRemembranceDay = newHoliday(time.November, 11).setLabel("Remembrance day")

	// AB, NB, NS, ON, PE
	caBoxingDay = christmas2.setLabel("Boxing day")

	// aka. Civic Holiday, Heritage Day, New Brunswick Day, Natal Day
	// Everyone except QC
	caCivicHoliday = newHolidayFloat(time.August, time.Monday, 1).setLabel("Civic holiday")

	// Everyone except NB, NL, PE
	caThanksgiving = newHolidayFloat(time.October, time.Monday, 2).setLabel("Thanksgiving")

	caBCFamilyDay         = newHolidayFactory(calculateBritishColumbiaFamilyDay).setLabel("Family day")
	caMBLouisRielDay      = newHolidayFloat(time.February, time.Monday, 3).setLabel("Louis Riel day")
	caNLOrangemensDay     = newHoliday(time.July, 12).setLabel("Orangemens Day")
	caNLSaintGeorgesDay   = newHoliday(time.April, 23).setLabel("Saint George's day")
	caNLSaintPatricksDay  = newHoliday(time.March, 17).setLabel("Saint Patrick's day")
	caNSHeritageDay       = newHolidayFloat(time.February, time.Monday, 3).setLabel("Heritage day")
	caPEEasterMonday      = easterMonday.setLabel("Easter monday")
	caPEGoldCupParadeDay  = newHolidayFloat(time.August, time.Friday, 3).setLabel("Gold Cup parade day")
	caPEIslanderDay       = newHolidayFloat(time.February, time.Monday, 3).setLabel("Islander day")
	caQCNationalHoliday   = newHoliday(time.June, 24).setLabel("National holiday")
	caQCPatriotsDay       = caVictoriaDay.setLabel("Patriot's day")
	caYTDiscoveryDay      = newHolidayFloat(time.August, time.Monday, 3).setLabel("Discovery day")
	caYTSaintJeanBaptiste = caDiscoveryDay.setLabel("Saint John the Baptist")
)

// addCanadianHolidays adds all Canadian holidays to the Calender
func addCanadianHolidays(c *Calendar) {
	c.addHoliday(
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
		c.addHoliday(
			caCivicHoliday,
			caFamilyDay,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-BC": // British Columbia
		c.addHoliday(
			caBCFamilyDay,
			caCivicHoliday,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-MB": // Manitoba
		c.addHoliday(
			caCivicHoliday,
			caMBLouisRielDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-NB": // New Brunswick
		c.addHoliday(
			caBoxingDay,
			caCivicHoliday,
			caFamilyDay,
			caRemembranceDay,
			caVictoriaDay,
		)
	case "CA-NL": // Newfoundland and Labrador
		c.addHoliday(
			caCivicHoliday,
			caDiscoveryDay,
			caNLOrangemensDay,
			caNLSaintGeorgesDay,
			caNLSaintPatricksDay,
			caRemembranceDay,
		)
	case "CA-NT": // Northwest Territories
		c.addHoliday(
			caAboriginalDay,
			caCivicHoliday,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-NS": // Nova Scotia
		c.addHoliday(
			caBoxingDay,
			caCivicHoliday,
			caNSHeritageDay,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-NU": // Nunavut
		c.addHoliday(
			caCivicHoliday,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-ON": // Ontario
		c.addHoliday(
			caBoxingDay,
			caCivicHoliday,
			caFamilyDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-PE": // Prince Edward Island
		c.addHoliday(
			caBoxingDay,
			caCivicHoliday,
			caPEEasterMonday,
			caPEGoldCupParadeDay,
			caPEIslanderDay,
			caRemembranceDay,
			caVictoriaDay,
		)
	case "CA-QC": // Quebec
		c.addHoliday(
			caDiscoveryDay,
			caQCNationalHoliday,
			caQCPatriotsDay,
			caThanksgiving,
		)
	case "CA-SK": // Saskatchewan
		c.addHoliday(
			caCivicHoliday,
			caFamilyDay,
			caRemembranceDay,
			caThanksgiving,
			caVictoriaDay,
		)
	case "CA-YT": // Yukon
		c.addHoliday(
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
func calculateBritishColumbiaFamilyDay(year int, loc *time.Location) holiday {
	switch {
	case year >= 2013 && year <= 2018:
		return newHolidayFloat(time.February, time.Monday, 2)
	case year >= 2019:
		return newHolidayFloat(time.February, time.Monday, 3)
	}

	return holiday{}
}
