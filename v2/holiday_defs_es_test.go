package cal

import (
	"testing"
	"time"
)

func TestSpainHolidays(t *testing.T) {
	c, err := newCalendarFromCountryCode("ES")
	if err != nil {
		t.Fatal(err)
	}

	assertHolidays(t, c, []testStruct{
		{time.Date(2019, 1, 1, 12, 0, 0, 0, time.UTC), true, "ESAñoNuevo"},
		{time.Date(2019, 1, 6, 12, 0, 0, 0, time.UTC), true, "ESReyes"},
		{time.Date(2019, 5, 1, 12, 0, 0, 0, time.UTC), true, "ESFiestaDelTrabajo"},
		{time.Date(2019, 8, 15, 12, 0, 0, 0, time.UTC), true, "ESAsuncionDeLaVirgen"},
		{time.Date(2019, 10, 12, 12, 0, 0, 0, time.UTC), true, "ESFiestaNacionalDeEspaña"},
		{time.Date(2019, 11, 1, 12, 0, 0, 0, time.UTC), true, "ESTodosLosSantos"},
		{time.Date(2019, 12, 6, 12, 0, 0, 0, time.UTC), true, "ESConstitucion"},
		{time.Date(2019, 12, 8, 12, 0, 0, 0, time.UTC), true, "ESInmaculadaConcepcion"},
		{time.Date(2019, 12, 25, 12, 0, 0, 0, time.UTC), true, "ESNavidad"},
	})
}
