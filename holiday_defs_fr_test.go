package cal

import (
	"testing"
	"time"
)

func TestFranceHolidays(t *testing.T) {
	c, err := newCalendarFromCountryCode("FR")
	if err != nil {
		t.Fatal(err)
	}

	assertHolidays(t, c, []testStruct{
		{time.Date(2017, 1, 1, 12, 0, 0, 0, time.UTC), true, "NouvelAn"},
		{time.Date(2017, 4, 17, 12, 0, 0, 0, time.UTC), true, "LundiDePâques"},
		{time.Date(2017, 5, 1, 12, 0, 0, 0, time.UTC), true, "FêteDuTravail"},
		{time.Date(2017, 5, 8, 12, 0, 0, 0, time.UTC), true, "Armistice1945"},
		{time.Date(2017, 5, 25, 12, 0, 0, 0, time.UTC), true, "JeudiDeLAscension"},
		{time.Date(2017, 6, 5, 12, 0, 0, 0, time.UTC), true, "LundiDePentecôte"},
		{time.Date(2017, 7, 14, 12, 0, 0, 0, time.UTC), true, "FêteNationale"},
		{time.Date(2017, 8, 15, 12, 0, 0, 0, time.UTC), true, "Assomption"},
		{time.Date(2017, 11, 1, 12, 0, 0, 0, time.UTC), true, "Toussaint"},
		{time.Date(2017, 11, 11, 12, 0, 0, 0, time.UTC), true, "Armistice1918"},
		{time.Date(2017, 12, 25, 12, 0, 0, 0, time.UTC), true, "Noël"},

		{time.Date(2018, 1, 1, 12, 0, 0, 0, time.UTC), true, "NouvelAn"},
		{time.Date(2018, 4, 2, 12, 0, 0, 0, time.UTC), true, "LundiDePâques"},
		{time.Date(2018, 5, 1, 12, 0, 0, 0, time.UTC), true, "FêteDuTravail"},
		{time.Date(2018, 5, 8, 12, 0, 0, 0, time.UTC), true, "Armistice1945"},
		{time.Date(2018, 5, 10, 12, 0, 0, 0, time.UTC), true, "JeudiDeLAscension"},
		{time.Date(2018, 5, 21, 12, 0, 0, 0, time.UTC), true, "LundiDePentecôte"},
		{time.Date(2018, 7, 14, 12, 0, 0, 0, time.UTC), true, "FêteNationale"},
		{time.Date(2018, 8, 15, 12, 0, 0, 0, time.UTC), true, "Assomption"},
		{time.Date(2018, 11, 1, 12, 0, 0, 0, time.UTC), true, "Toussaint"},
		{time.Date(2018, 11, 11, 12, 0, 0, 0, time.UTC), true, "Armistice1918"},
		{time.Date(2018, 12, 25, 12, 0, 0, 0, time.UTC), true, "Noël"},
	})
}

func TestElsassHolidays(t *testing.T) {
	for _, code := range []string{"FR-67", "FR-68", "FR-57", "FR-75"} {
		expected := code != "FR-75"

		c, err := NewLocalCalendar(code)
		if err != nil {
			t.Error(err)
			continue
		}

		assertHolidays(t, c, []testStruct{
			{time.Date(2019, time.April, 19, 0, 0, 0, 0, time.UTC), expected, "Vendredi Saint"},
			{time.Date(2020, time.December, 26, 0, 0, 0, 0, time.UTC), expected, "Saint Étienne"},
		})
	}
}
