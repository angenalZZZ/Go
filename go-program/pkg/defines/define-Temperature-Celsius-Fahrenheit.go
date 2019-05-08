package defines

import (
	"flag"
	"fmt"
)

// 摄氏 温度 输出字符串 接口 → fmt.Stringer
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// 摄氏 → 华氏 温度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

// 华氏 → 摄氏 温度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

//!+FlagCelsius

// FlagCelsius defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
// Usage (°F auto convert to °C) : 华氏 → 摄氏 温度
//   var temperature = FlagCelsius("t", 20.0, "the temperature")
func FlagCelsius(name string, value Celsius, usage string) *Celsius {
	f := flagCelsius{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

//!-FlagCelsius

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
