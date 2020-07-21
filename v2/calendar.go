// Package cal provides extensions to the time package for handling holidays
// and work days (business days).
//
// Work days are defined as Monday through Friday on dates that are not
// holidays or the alternate observation dates for holidays.
//
// As in the time package, all calculations assume a Gregorian calendar.
//
// (c) 2014 Rick Arnold. Licensed under the BSD license (see LICENSE).
package cal

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// IsWeekend reports whether the given date falls on a weekend.
func IsWeekend(date time.Time) bool {
	day := date.Weekday()
	return day == time.Saturday || day == time.Sunday
}

// IsWeekdayN reports whether the given date is the nth occurrence of the
// day in the month.
//
// The value of n affects the direction of counting:
//   n > 0: counting begins at the first day of the month.
//   n == 0: the result is always false.
//   n < 0: counting begins at the end of the month.
func IsWeekdayN(date time.Time, day time.Weekday, n int) bool {
	cday := date.Weekday()
	if cday != day || n == 0 {
		return false
	}

	if n > 0 {
		return (date.Day()-1)/7 == (n - 1)
	}

	n = -n
	last := time.Date(date.Year(), date.Month()+1,
		1, 12, 0, 0, 0, date.Location())
	lastCount := 0
	for {
		last = last.AddDate(0, 0, -1)
		if last.Weekday() == day {
			lastCount++
		}
		if lastCount == n || last.Month() != date.Month() {
			break
		}
	}

	return lastCount == n && last.Month() == date.Month() &&
		last.Day() == date.Day()
}

// WorkdayFn reports whether the given date is a workday.
// This is useful for situations where work days change throughout the year.
//
// If your workdays are fixed (Mon-Fri for example) then a WorkdayFn
// is not necessary and you can use cal.SetWorkday() instead.
type WorkdayFn func(date time.Time) bool

// Calendar represents a yearly calendar with a list of holidays.
type Calendar struct {
	holidays    [13][]Holiday // 0 for offset based holidays, 1-12 for month based
	workday     [7]bool       // flags to indicate a day of the week is a workday
	WorkdayFunc WorkdayFn     // optional function to override workday flags
	Observed    ObservedRule
}

// NewCalendar creates a new Calendar with no holidays defined
// and work days of Monday through Friday.
func NewCalendar() *Calendar {
	c := &Calendar{}

	for i := range c.holidays {
		c.holidays[i] = make([]Holiday, 0, 2)
	}
	c.workday[time.Monday] = true
	c.workday[time.Tuesday] = true
	c.workday[time.Wednesday] = true
	c.workday[time.Thursday] = true
	c.workday[time.Friday] = true
	return c
}

// newCalendarFromCountryCode returns the Holiday calendar associated to the
// country designated by the given uppercase ISO 3166-1 alpha-2 country code.
func newCalendarFromCountryCode(code string) (*Calendar, error) {
	c := NewCalendar()

	// As a special case, allow creating an empty calendar using the
	// unused/reserved code ZZ.
	if code == "ZZ" {
		return c, nil
	}

	if obs, ok := map[string]ObservedRule{
		// TODO: Complete this list when the knowledge is there.
		"CA": ObservedNearest,
		"GB": ObservedNearest,
		"US": ObservedNearest,
	}[code]; ok {
		c.Observed = obs
	}

	// TODO: Consistency on func names.
	fn, ok := map[string]func(*Calendar){
		"AT": addAustrianHolidays,
		"AU": addAustralianHolidays,
		"BE": addBelgiumHolidays,
		"CA": addCanadianHolidays,
		"DE": addGermanHolidays,
		"DK": addDanishHolidays,
		"ES": addSpainHolidays,
		"FR": addFranceHolidays,
		"GB": addBritishHolidays,
		"IT": addItalianHolidays,
		"NL": addDutchHolidays,
		"NO": addNorwegianHolidays,
		"NZ": addNewZealandHoliday,
		"PL": addPolandHolidays,
		"SE": addSwedishHolidays,
		"US": addUSHolidays,
	}[code]
	if !ok {
		return nil, fmt.Errorf("no calendar exists for country %s", code)
	}

	fn(c)

	return c, nil
}

// NewLocalCalendar creates a Calendar from an ISO-3166-1 Alpha 2 or ISO-3166-2.
// To obtain a valid but empty calendar you can use the "ZZ" country code.
func NewLocalCalendar(code string) (*Calendar, error) {
	parts := strings.SplitN(code, "-", 2)
	if len(parts) < 1 {
		return nil, fmt.Errorf("%s is not a valid ISO-3166-1 Alpha 2 or ISO-3166-2 code", code)
	}

	c, err := newCalendarFromCountryCode(parts[0])
	if err != nil {
		return nil, err
	}
	if len(parts) < 2 {
		return c, nil
	}

	fn, ok := map[string]func(*Calendar, string) error{
		"CA": addCanadaProvinceHolidays,
		"DE": addGermanyStateHolidays,
		"FR": addFranceDepartmentHolidays,
	}[parts[0]]
	if !ok {
		return c, nil
	}

	if err := fn(c, code); err != nil {
		return nil, err
	}

	return c, nil
}

// GetHolidays returns the list of all holidays defined in the calendar between start and end
func (c Calendar) GetHolidays(start, end time.Time) []struct {
	Date  time.Time
	Label string
} {
	var ret []struct {
		Date  time.Time
		Label string
	}

	// Knowingly Quadratic
	for v := truncate(start); v.Before(end); v = v.AddDate(0, 0, 1) {
		date, label, ok := c.GetHoliday(v)
		if ok {
			ret = append(ret, struct {
				Date  time.Time
				Label string
			}{date, label})
		}
	}

	return ret
}

// truncate truncates a Date to 00h00m00s
func truncate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// AddHoliday adds a Holiday to the calendar's list.
func (c *Calendar) AddHoliday(h ...Holiday) {
	for _, hd := range h {
		c.holidays[hd.Month] = append(c.holidays[hd.Month], hd)
	}
}

// SetWorkday changes the given day's status as a standard working day
func (c *Calendar) SetWorkday(day time.Weekday, workday bool) {
	c.workday[day] = workday
}

// IsHoliday reports whether a given date is a Holiday. It does not account
// for the observation of holidays on alternate days.
func (c *Calendar) IsHoliday(date time.Time) bool {
	_, _, ok := c.GetHoliday(date)
	return ok
}

func (c *Calendar) GetHoliday(date time.Time) (time.Time, string, bool) {
	idx := date.Month()
	label := func(h Holiday, date time.Time) string {
		if h.Label == "" {
			return date.Format("2006-01-02")
		}
		return h.Label
	}

	for i := range c.holidays[idx] {
		if c.holidays[idx][i].Matches(date) {
			return truncate(date), label(c.holidays[idx][i], date), true
		}
	}

	for i := range c.holidays[0] {
		if c.holidays[0][i].Matches(date) {
			return truncate(date), label(c.holidays[0][i], date), true
		}
	}

	return time.Time{}, "", false
}

// IsWorkday reports whether a given date is a work day (business day).
func (c *Calendar) IsWorkday(date time.Time) bool {
	day := date.Weekday()

	var workday bool
	if c.WorkdayFunc == nil {
		workday = c.workday[day]
	} else {
		workday = c.WorkdayFunc(date)
	}

	if !workday || c.IsHoliday(date) {
		return false
	}

	if c.Observed == ObservedExact {
		return true
	}

	if (c.Observed == ObservedMonday || c.Observed == ObservedMondayTuesday) &&
		day == time.Monday {
		sun := date.AddDate(0, 0, -1)
		sat := date.AddDate(0, 0, -2)
		return !c.IsHoliday(sat) && !c.IsHoliday(sun)
	}

	if c.Observed == ObservedMondayTuesday && day == time.Tuesday {
		mon := date.AddDate(0, 0, -1)
		sun := date.AddDate(0, 0, -2)
		sat := date.AddDate(0, 0, -3)
		return !(c.IsHoliday(sat) && c.IsHoliday(sun)) &&
			!(c.IsHoliday(sat) && c.IsHoliday(mon)) &&
			!(c.IsHoliday(sun) && c.IsHoliday(mon))
	}

	if c.Observed == ObservedNearest {
		if day == time.Friday {
			sat := date.AddDate(0, 0, 1)
			return !c.IsHoliday(sat)
		}

		if day == time.Monday {
			sun := date.AddDate(0, 0, -1)
			return !c.IsHoliday(sun)
		}
	}

	return true
}

// countWorkdays reports the number of workdays from the given date to the end
// of the month.
func (c *Calendar) countWorkdays(dt time.Time, month time.Month) int {
	n := 0
	for ; month == dt.Month(); dt = dt.AddDate(0, 0, 1) {
		if c.IsWorkday(dt) {
			n++
		}
	}
	return n
}

// Workdays reports the total number of workdays for the given year and month.
func (c *Calendar) Workdays(year int, month time.Month) int {
	return c.countWorkdays(time.Date(year, month, 1, 12, 0, 0, 0, time.UTC), month)
}

// WorkdaysRemain reports the total number of remaining workdays in the month
// for the given date.
func (c *Calendar) WorkdaysRemain(date time.Time) int {
	return c.countWorkdays(date.AddDate(0, 0, 1), date.Month())
}

// WorkdayN reports the day of the month that corresponds to the nth workday
// for the given year and month.
//
// The value of n affects the direction of counting:
//   n > 0: counting begins at the first day of the month.
//   n == 0: the result is always 0.
//   n < 0: counting begins at the end of the month.
func (c *Calendar) WorkdayN(year int, month time.Month, n int) int {
	var date time.Time
	var add int
	if n == 0 {
		return 0
	}

	if n > 0 {
		date = time.Date(year, month, 1, 12, 0, 0, 0, time.UTC)
		add = 1
	} else {
		date = time.Date(year, month+1, 1, 12, 0, 0, 0, time.UTC).AddDate(0, 0, -1)
		add = -1
		n = -n
	}

	ndays := 0
	for ; month == date.Month(); date = date.AddDate(0, 0, add) {
		if c.IsWorkday(date) {
			ndays++
			if ndays == n {
				return date.Day()
			}
		}
	}
	return 0
}

// WorkdaysFrom reports the date of a workday that is offset days
// away from start.
//
// If n > 0, then the date returned is start + offset workdays.
// If n == 0, then the date is returned unchanged.
// If n < 0, then the date returned is start - offset workdays.
func (c *Calendar) WorkdaysFrom(start time.Time, offset int) time.Time {
	date := start
	var add int

	if offset == 0 {
		return start
	}

	if offset > 0 {
		add = 1
	} else {
		add = -1
		offset = -offset
	}

	for ndays := 0; ndays < offset; {
		date = date.AddDate(0, 0, add)
		if c.IsWorkday(date) {
			ndays++
		}
	}

	return date
}

// CountHolidayHoursWithOffset returns the number of working hours in a range starting from the consumed start date
// to the end date set by the offset
func (c *Calendar) CountHolidayHoursWithOffset(start time.Time, offsetHour int) int {
	days := int(math.Ceil(float64(offsetHour) / float64(24)))

	holidayHours := 0
	day := 0
	for day <= days {
		date := start.AddDate(0, 0, day)
		if !c.IsWorkday(date) {
			holidayHours += 24
			days++
		}
		day++
	}

	return holidayHours
}

//CountWorkdays return amount of workdays between start and end dates
func (c *Calendar) CountWorkdays(start, end time.Time) int64 {
	factor := 1
	if end.Before(start) {
		factor = -1
		start, end = end, start
	}
	result := 0
	var i time.Time
	for i = start; i.Before(end); i = i.AddDate(0, 0, 1) {
		if c.IsWorkday(i) {
			result++
		}
	}
	if i.Equal(end) && c.IsWorkday(end) {
		result++
	}
	return int64(factor * result)
}

// AddSkipNonWorkdays returns start time plus d working duration
func (c *Calendar) AddSkipNonWorkdays(start time.Time, d time.Duration) time.Time {
	const day = 24 * time.Hour
	s := start
	for {
		for !c.IsWorkday(s) {
			s = s.Add(day)
		}

		if d >= day { // nolint:gocritic
			s = s.Add(day)
			d -= day
		} else if d > 0 {
			s = s.Add(d)
			d = 0
		} else {
			break
		}
	}
	return s
}

// SubSkipNonWorkdays  returns start time minus d working duration
func (c *Calendar) SubSkipNonWorkdays(start time.Time, d time.Duration) time.Time {
	const day = 24 * time.Hour * -1
	s := start
	for {
		for !c.IsWorkday(s) {
			s = s.Add(day)
		}

		if d >= day*-1 { // nolint:gocritic
			s = s.Add(day)
			d += day
		} else if d > 0 {
			s = s.Add(-d)
			d = 0
		} else {
			break
		}
	}
	return s
}
