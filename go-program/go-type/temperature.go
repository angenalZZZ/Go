package go_type

import "fmt"

// 摄氏 → 华氏 温度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

// 华氏 → 摄氏 温度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// 摄氏 温度 输出字符串 接口 → fmt.Stringer
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
