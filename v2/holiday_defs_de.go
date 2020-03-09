// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"fmt"
	"time"
)

// Holidays in Germany
var (
	deNeujahr                  = newYear.SetLabel("Neujahr")
	deHeiligeDreiKoenige       = newHoliday(time.January, 6).SetLabel("Heilige Drei Koenige")
	deInternationalerFrauentag = newHoliday(time.March, 8).SetLabel("International Frauentag")
	deKarFreitag               = goodFriday.SetLabel("Kar Freitag")
	deOstersonntag             = newHolidayFunc(calculateOstersonntag).SetLabel("Ostersonntag")
	deOstermontag              = easterMonday.SetLabel("Ostermontag")
	deTagderArbeit             = newHoliday(time.May, 1).SetLabel("Tag der Arbeit")
	deChristiHimmelfahrt       = newHolidayFunc(calculateHimmelfahrt).SetLabel("ChristiHimmelFahrt")
	dePfingstsonntag           = newHolidayFunc(calculatePfingstSonntag).SetLabel("Pfingstsonntag")
	dePfingstmontag            = newHolidayFunc(calculatePfingstMontag).SetLabel("Pfingstmontag")
	deFronleichnam             = newHolidayFunc(calculateFronleichnam).SetLabel("Fronleichnam")
	deMariaHimmelfahrt         = newHoliday(time.August, 15).SetLabel("Maria Himmelfahrt")
	deWeltkindertag            = newHoliday(time.September, 20).SetLabel("Weltkindertag")
	deTagderDeutschenEinheit   = newHoliday(time.October, 3).SetLabel("Tag der Deutschen Einheit")
	deReformationstag          = newHoliday(time.October, 31).SetLabel("Reformationstag")
	deReformationstag2017      = newHolidayExact(time.October, 31, 2017).SetLabel("Reformationstag")
	deAllerheiligen            = newHoliday(time.November, 1).SetLabel("AllerHeiligen")
	deBußUndBettag             = newHolidayFunc(calculateBußUndBettag).SetLabel("Buß- Und Bettag")
	deErsterWeihnachtstag      = christmas.SetLabel("Erster Weihnachtstag")
	deZweiterWeihnachtstag     = christmas2.SetLabel("Zweiter Weihnachtstag")
)

// addGermanHolidays adds all German holidays to the Calendar
func addGermanHolidays(c *Calendar) {
	c.addHoliday(
		deNeujahr,
		deKarFreitag,
		deOstermontag,
		deTagderArbeit,
		deChristiHimmelfahrt,
		dePfingstmontag,
		deTagderDeutschenEinheit,
		deErsterWeihnachtstag,
		deZweiterWeihnachtstag,
	)
}

// addGermanyStateHolidays adds german state holidays to the calendar
func addGermanyStateHolidays(c *Calendar, state string) error { // nolint:funlen
	switch state {
	case "DE-BB": // Brandenburg
		c.addHoliday(
			deOstersonntag,
			dePfingstsonntag,
			deReformationstag,
		)
	case "DE-BE": // Berlin
		c.addHoliday(
			deInternationalerFrauentag,
			deReformationstag2017,
		)
	case "DE-BW": // Baden-Württemberg
		c.addHoliday(
			deHeiligeDreiKoenige,
			deFronleichnam,
			deAllerheiligen,
			deReformationstag2017,
		)
	case "DE-BY": // Bayern
		c.addHoliday(
			deHeiligeDreiKoenige,
			deFronleichnam,
			deMariaHimmelfahrt,
			deAllerheiligen,
			deReformationstag2017,
		)
	case "DE-HB": // Bremen
		c.addHoliday(deReformationstag)
	case "DE-HE": // Hessen
		c.addHoliday(
			deOstersonntag,
			dePfingstsonntag,
			deFronleichnam,
			deReformationstag2017,
		)
	case "DE-HH": // Hamburg
		c.addHoliday(deReformationstag)
	case "DE-MV": // Mecklenburg-Vorpommern
		c.addHoliday(deReformationstag)
	case "DE-NI": // Niedersachsen
		c.addHoliday(deReformationstag)
	case "DE-NW": // Nordrhein-Westfalen
		c.addHoliday(
			deFronleichnam,
			deAllerheiligen,
			deReformationstag2017,
		)
	case "DE-RP": // Rheinland-Pfalz
		c.addHoliday(
			deFronleichnam,
			deAllerheiligen,
			deReformationstag2017,
		)
	case "DE-SN", "DE-SA": // Sachsen (keep wrong "SA" code for compatibility)
		c.addHoliday(
			deReformationstag,
			deBußUndBettag,
		)
	case "DE-SH": // Schleswig-Holstein
		c.addHoliday(deReformationstag)
	case "DE-SL": // Saarland
		c.addHoliday(
			deFronleichnam,
			deMariaHimmelfahrt,
			deAllerheiligen,
			deReformationstag2017,
		)
	case "DE-ST": // Sachen-Anhalt
		c.addHoliday(
			deHeiligeDreiKoenige,
			deReformationstag,
		)
	case "DE-TH": // Thüringen
		c.addHoliday(
			deWeltkindertag,
			deReformationstag,
		)
	default:
		return fmt.Errorf("unknown Länder %s", state)
	}

	return nil
}

func calculateOstersonntag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	return easter.Month(), easter.Day()
}

func calculateHimmelfahrt(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 39 days after Easter Sunday
	em := easter.AddDate(0, 0, +39)
	return em.Month(), em.Day()
}

func calculatePfingstSonntag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +49)
	return em.Month(), em.Day()
}

func calculatePfingstMontag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +50)
	return em.Month(), em.Day()
}

func calculateFronleichnam(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +60)
	return em.Month(), em.Day()
}

func calculateBußUndBettag(year int, loc *time.Location) (time.Month, int) {
	t := time.Date(year, 11, 23, 0, 0, 0, 0, loc)

	for i := -1; i > -10; i-- {
		d := t.Add(time.Hour * 24 * time.Duration(i))
		if d.Weekday() == time.Wednesday {
			t = d
			break
		}
	}
	return t.Month(), t.Day()
}
