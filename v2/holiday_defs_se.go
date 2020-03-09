// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import "time"

// Holidays in Sweden
// Reference https://sv.wikipedia.org/wiki/Helgdagar_i_Sverige
// Days with the [2] notation, meaning days with reduced working hours
// haven't been added, as this is not regulated by law.
var (
	seNyarsdagen           = newYear.setLabel("Nyarsdagen")
	seTrettondedagJul      = newHoliday(time.January, 6).setLabel("Trettondedag Jul")
	seLangfredagen         = goodFriday.setLabel("Langfredagen")
	sePaskdagen            = newHolidayFunc(calculatePaskdagen).setLabel("Paskdagen")
	seAnnandagPask         = easterMonday.setLabel("Annandag Pask")
	seForstaMaj            = newHoliday(time.May, 1).setLabel("Forsta Maj")
	seKristiHimmelfardsdag = newHolidayFunc(calculateKristiHimmelfardsdag).setLabel("Kristi Himmelfardsdag")
	sePingstdagen          = newHolidayFunc(calculatePingstdagen).setLabel("Pingstdagen")
	seNationaldagen        = newHoliday(time.June, 6).setLabel("Nationaldagen")
	seMidsommarafton       = newHolidayFunc(calculateMidsommarafton).setLabel("Midsommarafton")
	seMidsommardagen       = newHolidayFunc(calculateMidsommardagen).setLabel("Midsommardagen")
	seAllaHelgonsDag       = newHolidayFunc(calculateAllaHelgonsDag).setLabel("Alla helgons dag")
	seJulafton             = newHoliday(time.December, 24).setLabel("Julafton")
	seJuldagen             = christmas.setLabel("Juldagen")
	seAnnandagJul          = christmas2.setLabel("Annandag jul")
	seNewYearsEve          = newHoliday(time.December, 31)
)

// addSwedishHolidays adds all Swedish holidays to the Calendar
func addSwedishHolidays(c *Calendar) {
	c.addHoliday(
		seNyarsdagen,
		seTrettondedagJul,
		seLangfredagen,
		sePaskdagen,
		seAnnandagPask,
		seForstaMaj,
		seKristiHimmelfardsdag,
		sePingstdagen,
		seNationaldagen,
		seMidsommarafton,
		seMidsommardagen,
		seAllaHelgonsDag,
		seJulafton,
		seJuldagen,
		seAnnandagJul,
		seNewYearsEve,
	)
}

func calculatePaskdagen(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	return easter.Month(), easter.Day()
}

func calculateKristiHimmelfardsdag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 39 days after Easter Sunday
	em := easter.AddDate(0, 0, +39)
	return em.Month(), em.Day()
}

func calculatePingstdagen(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	em := easter.AddDate(0, 0, +49)
	return em.Month(), em.Day()
}

func calculateMidsommarafton(year int, loc *time.Location) (time.Month, int) {
	t := time.Date(year, 6, 25, 0, 0, 0, 0, loc)

	for i := -1; i > -6; i-- {
		d := t.Add(time.Hour * 24 * time.Duration(i))
		if d.Weekday() == time.Friday {
			t = d
			break
		}
	}
	return t.Month(), t.Day()
}

func calculateMidsommardagen(year int, loc *time.Location) (time.Month, int) {
	t := time.Date(year, 6, 26, 0, 0, 0, 0, loc)

	for i := -1; i > -6; i-- {
		d := t.Add(time.Hour * 24 * time.Duration(i))
		if d.Weekday() == time.Saturday {
			t = d
			break
		}
	}
	return t.Month(), t.Day()
}

func calculateAllaHelgonsDag(year int, loc *time.Location) (time.Month, int) {
	t := time.Date(year, 11, 6, 0, 0, 0, 0, loc)

	for i := -1; i > -7; i-- {
		d := t.Add(time.Hour * 24 * time.Duration(i))
		if d.Weekday() == time.Saturday {
			t = d
			break
		}
	}
	return t.Month(), t.Day()
}
