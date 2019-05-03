package go_c

/*
#include <stdio.h>
#include "./c.h" // 接口申明

// 自定义 C 函数
static void printStr(const char* str) {
	puts(str)
}

// 指针
static void pointers() {
	void *p = NULL; // 无类型的指针
	uintptr_t q = (uintptr_t)(p); // 数值化的指针 <stdint.h>
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// 入口
// 禁用编译器对复杂结构体检查的耗时: GODEBUG=cgocheck=0 go run *.go
func init() {
	str := "This is a C string.\n"

	// Go调用C 原生 stdio 函数
	C.puts(C.CString(str))

	// Go调用C 由C实现的接口 自定义 C 函数
	C.printStr(C.CString(str))

	// Go调用C 由Go实现的接口
	C.printStr1(C.CString(str))

	// Go调用C 由Go实现的接口 无需转换参数: string <=> _GoString_
	C.printStr2(str)

	// 指针
	var p unsafe.Pointer = nil // 无类型的真实指针 GC会管理该指针变量
	var q uintptr = uintptr(p) // 数值化的转换指针 只有转换作用 GC不会处理它[非指针类型]
	q = q & 0
}

// 关键字 export 导出中间文件 _cgo_export.h 供 C 调用

/* 实现 C 接口 C调用Go printStr.h */
//export printStr1
func PrintStr1(str *C.char) {
	fmt.Print(C.GoString(str))
}

/* 实现 C 接口 C调用Go # go ^1.10 */
//export printStr2
func PrintStr2(s string) {
	fmt.Print(s)
}
