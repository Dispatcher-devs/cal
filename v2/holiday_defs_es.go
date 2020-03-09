package cal

import "time"

// Holidays in Spain
var (
	esAñoNuevo               = newYear.SetLabel("Año Nuevo")
	esReyes                  = newHoliday(time.January, 6).SetLabel("Reyes")
	esFiestaDelTrabajo       = newHoliday(time.May, 1).SetLabel("Fiesta del trabajo")
	esAsuncionDeLaVirgen     = newHoliday(time.August, 15).SetLabel("Asuncion de la virgen")
	esFiestaNacionalDeEspaña = newHoliday(time.October, 12).SetLabel("Fiesta Nacional de España")
	esTodosLosSantos         = newHoliday(time.November, 1).SetLabel("Todos los Santos")
	esConstitucion           = newHoliday(time.December, 6).SetLabel("Constitucion") // nolint:mispell
	esInmaculadaConcepcion   = newHoliday(time.December, 8).SetLabel("Inmaculada concepcion")
	esNavidad                = christmas.SetLabel("Navidad")
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
