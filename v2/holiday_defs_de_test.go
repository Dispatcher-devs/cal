package cal

import (
	"testing"
	"time"
)

func TestGermanHolidays(t *testing.T) {
	c, err := newCalendarFromCountryCode("DE")
	if err != nil {
		t.Fatal(err)
	}

	assertHolidays(t, c, []testStruct{
		{time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC), true, "Neujahr"},
		{time.Date(2016, 3, 25, 12, 0, 0, 0, time.UTC), true, "Karfreitag"},
		{time.Date(2016, 3, 28, 12, 0, 0, 0, time.UTC), true, "Ostermontag"},
		{time.Date(2016, 5, 1, 12, 0, 0, 0, time.UTC), true, "Tag der Arbeit"},
		{time.Date(2016, 5, 5, 12, 0, 0, 0, time.UTC), true, "Himmelfahrt"},
		{time.Date(2000, 6, 1, 12, 0, 0, 0, time.UTC), true, "Himmelfahrt"},
		{time.Date(2016, 5, 16, 12, 0, 0, 0, time.UTC), true, "Pfingstmontag"},
		{time.Date(2000, 6, 12, 12, 0, 0, 0, time.UTC), true, "Pfingstmontag"},
		{time.Date(2016, 10, 3, 12, 0, 0, 0, time.UTC), true, "Tag der deutschen Einheit"},
		{time.Date(2016, 12, 25, 12, 0, 0, 0, time.UTC), true, "1. Weihnachtstag"},
		{time.Date(2016, 12, 26, 12, 0, 0, 0, time.UTC), true, "2. Weihnachtstag"},

		{time.Date(2017, 1, 1, 12, 0, 0, 0, time.UTC), true, "Neujahr"},
	})
}

func TestAddGermanyStateHolidays(t *testing.T) {
	tests := []struct {
		state string
		tests []testStruct
	}{
		{
			"BB",
			[]testStruct{
				{time.Date(2017, 4, 16, 12, 0, 0, 0, time.UTC), true, "DE Ostersonntag"},
				{time.Date(2017, 6, 4, 12, 0, 0, 0, time.UTC), true, "DE Pfingstsonntag"},
				{time.Date(2017, 10, 31, 12, 0, 0, 0, time.UTC), true, "DE Reformationstag"},
			},
		},
		{
			"BW",
			[]testStruct{
				{time.Date(2017, 1, 6, 12, 0, 0, 0, time.UTC), true, "BW Heilige Drei Könige"},
				{time.Date(2017, 6, 15, 12, 0, 0, 0, time.UTC), true, "BW Fronleichnam"},
				{time.Date(2017, 11, 1, 12, 0, 0, 0, time.UTC), true, "BW Allerheiligen"},
			},
		},
		{
			"BY",
			[]testStruct{
				{time.Date(2017, 1, 6, 12, 0, 0, 0, time.UTC), true, "BY Heilige Drei Könige"},
				{time.Date(2017, 6, 15, 12, 0, 0, 0, time.UTC), true, "BY Fronleichnam"},
				{time.Date(2017, 8, 15, 12, 0, 0, 0, time.UTC), true, "BY Mariä Himmelfahrt"},
				{time.Date(2017, 11, 1, 12, 0, 0, 0, time.UTC), true, "BY Allerheiligen"},
			},
		},
		{
			"HE",
			[]testStruct{
				{time.Date(2017, 6, 15, 12, 0, 0, 0, time.UTC), true, "HE Fronleichnam"},
			},
		},
		{
			"MV",
			[]testStruct{
				{time.Date(2017, 10, 31, 12, 0, 0, 0, time.UTC), true, "MV Reformationstag"},
			},
		},
		{
			"NW",
			[]testStruct{
				{time.Date(2017, 6, 15, 12, 0, 0, 0, time.UTC), true, "NW Fronleichnam"},
				{time.Date(2017, 11, 1, 12, 0, 0, 0, time.UTC), true, "NW Allerheiligen"},
			},
		},
		{
			"RP",
			[]testStruct{
				{time.Date(2017, 6, 15, 12, 0, 0, 0, time.UTC), true, "RP Fronleichnam"},
				{time.Date(2017, 11, 1, 12, 0, 0, 0, time.UTC), true, "RP Allerheiligen"},
			},
		},
		{
			"SA",
			[]testStruct{
				{time.Date(2017, 10, 31, 12, 0, 0, 0, time.UTC), true, "SA Reformationstag"},
				{time.Date(2017, 11, 22, 12, 0, 0, 0, time.UTC), true, "SA Buß- und Bettag"},
			},
		},
		{
			"SL",
			[]testStruct{
				{time.Date(2017, 6, 15, 12, 0, 0, 0, time.UTC), true, "SL Fronleichnam"},
				{time.Date(2017, 8, 15, 12, 0, 0, 0, time.UTC), true, "SL Mariä Himmelfahrt"},
				{time.Date(2017, 11, 1, 12, 0, 0, 0, time.UTC), true, "SL Allerheiligen"},
			},
		},
		{
			"ST",
			[]testStruct{
				{time.Date(2017, 1, 6, 12, 0, 0, 0, time.UTC), true, "ST Heilige Drei Könige"},
				{time.Date(2017, 10, 31, 12, 0, 0, 0, time.UTC), true, "ST Reformationstag"},
			},
		},
		{
			"TH",
			[]testStruct{
				{time.Date(2017, 10, 31, 12, 0, 0, 0, time.UTC), true, "TH Reformationstag"},
			},
		},
	}

	for _, test := range tests {
		c, err := NewLocalCalendar("DE-" + test.state)
		if err != nil {
			t.Error(err)
			continue
		}

		assertHolidays(t, c, test.tests)
	}
}
