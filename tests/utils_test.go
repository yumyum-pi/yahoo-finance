package test

import (
	"testing"

	yahooFinace "github.com/yumyum-pi/yahooFinance"
)

type s2fUnitType struct {
	s string
	b bool
	f float64
}

func TestString2Float(t *testing.T) {
	var units []s2fUnitType = []s2fUnitType{
		{"Hello", true, 0},
		{"X42", false, 42},
		{"I'm", true, 0},
		{"a", true, 0},
		{"Y-32.35", false, -32.35},
		{"string", true, 0},
		{"Z30", false, 30},
		{"1,432.90", false, 1432.90},
		{" 1,432.90", false, 1432.90},
		{" 1,432.90.2341234", false, 1432.90},
		{" 1,432.90.asdfasd", false, 1432.90},
	}
	var f float64
	var err error
	for i := 0; i < len(units); i++ {
		f, err = yahooFinace.String2Float(units[i].s)
		if err != nil {
			if !units[i].b {
				t.Errorf("unexpected error at index %d: %s", i, err)

			}
			// do something
		} else if f != units[i].f {
			t.Errorf("Was expecting %f, got %f on %s", units[i].f, f, units[i].s)
		}
	}

}
