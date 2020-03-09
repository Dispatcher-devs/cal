package cal

import "time"

// Holidays in Spain
var (
	esAñoNuevo               = newYear.setLabel("Año Nuevo")
	esReyes                  = newHoliday(time.January, 6).setLabel("Reyes")
	esFiestaDelTrabajo       = newHoliday(time.May, 1).setLabel("Fiesta del trabajo")
	esAsuncionDeLaVirgen     = newHoliday(time.August, 15).setLabel("Asuncion de la virgen")
	esFiestaNacionalDeEspaña = newHoliday(time.October, 12).setLabel("Fiesta Nacional de España")
	esTodosLosSantos         = newHoliday(time.November, 1).setLabel("Todos los Santos")
	esConstitucion           = newHoliday(time.December, 6).setLabel("Constitucion") // nolint:mispell
	esInmaculadaConcepcion   = newHoliday(time.December, 8).setLabel("Inmaculada concepcion")
	esNavidad                = christmas.setLabel("Navidad")
)

// addSpainHolidays adds all Spain holidays to the Calendar
func addSpainHolidays(c *Calendar) {
	c.addHoliday(
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
