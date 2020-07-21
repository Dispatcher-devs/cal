package cal

import (
	"fmt"
	"strconv"
	"time"
)

// Holidays in France
var (
	frNouvelAn          = newYear.SetLabel("Nouvel an")
	frLundiDePâques     = easterMonday.SetLabel("Lundi de pâques")
	frFêteDuTravail     = ecbLabourDay.SetLabel("Fête du travail")
	frArmistice1945     = NewHoliday(time.May, 8).SetLabel("Armistice 1945")
	frJeudiDeLAscension = NewHolidayFunc(calculateJeudiDeLAscension).SetLabel("Jeudi de l'ascension")
	frLundiDePentecôte  = NewHolidayFunc(calculateLundiDePentecôte).SetLabel("Lundi de pentecôte")
	frFêteNationale     = NewHoliday(time.July, 14).SetLabel("Fête nationale")
	frAssomption        = NewHoliday(time.August, 15).SetLabel("Assomption")
	frToussaint         = NewHoliday(time.November, 1).SetLabel("Toussaint")
	frArmistice1918     = NewHoliday(time.November, 11).SetLabel("Armistice 1918")
	frNoël              = christmas.SetLabel("Noël")

	// Two more days for the Alsace-Moselle area.
	// https://www.legifrance.gouv.fr/affichCode.do?idSectionTA=LEGISCTA000006178008&cidTexte=LEGITEXT000006072050
	frSaintÉtienne  = NewHoliday(time.December, 26).SetLabel("Saint Étienne")
	frVendrediSaint = goodFriday.SetLabel("Vendredi saint")

	// The end of slavery is observed at different dates for different DOM-TOMs.
	frGuadeloupeAbolition      = NewHoliday(time.May, 27).SetLabel("Abolition")
	frMartiniqueAbolition      = NewHoliday(time.May, 22).SetLabel("Abolition")
	frMayotteAbolition         = NewHoliday(time.April, 27).SetLabel("Abolition")
	frRéunionAbolition         = NewHoliday(time.December, 20).SetLabel("Abolition")
	frSaintBarthélemyAbolition = NewHoliday(time.October, 9).SetLabel("Abolition")
	frSaintMartinAbolition     = NewHolidayFunc(calculateSaintMartinAbolition).SetLabel("Abolition")
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

// addFranceHolidays adds all France holidays to the Calendar
func addFranceHolidays(c *Calendar) {
	c.AddHoliday(
		frNouvelAn,
		frLundiDePâques,
		frFêteDuTravail,
		frArmistice1945,
		frJeudiDeLAscension,
		frLundiDePentecôte,
		frFêteNationale,
		frAssomption,
		frToussaint,
		frArmistice1918,
		frNoël,
	)
}

// https://en.wikipedia.org/wiki/ISO_3166-2:FR
func addFranceDepartmentHolidays(c *Calendar, code string) error {
	// DOM-TOM, non numeric codes
	switch code {
	case "FR-CP", "FR-GF", "FR-NC", "FR-PF", "FR-PM", "FR-TF", "FR-WF":
		return nil
	case "FR-GP":
		c.AddHoliday(frGuadeloupeAbolition)
		return nil
	case "FR-MQ":
		c.AddHoliday(frMartiniqueAbolition)
		return nil
	case "FR-RE":
		c.AddHoliday(frRéunionAbolition)
		return nil
	case "FR-YT":
		c.AddHoliday(frMayotteAbolition)
		return nil
	case "FR-BL":
		c.AddHoliday(frSaintBarthélemyAbolition)
		return nil
	case "FR-MF":
		c.AddHoliday(frSaintMartinAbolition)
		return nil
	}

	if len(code) != 5 { // FR-XX
		return fmt.Errorf("invalid ISO-3166-2 department code: %s", code)
	}

	dep, err := strconv.Atoi(code[3:])
	if err != nil {
		return fmt.Errorf("non-numeric ISO-3166-2 department code: %s", code[4:])
	}

	if dep < 1 || dep > 95 {
		return fmt.Errorf("out of bound department number: %d", dep)
	}

	// {Bas,Haut}-Rhin, Moselle
	if dep == 67 || dep == 68 || dep == 57 {
		c.AddHoliday(
			frSaintÉtienne,
			frVendrediSaint,
		)
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
