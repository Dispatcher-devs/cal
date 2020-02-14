package cal

import "time"

// Holidays in Spain
var (
	esAñoNuevo               = newYear
	esReyes                  = NewHoliday(time.January, 6)
	esFiestaDelTrabajo       = NewHoliday(time.May, 1)
	esAsuncionDeLaVirgen     = NewHoliday(time.August, 15)
	esFiestaNacionalDeEspaña = NewHoliday(time.October, 12)
	esTodosLosSantos         = NewHoliday(time.November, 1)
	esConstitucion           = NewHoliday(time.December, 6)
	esInmaculadaConcepcion   = NewHoliday(time.December, 8)
	esNavidad                = christmas
)

// addSpainHolidays adds all Spain holidays to the Calendar
func addSpainHolidays(c *Calendar) {
	c.AddHoliday(
		esAñoNuevo,
		esReyes,
		esFiestaDelTrabajo,
		esAsuncionDeLaVirgen,
		esFiestaNacionalDeEspaña,
		esTodosLosSantos,
		esConstitucion,
		esInmaculadaConcepcion,
		esNavidad,
	)
}
