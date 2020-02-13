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
	CANewYear      = NewYear
	CAGoodFriday   = GoodFriday
	CACanadaDay    = NewHoliday(time.July, 1)
	CALabourDay    = NewHolidayFloat(time.September, time.Monday, 1)
	CAChristmasDay = Christmas

	// Per-province holidays, some are on the same day but are not the same
	// holidays, entries are duplicated for future-proofing.

	// AB, BC (special case below), NB, ON, SK
	CAFamilyDay = NewHolidayFloat(time.February, time.Monday, 3)

	// Acts as QC National Patriot's Day and Victoria Day for all the rest except NL.
	CAVictoriaDay = NewHolidayFloat(time.May, time.Monday, 3)

	// NT, YT
	CAAboriginalDay = NewHoliday(time.June, 21)

	// NL, QC, YT
	CADiscoveryDay = NewHoliday(time.June, 24)

	// Everyone except MB, ON, QC
	CARemembranceDay = NewHoliday(time.November, 11)

	// AB, NB, NS, ON, PE
	CABoxingDay = Christmas2

	// aka. Civic Holiday, Heritage Day, New Brunswick Day, Natal Day
	// Everyone except QC
	CACivicHoliday = NewHolidayFloat(time.August, time.Monday, 1)

	// Everyone except NB, NL, PE
	CAThanksgiving = NewHolidayFloat(time.October, time.Monday, 2)

	CABCFamilyDay         = NewHolidayFactory(calculateBritishColumbiaFamilyDay)
	CAMBLouisRielDay      = NewHolidayFloat(time.February, time.Monday, 3)
	CANLOrangemensDay     = NewHoliday(time.July, 12)
	CANLSaintGeorgesDay   = NewHoliday(time.April, 23)
	CANLSaintPatricksDay  = NewHoliday(time.March, 17)
	CANSHeritageDay       = NewHolidayFloat(time.February, time.Monday, 3)
	CAPEEasterMonday      = EasterMonday
	CAPEGoldCupParadeDay  = NewHolidayFloat(time.August, time.Friday, 3)
	CAPEIslanderDay       = NewHolidayFloat(time.February, time.Monday, 3)
	CAQCNationalHoliday   = NewHoliday(time.June, 24)
	CAQCPatriotsDay       = CAVictoriaDay
	CAYTDiscoveryDay      = NewHolidayFloat(time.August, time.Monday, 3)
	CAYTSaintJeanBaptiste = CADiscoveryDay
)

// AddCanadianHolidays adds all Canadian holidays to the Calender
func AddCanadianHolidays(c *Calendar) {
	c.AddHoliday(
		CABoxingDay,
		CANewYear,
		CAGoodFriday,
		CACanadaDay,
		CALabourDay,
		CAChristmasDay,
	)
}

// AddCanadianProvinceHolidays adds Canadian state holidays to the calendar
func AddCanadianProvinceHolidays(c *Calendar, province string) error { // nolint:funlen
	switch province {
	case "AB": // Alberta
		c.AddHoliday(
			CACivicHoliday,
			CAFamilyDay,
			CARemembranceDay,
			CAThanksgiving,
			CAVictoriaDay,
		)
	case "BC": // British Columbia
		c.AddHoliday(
			CABCFamilyDay,
			CACivicHoliday,
			CARemembranceDay,
			CAThanksgiving,
			CAVictoriaDay,
		)
	case "MB": // Manitoba
		c.AddHoliday(
			CACivicHoliday,
			CAMBLouisRielDay,
			CAThanksgiving,
			CAVictoriaDay,
		)
	case "NB": // New Brunswick
		c.AddHoliday(
			CABoxingDay,
			CACivicHoliday,
			CAFamilyDay,
			CARemembranceDay,
			CAVictoriaDay,
		)
	case "NL": // Newfoundland and Labrador
		c.AddHoliday(
			CACivicHoliday,
			CADiscoveryDay,
			CANLOrangemensDay,
			CANLSaintGeorgesDay,
			CANLSaintPatricksDay,
			CARemembranceDay,
		)
	case "NT": // Northwest Territories
		c.AddHoliday(
			CAAboriginalDay,
			CACivicHoliday,
			CARemembranceDay,
			CAThanksgiving,
			CAVictoriaDay,
		)
	case "NS": // Nova Scotia
		c.AddHoliday(
			CABoxingDay,
			CACivicHoliday,
			CANSHeritageDay,
			CARemembranceDay,
			CAThanksgiving,
			CAVictoriaDay,
		)
	case "NU": // Nunavut
		c.AddHoliday(
			CACivicHoliday,
			CARemembranceDay,
			CAThanksgiving,
			CAVictoriaDay,
		)
	case "ON": // Ontario
		c.AddHoliday(
			CABoxingDay,
			CACivicHoliday,
			CAFamilyDay,
			CAThanksgiving,
			CAVictoriaDay,
		)
	case "PE": // Prince Edward Island
		c.AddHoliday(
			CABoxingDay,
			CACivicHoliday,
			CAPEEasterMonday,
			CAPEGoldCupParadeDay,
			CAPEIslanderDay,
			CARemembranceDay,
			CAVictoriaDay,
		)
	case "QC": // Quebec
		c.AddHoliday(
			CADiscoveryDay,
			CAQCNationalHoliday,
			CAQCPatriotsDay,
			CAThanksgiving,
		)
	case "SK": // Saskatchewan
		c.AddHoliday(
			CACivicHoliday,
			CAFamilyDay,
			CARemembranceDay,
			CAThanksgiving,
			CAVictoriaDay,
		)
	case "YT": // Yukon
		c.AddHoliday(
			CAAboriginalDay,
			CACivicHoliday,
			CARemembranceDay,
			CAThanksgiving,
			CAVictoriaDay,
			CAYTDiscoveryDay,
			CAYTSaintJeanBaptiste, // Saint Jean-Baptiste
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
