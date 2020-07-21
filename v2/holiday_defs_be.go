package cal

import "time"

// Holidays in Belgium
var (
	beNieuwjaar                = newYear.SetLabel("Nieuwjaar")
	bePaasmaandag              = easterMonday.SetLabel("Paasmaandag")
	beDagVanDeArbeid           = ecbLabourDay.SetLabel("Dag van de Arbeid")
	beOnzeLieveHeerHemelvaart  = NewHolidayFunc(calculateOnzeLieveHeerHemelvaart).SetLabel("Onze Lieve Heer Hemelvaart")
	bePinkstermaandag          = NewHolidayFunc(calculatePinkstermaandag).SetLabel("Pinkstermaandag")
	beNationaleFeestdag        = NewHoliday(time.July, 21).SetLabel("Nationale Feestdag")
	beOnzeLieveVrouwHemelvaart = NewHoliday(time.August, 15).SetLabel("Onze Lieve Vrouw Hemelvaart")
	beAllerheiligen            = NewHoliday(time.November, 1).SetLabel("Allerheiligen")
	beWapenstilstand           = NewHoliday(time.November, 11).SetLabel("Wepenstilstand")
	beKerstmis                 = christmas.SetLabel("Kerstmis")
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
