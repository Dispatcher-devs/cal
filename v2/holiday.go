// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"time"
)

// ObservedRule represents a rule for observing a holiday that falls
// on a weekend.
type ObservedRule int

// ObservedRule are the specific ObservedRules
const (
	ObservedExact   ObservedRule = iota // the exact day only, default value
	ObservedNearest                     // nearest weekday (Friday or Monday)
	ObservedMonday                      // Monday always

	// As above, but also accounts for Christmas Day being on a weekend (which
	// pushes Boxing Day to Tuesday)
	ObservedMondayTuesday
)

// holidayFn calculates the occurrence of a holiday for the given year.
// This is useful for holidays like Easter that depend on complex rules.
type holidayFn func(year int, loc *time.Location) (month time.Month, day int)

// holidayFactory creates a holiday for the given year letting you use the
// regular Holiday/Cal facilities to define holidays that had yearly variations.
type holidayFactory func(year int, loc *time.Location) holiday

// Holiday holds information about the yearly occurrence of a holiday.
//
// A valid holiday consists of one of the following:
//   - Month and Day (such as March 14 for Pi Day)
//   - Month, Weekday, and Offset (such as the second Monday of October for Columbus Day)
//   - Offset (such as the 183rd day of the year for the start of the second half)
//   - Month, Day, and Year (in case you want to specify holidays exactly for each year)
//   - Func (to calculate the holiday)
type holiday struct {
	month   time.Month
	weekday time.Weekday
	day     int
	offset  int
	year    int
	fn      holidayFn
	factory holidayFactory
	label   string

	// last values used to calculate month and day with Func
	lastYear int
	lastLoc  *time.Location
}

func (h holiday) setLabel(label string) holiday {
	ret := h
	ret.label = label
	return ret
}

// newHoliday creates a new Holiday instance for an exact day of a month.
func newHoliday(month time.Month, day int) holiday {
	return holiday{month: month, day: day}
}

// newHolidayExact creates a new Holiday instance for an exact day of a month and year.
func newHolidayExact(month time.Month, day int, year int) holiday {
	return holiday{month: month, day: day, year: year}
}

// newHolidayFloat creates a new Holiday instance for an offset-based day of
// a month.
func newHolidayFloat(month time.Month, weekday time.Weekday, offset int) holiday {
	return holiday{month: month, weekday: weekday, offset: offset}
}

// newHolidayFunc creates a new Holiday instance that uses a function to
// calculate the day and month.
func newHolidayFunc(fn holidayFn) holiday {
	return holiday{fn: fn}
}

// newHolidayFactory creates a new holiday instance that uses a function to
// create the holiday for a given year.
func newHolidayFactory(fn holidayFactory) holiday {
	return holiday{factory: fn}
}

// matches determines whether the given date is the one referred to by the
// Holiday.
func (h *holiday) matches(date time.Time) bool {
	if h.factory != nil {
		generated := h.factory(date.Year(), date.Location())
		h.lastYear = date.Year()
		h.lastLoc = date.Location()
		return generated.matches(date)
	}

	if h.fn != nil && (date.Year() != h.lastYear || date.Location() != h.lastLoc) {
		h.month, h.day = h.fn(date.Year(), date.Location())
		h.lastYear = date.Year()
		h.lastLoc = date.Location()
	}

	if h.month > 0 {
		if date.Month() != h.month {
			return false
		}
		if h.year > 0 && date.Year() != h.year {
			return false
		}
		if h.day > 0 {
			return date.Day() == h.day
		}
		if h.weekday > 0 && h.offset != 0 {
			return IsWeekdayN(date, h.weekday, h.offset)
		}
	} else if h.offset > 0 {
		return date.YearDay() == h.offset
	}

	return false
}
