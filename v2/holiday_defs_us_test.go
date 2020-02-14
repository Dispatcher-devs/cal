package cal

import (
	"testing"
	"time"
)

func TestAddUSHolidays(t *testing.T) {
	c, err := newCalendarFromCountryCode("US")
	if err != nil {
		t.Fatal(err)
	}

	assertHolidays(t, c, []testStruct{
		{time.Date(2017, time.January, 1, 12, 0, 0, 0, time.UTC), true, "New Year"},
		{time.Date(2017, time.January, 16, 12, 0, 0, 0, time.UTC), true, "MLK"},
		{time.Date(2017, time.February, 20, 12, 0, 0, 0, time.UTC), true, "Presidents"},
		{time.Date(2017, time.May, 29, 12, 0, 0, 0, time.UTC), true, "Memorial"},
		{time.Date(2017, time.July, 4, 12, 0, 0, 0, time.UTC), true, "Independence"},
		{time.Date(2017, time.September, 4, 12, 0, 0, 0, time.UTC), true, "Labor"},
		{time.Date(2017, time.October, 9, 12, 0, 0, 0, time.UTC), true, "Columbus"},
		{time.Date(2017, time.November, 11, 12, 0, 0, 0, time.UTC), true, "Veterans"},
		{time.Date(2017, time.November, 23, 12, 0, 0, 0, time.UTC), true, "Thanksgiving"},
		{time.Date(2017, time.December, 25, 12, 0, 0, 0, time.UTC), true, "Christmas"},
	})
}
