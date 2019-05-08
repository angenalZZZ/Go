package types

import "sort"

// <类型声明> ---------------------------------------- 字符串 --------------------------------------

// 字符串切片 → []string
type SS sort.StringSlice

// <类型声明> ---------------------------------------- 字典 ----------------------------------------

// 字典 → map
type Q map[string]interface{}

// <类型声明> ---------------------------------------- 温度 ----------------------------------------

// 摄氏 温度 Temperature
type Celsius float64

// 华氏 温度 Temperature
type Fahrenheit float64

// <类型声明> ---------------------------------------- 数据 ----------------------------------------

type ByteS16 [16]byte
type ByteS24 [24]byte
type ByteS32 [32]byte
type ByteS48 [48]byte
type ByteS60 [60]byte

type FloatS16 [16]float64
type FloatS24 [24]float64
type FloatS32 [32]float64
type FloatS48 [48]float64
type FloatS60 [60]float64

type IntS24 [24]int
type IntS60 [60]int

// 斐波那契数列
type Fibonacci struct{}
