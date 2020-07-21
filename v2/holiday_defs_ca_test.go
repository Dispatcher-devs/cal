package cal

import (
	"testing"
	"time"
)

func TestCanadianHolidays(t *testing.T) {
	c, err := newCalendarFromCountryCode("CA")
	if err != nil {
		t.Fatal(err)
	}

	assertHolidays(t, c, []testStruct{
		{time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC), true, "New Years"},
		{time.Date(2020, time.April, 10, 12, 0, 0, 0, time.UTC), true, "Good Friday"},
		{time.Date(2020, time.July, 1, 12, 0, 0, 0, time.UTC), true, "Canada Day"},
		{time.Date(2020, time.September, 7, 12, 0, 0, 0, time.UTC), true, "Labour Day"},
		{time.Date(2020, time.December, 25, 12, 0, 0, 0, time.UTC), true, "Christmas"},
	})
}

func TestCanadianStateHolidays(t *testing.T) {
	tests := []struct {
		state string
		tests []testStruct
	}{
		{
			"AB",
			[]testStruct{
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "AB Family Day"},
				{time.Date(2020, time.May, 18, 12, 0, 0, 0, time.UTC), true, "AB Victoria Day"},
				{time.Date(2020, time.August, 3, 12, 0, 0, 0, time.UTC), true, "AB Heritage Day"},
			},
		},
		{
			"BC",
			[]testStruct{
				{time.Date(2012, time.February, 13, 12, 0, 0, 0, time.UTC), false, "BC First Family day"},
				{time.Date(2012, time.February, 20, 12, 0, 0, 0, time.UTC), false, "BC Family day"},
				{time.Date(2013, time.February, 11, 12, 0, 0, 0, time.UTC), true, "BC Family day"},
				{time.Date(2018, time.February, 12, 12, 0, 0, 0, time.UTC), true, "BC Family day"},
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "BC Family day"},

				{time.Date(2019, time.August, 5, 12, 0, 0, 0, time.UTC), true, "BC Civic Holiday"},
				{time.Date(2020, time.August, 3, 12, 0, 0, 0, time.UTC), true, "BC Civic Holiday"},
			},
		},
		{
			"MB",
			[]testStruct{
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "MB Louis Riel"},
			},
		},
		{
			"NL",
			[]testStruct{
				{time.Date(2020, time.March, 17, 12, 0, 0, 0, time.UTC), true, "NL Saint Patrick's Day"},
				{time.Date(2020, time.April, 23, 12, 0, 0, 0, time.UTC), true, "NL Saint George's Day"},
				{time.Date(2020, time.July, 12, 12, 0, 0, 0, time.UTC), true, "NL Orangemen's Day"},
			},
		},
		{
			"NT",
			[]testStruct{
				{time.Date(2020, time.June, 21, 12, 0, 0, 0, time.UTC), true, "NT Aboriginal Day"},
			},
		},
		{
			"NS",
			[]testStruct{
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "NS Heritage Day"},
			},
		},
		{
			"PE",
			[]testStruct{
				{time.Date(2020, time.April, 13, 12, 0, 0, 0, time.UTC), true, "PE Easter Monday"},
				{time.Date(2020, time.August, 21, 12, 0, 0, 0, time.UTC), true, "PE Gold Cup Parade"},
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "PE Islander Day"},
			},
		},
		{
			"QC",
			[]testStruct{
				{time.Date(2023, time.June, 24, 12, 0, 0, 0, time.UTC), true, "QC National Holiday"},
				{time.Date(2020, time.August, 1, 12, 0, 0, 0, time.UTC), false, "QC no Civic Holiday"},
				{time.Date(2020, time.May, 18, 12, 0, 0, 0, time.UTC), true, "QC Patriot's Day"},
			},
		},
		{
			"YT",
			[]testStruct{
				{time.Date(2020, time.August, 17, 12, 0, 0, 0, time.UTC), true, "YT Discovery Day"},
				{time.Date(2020, time.June, 24, 12, 0, 0, 0, time.UTC), true, "YT Saint Jean Baptiste"},
			},
		},
	}

	for _, test := range tests {
		c, err := NewLocalCalendar("CA-" + test.state)
		if err != nil {
			t.Error(err)
			continue
		}

		assertHolidays(t, c, test.tests)
	}
}
