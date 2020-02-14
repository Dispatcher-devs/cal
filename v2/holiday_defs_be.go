package cal

import "time"

// Holidays in Belgium
var (
	beNieuwjaar                = newYear
	bePaasmaandag              = easterMonday
	beDagVanDeArbeid           = ecbLabourDay
	beOnzeLieveHeerHemelvaart  = NewHolidayFunc(calculateOnzeLieveHeerHemelvaart)
	bePinkstermaandag          = NewHolidayFunc(calculatePinkstermaandag)
	beNationaleFeestdag        = NewHoliday(time.July, 21)
	beOnzeLieveVrouwHemelvaart = NewHoliday(time.August, 15)
	beAllerheiligen            = NewHoliday(time.November, 1)
	beWapenstilstand           = NewHoliday(time.November, 11)
	beKerstmis                 = christmas
)

// addBelgiumHolidays adds all Belgium holidays to the Calendar
func addBelgiumHolidays(c *Calendar) {
	c.AddHoliday(
		beNieuwjaar,
		bePaasmaandag,
		beDagVanDeArbeid,
		beOnzeLieveHeerHemelvaart,
		bePinkstermaandag,
		beNationaleFeestdag,
		beOnzeLieveVrouwHemelvaart,
		beAllerheiligen,
		beWapenstilstand,
		beKerstmis,
	)
}

func calculateOnzeLieveHeerHemelvaart(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 39 days after Easter Sunday
	t := easter.AddDate(0, 0, +39)
	return t.Month(), t.Day()
}

func calculatePinkstermaandag(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	// 50 days after Easter Sunday
	t := easter.AddDate(0, 0, +50)
	return t.Month(), t.Day()
}
