package go_type

import (
	"fmt"

	api_models "github.com/angenalZZZ/Go/go-program/api-models"
)

// 类型检查
func TestTypeCheck() {
	var p api_models.IPoint = &api_models.Point{X: 1, Y: 2}
	var p2 = make([]api_models.Point, 2)
	fmt.Println("-------------------------\n类型检查：")

	// type assertion (*指针类型)
	if p0, ok := p.(*api_models.Point); ok {
		fmt.Printf("  类型断言: %p  %p\n", &p, &p0)
	}

	//var v1 bool
	//var v2 byte   // uint8  [true 或 false]
	//var v3 rune   // uint8, uint16, uint32 [unicode 编码: 1, 2, 4 个字节]
	//var v4 int    // 32位
	//var v40 uint  // 64位
	//var v5 int8   // -128~127
	//var v50 uint8 // 0 ~ 255
	//var v6 int16
	//var v60 uint16
	//var v7 int32
	//var v70 uint32
	//var v8 int64
	//var v80 uint64
	//var v9 uintptr // 存储指针的 uint32 或 uint64
	//var f1 float32 // 小数位数精确到  7 位
	//var f2 float64 // 小数位数精确到 15 位
	//var c1 complex64
	//var c2 complex128
	//var s1 string  // readonly byte slice
	//var s2 stringS

	fmt.Println(`  格式化p：%v %+v %T %#v make(Slice::Point)`)
	fmt.Printf("  格式化p：%v %+v %T %#v [%d]Point\n", p, p, p, p, cap(p2))
	fmt.Printf("  格式化i：%c %8.1f %8.2f %8x\n", 65, 12.5, 12.509, 54349)

	// 类型检查 指针
	PtrTypeCheck()

	// 二维数组
	TwoImensionalArrays(4, 2)
}
