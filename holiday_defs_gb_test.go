package cal

import (
	"testing"
	"time"
)

func TestAddBritishHolidays(t *testing.T) {
	c, err := NewCalendarFromCountryCode("GB")
	if err != nil {
		t.Fatal(err)
	}

	assertHolidays(t, c, []testStruct{
		{time.Date(2017, time.January, 2, 12, 0, 0, 0, time.UTC), true, "New Year"},
		{time.Date(2017, time.April, 14, 12, 0, 0, 0, time.UTC), true, "Good Friday"},
		{time.Date(2017, time.April, 17, 12, 0, 0, 0, time.UTC), true, "Easter Monday"},
		{time.Date(2017, time.May, 1, 12, 0, 0, 0, time.UTC), true, "Early May"},
		{time.Date(2017, time.May, 29, 12, 0, 0, 0, time.UTC), true, "Spring"},
		{time.Date(2017, time.August, 28, 12, 0, 0, 0, time.UTC), true, "Summer"},
		{time.Date(2017, time.December, 25, 12, 0, 0, 0, time.UTC), true, "Christmas"},
		{time.Date(2017, time.December, 26, 12, 0, 0, 0, time.UTC), true, "Boxing Day"},
		{time.Date(2019, time.May, 6, 12, 0, 0, 0, time.UTC), true, "Early May"},
		{time.Date(2020, time.May, 8, 12, 0, 0, 0, time.UTC), true, "Early May"},
		{time.Date(2021, time.May, 3, 12, 0, 0, 0, time.UTC), true, "Early May"},
		{time.Date(2011, time.January, 3, 12, 0, 0, 0, time.UTC), true, "New Year"},
	})
}
