package cal

import "time"

// Holidays in Belgium
var (
	beNieuwjaar                = newYear.setLabel("Nieuwjaar")
	bePaasmaandag              = easterMonday.setLabel("Paasmaandag")
	beDagVanDeArbeid           = ecbLabourDay.setLabel("Dag van de Arbeid")
	beOnzeLieveHeerHemelvaart  = newHolidayFunc(calculateOnzeLieveHeerHemelvaart).setLabel("Onze Lieve Heer Hemelvaart")
	bePinkstermaandag          = newHolidayFunc(calculatePinkstermaandag).setLabel("Pinkstermaandag")
	beNationaleFeestdag        = newHoliday(time.July, 21).setLabel("Nationale Feestdag")
	beOnzeLieveVrouwHemelvaart = newHoliday(time.August, 15).setLabel("Onze Lieve Vrouw Hemelvaart")
	beAllerheiligen            = newHoliday(time.November, 1).setLabel("Allerheiligen")
	beWapenstilstand           = newHoliday(time.November, 11).setLabel("Wepenstilstand")
	beKerstmis                 = christmas.setLabel("Kerstmis")
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
