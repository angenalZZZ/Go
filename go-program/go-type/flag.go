package go_type

import (
	"flag"
)

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
