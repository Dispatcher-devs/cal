package cal

import "time"

// Holidays in Belgium
var (
	beNieuwjaar                = newYear.SetLabel("Nieuwjaar")
	bePaasmaandag              = easterMonday.SetLabel("Paasmaandag")
	beDagVanDeArbeid           = ecbLabourDay.SetLabel("Dag van de Arbeid")
	beOnzeLieveHeerHemelvaart  = newHolidayFunc(calculateOnzeLieveHeerHemelvaart).SetLabel("Onze Lieve Heer Hemelvaart")
	bePinkstermaandag          = newHolidayFunc(calculatePinkstermaandag).SetLabel("Pinkstermaandag")
	beNationaleFeestdag        = newHoliday(time.July, 21).SetLabel("Nationale Feestdag")
	beOnzeLieveVrouwHemelvaart = newHoliday(time.August, 15).SetLabel("Onze Lieve Vrouw Hemelvaart")
	beAllerheiligen            = newHoliday(time.November, 1).SetLabel("Allerheiligen")
	beWapenstilstand           = newHoliday(time.November, 11).SetLabel("Wepenstilstand")
	beKerstmis                 = christmas.SetLabel("Kerstmis")
)

// addBelgiumHolidays adds all Belgium holidays to the Calendar
func addBelgiumHolidays(c *Calendar) {
	c.addHoliday(
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
