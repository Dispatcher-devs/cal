package cal

import (
	"testing"
	"time"
)

func TestCanadianHolidays(t *testing.T) {
	c, err := NewCalendarFromCountryCode("CA")
	if err != nil {
		t.Fatal(err)
	}

	tests := []testStruct{
		{time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC), true, "New Years"},
		{time.Date(2020, time.April, 10, 12, 0, 0, 0, time.UTC), true, "Good Friday"},
		{time.Date(2020, time.July, 1, 12, 0, 0, 0, time.UTC), true, "Canada Day"},
		{time.Date(2020, time.September, 7, 12, 0, 0, 0, time.UTC), true, "Labour Day"},
		{time.Date(2020, time.December, 25, 12, 0, 0, 0, time.UTC), true, "Christmas"},
	}

	for _, test := range tests {
		got := c.IsHoliday(test.t)
		if got != test.want {
			t.Errorf("got: %t for %s; want: %t (%s)", got, test.name, test.want, test.t)
		}
	}
}

func TestCanadianStateHolidays(t *testing.T) {
	tests := []struct {
		state string
		tests []testStruct
	}{
		{
			"AB",
			[]testStruct{
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "Family   Day"},
				{time.Date(2020, time.May, 18, 12, 0, 0, 0, time.UTC), true, "Victoria Day"},
				{time.Date(2020, time.August, 3, 12, 0, 0, 0, time.UTC), true, "Heritage Day"},
			},
		},
		{
			"BC",
			[]testStruct{
				{time.Date(2012, time.February, 13, 12, 0, 0, 0, time.UTC), false, "First Family day"},
				{time.Date(2012, time.February, 20, 12, 0, 0, 0, time.UTC), false, "Family day"},
				{time.Date(2013, time.February, 11, 12, 0, 0, 0, time.UTC), true, "Family day"},
				{time.Date(2018, time.February, 12, 12, 0, 0, 0, time.UTC), true, "Family day"},
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "Family day"},

				{time.Date(2019, time.August, 5, 12, 0, 0, 0, time.UTC), true, "Civic Holiday"},
				{time.Date(2020, time.August, 3, 12, 0, 0, 0, time.UTC), true, "Civic Holiday"},
			},
		},
		{
			"MB",
			[]testStruct{
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "Louis Riel"},
			},
		},
		{
			"NL",
			[]testStruct{
				{time.Date(2020, time.March, 17, 12, 0, 0, 0, time.UTC), true, "Saint Patrick's Day"},
				{time.Date(2020, time.April, 23, 12, 0, 0, 0, time.UTC), true, "Saint George's Day"},
				{time.Date(2020, time.July, 12, 12, 0, 0, 0, time.UTC), true, "Orangemen's Day"},
			},
		},
		{
			"NT",
			[]testStruct{
				{time.Date(2020, time.June, 21, 12, 0, 0, 0, time.UTC), true, "Aboriginal Day"},
			},
		},
		{
			"NS",
			[]testStruct{
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "Heritage Day"},
			},
		},
		{
			"PE",
			[]testStruct{
				{time.Date(2020, time.April, 13, 12, 0, 0, 0, time.UTC), true, "Easter Monday"},
				{time.Date(2020, time.August, 21, 12, 0, 0, 0, time.UTC), true, "Gold Cup Parade"},
				{time.Date(2020, time.February, 17, 12, 0, 0, 0, time.UTC), true, "Islander Day"},
			},
		},
		{
			"QC",
			[]testStruct{
				{time.Date(2023, time.June, 24, 12, 0, 0, 0, time.UTC), true, "National holiday"},
				{time.Date(2020, time.August, 1, 12, 0, 0, 0, time.UTC), false, "no Civic Holiday"},
				{time.Date(2020, time.May, 18, 12, 0, 0, 0, time.UTC), true, "Patriot's Day"},
			},
		},
		{
			"YT",
			[]testStruct{
				{time.Date(2020, time.August, 17, 12, 0, 0, 0, time.UTC), true, "Discovery Day"},
				{time.Date(2020, time.June, 24, 12, 0, 0, 0, time.UTC), true, "Saint Jean Baptiste"},
			},
		},
	}

	for _, test := range tests {
		c, err := NewCalendarFromCountryCode("CA")
		if err != nil {
			t.Fatal(err)
		}

		if err := AddCanadianProvinceHolidays(c, test.state); err != nil {
			t.Fatal(err)
		}

		for _, day := range test.tests {
			got := c.IsHoliday(day.t)
			if got != day.want {
				t.Errorf("state: %s got: %t for %s; want: %t (%s)", test.state, got, day.name, day.want, day.t)
			}
		}
	}
}
