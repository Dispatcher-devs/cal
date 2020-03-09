package cal

import "time"

// Holidays in Spain
var (
	esAñoNuevo               = newYear.SetLabel("Año Nuevo")
	esReyes                  = NewHoliday(time.January, 6).SetLabel("Reyes")
	esFiestaDelTrabajo       = NewHoliday(time.May, 1).SetLabel("Fiesta del trabajo")
	esAsuncionDeLaVirgen     = NewHoliday(time.August, 15).SetLabel("Asuncion de la virgen")
	esFiestaNacionalDeEspaña = NewHoliday(time.October, 12).SetLabel("Fiesta Nacional de España")
	esTodosLosSantos         = NewHoliday(time.November, 1).SetLabel("Todos los Santos")
	esConstitucion           = NewHoliday(time.December, 6).SetLabel("Constitucion") // nolint:mispell
	esInmaculadaConcepcion   = NewHoliday(time.December, 8).SetLabel("Inmaculada concepcion")
	esNavidad                = christmas.SetLabel("Navidad")
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
