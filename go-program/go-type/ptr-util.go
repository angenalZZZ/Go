package go_type

import (
	"fmt"
	"unsafe" // unsafe包,可以像C一样去操作内存
)

// 类型检查 指针
func PtrTypeCheck() {

	// array int
	a := [4]int{0, 1, 2, 3}
	a0 := unsafe.Pointer(&a[0])
	a3 := unsafe.Pointer(uintptr(a0) + 3*unsafe.Sizeof(a[0]))
	*(*int)(a3) = 4
	fmt.Printf("  指针：array int: a = %v  %x..%dbytes..%x\n", a, uintptr(a0), (uintptr(a3)-uintptr(a0))/8, uintptr(a3)) // [0 1 2 4] byte=uint8

	// struct Person
	type Person struct {
		name   string
		age    int
		gender byte
	}
	who := Person{"John Mono", 30, 0}
	p := unsafe.Pointer(&who)                                                   // 指针 类似 C 语言的 void* 与其他语言的指针,相互转换的桥梁
	name := (*string)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(who.name)))   // 指针 struct member: name
	age := (*int)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(who.age)))        // 指针 struct member: age
	gender := (*byte)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(who.gender))) // 指针 struct member: gender
	*name = "Alice"
	*age = 28
	*gender = 1
	fmt.Printf("  指针：struct Person: a = %v\n", who) // {Alice 28 1}

}
