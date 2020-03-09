package cal

import (
	"fmt"
	"strconv"
	"time"
)

// Holidays in France
var (
	frNouvelAn          = newYear.setLabel("Nouvel an")
	frLundiDePâques     = easterMonday.setLabel("Lundi de pâques")
	frFêteDuTravail     = ecbLabourDay.setLabel("Fête du travail")
	frArmistice1945     = newHoliday(time.May, 8).setLabel("Armistice 1945")
	frJeudiDeLAscension = newHolidayFunc(calculateJeudiDeLAscension).setLabel("Jeudi de l'ascension")
	frLundiDePentecôte  = newHolidayFunc(calculateLundiDePentecôte).setLabel("Lundi de pentecôte")
	frFêteNationale     = newHoliday(time.July, 14).setLabel("Fête nationale")
	frAssomption        = newHoliday(time.August, 15).setLabel("Assomption")
	frToussaint         = newHoliday(time.November, 1).setLabel("Toussaint")
	frArmistice1918     = newHoliday(time.November, 11).setLabel("Armistice 1918")
	frNoël              = christmas.setLabel("Noël")

	// Two more days for the Alsace-Moselle area.
	// https://www.legifrance.gouv.fr/affichCode.do?idSectionTA=LEGISCTA000006178008&cidTexte=LEGITEXT000006072050
	frSaintÉtienne  = newHoliday(time.December, 26).setLabel("Saint Étienne")
	frVendrediSaint = goodFriday.setLabel("Vendredi saint")

	// The end of slavery is observed at different dates for different DOM-TOMs.
	frGuadeloupeAbolition      = newHoliday(time.May, 27).setLabel("Abolition")
	frMartiniqueAbolition      = newHoliday(time.May, 22).setLabel("Abolition")
	frMayotteAbolition         = newHoliday(time.April, 27).setLabel("Abolition")
	frRéunionAbolition         = newHoliday(time.December, 20).setLabel("Abolition")
	frSaintBarthélemyAbolition = newHoliday(time.October, 9).setLabel("Abolition")
	frSaintMartinAbolition     = newHolidayFunc(calculateSaintMartinAbolition).setLabel("Abolition")
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
	c.addHoliday(
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
		c.addHoliday(frGuadeloupeAbolition)
		return nil
	case "FR-MQ":
		c.addHoliday(frMartiniqueAbolition)
		return nil
	case "FR-RE":
		c.addHoliday(frRéunionAbolition)
		return nil
	case "FR-YT":
		c.addHoliday(frMayotteAbolition)
		return nil
	case "FR-BL":
		c.addHoliday(frSaintBarthélemyAbolition)
		return nil
	case "FR-MF":
		c.addHoliday(frSaintMartinAbolition)
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
		c.addHoliday(
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
