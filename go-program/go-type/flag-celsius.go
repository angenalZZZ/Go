package go_type

import "fmt"

/*
//!+flag value
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
//!-flag value
*/

//!+flagCelsius
// *flagCelsius satisfies the flag.Value interface.
type flagCelsius struct{ Celsius }

func (f *flagCelsius) Set(s string) error {
	var u string
	var v float64
	if _, e := fmt.Sscanf(s, "%f%s", &v, &u); e != nil {
		switch u {
		case "C", "°C":
			f.Celsius = Celsius(v)
			return nil
		case "F", "°F":
			f.Celsius = FToC(Fahrenheit(v))
			return nil
		}
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//!-flagCelsius
