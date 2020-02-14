package cal

import (
	"errors"
	"time"
)

// Holidays in France
var (
	FRNouvelAn          = NewYear
	FRLundiDePâques     = EasterMonday
	FRFêteDuTravail     = ECBLabourDay
	FRArmistice1945     = NewHoliday(time.May, 8)
	FRJeudiDeLAscension = NewHolidayFunc(calculateJeudiDeLAscension)
	FRLundiDePentecôte  = NewHolidayFunc(calculateLundiDePentecôte)
	FRFêteNationale     = NewHoliday(time.July, 14)
	FRAssomption        = NewHoliday(time.August, 15)
	FRToussaint         = NewHoliday(time.November, 1)
	FRArmistice1918     = NewHoliday(time.November, 11)
	FRNoël              = Christmas

	// Two more days for the Alsace-Moselle area.
	// https://www.legifrance.gouv.fr/affichCode.do?idSectionTA=LEGISCTA000006178008&cidTexte=LEGITEXT000006072050
	FRSaintÉtienne  = NewHoliday(time.December, 26)
	FRVendrediSaint = GoodFriday

	// The end of slavery is observed at different dates for different DOM-TOMs.
	FRGuadeloupeAbolition      = NewHoliday(time.May, 27)
	FRMartiniqueAbolition      = NewHoliday(time.May, 22)
	FRMayotteAbolition         = NewHoliday(time.April, 27)
	FRRéunionAbolition         = NewHoliday(time.December, 20)
	FRSaintBarthélemyAbolition = NewHoliday(time.October, 9)
	FRSaintMartinAbolition     = NewHolidayFunc(calculateSaintMartinAbolition)
)

func calculateSaintMartinAbolition(year int, loc *time.Location) (time.Month, int) {
	// nolint:misspell
	// cf. circulaire interministérielle du 3 mai 2019 relatives aux commémorations liées à l’esclavage
	// cf. loi n°2017-256 du 28 février 2017
	if year >= 2018 {
		return time.May, 28
	}

	return time.May, 27
}

// AddFranceHolidays adds all France holidays to the Calendar
func AddFranceHolidays(c *Calendar) {
	c.AddHoliday(
		FRNouvelAn,
		FRLundiDePâques,
		FRFêteDuTravail,
		FRArmistice1945,
		FRJeudiDeLAscension,
		FRLundiDePentecôte,
		FRFêteNationale,
		FRAssomption,
		FRToussaint,
		FRArmistice1918,
		FRNoël,
	)
}

func AddFranceDepartmentHolidays(c *Calendar, code string) error {
	if len(code) < 2 {
		return errors.New(`expected a postal code or a department number, eg. "91290", "03", or "972"`)
	}

	code2 := code[:2]
	if code2 == "67" || code2 == "68" || code2 == "57" {
		c.AddHoliday(
			FRSaintÉtienne,
			FRVendrediSaint,
		)
		return nil
	}

	if len(code) >= 3 {
		switch code[:3] {
		case "971":
			c.AddHoliday(FRGuadeloupeAbolition)
		case "972":
			c.AddHoliday(FRMartiniqueAbolition)
		case "974":
			c.AddHoliday(FRRéunionAbolition)
		case "976":
			c.AddHoliday(FRMayotteAbolition)
		case "977":
			c.AddHoliday(FRSaintBarthélemyAbolition)
		case "978":
			c.AddHoliday(FRSaintMartinAbolition)
		}
		return nil
	}

	return nil
}

func calculateJeudiDeLAscension(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 39 days after Easter Sunday
	t := easter.AddDate(0, 0, +39)
	return t.Month(), t.Day()
}

func calculateLundiDePentecôte(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	t := easter.AddDate(0, 0, +50)
	return t.Month(), t.Day()
}
