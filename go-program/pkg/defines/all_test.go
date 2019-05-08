package defines

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
	"unicode/utf8"
	"unsafe"

	api_models "github.com/angenalZZZ/Go/go-program/api-models"
)

/**
命令行参数
*/
var temperature = FlagCelsius("t", 20.0, "the temperature")

// 类型检查
func TestTypeCheck(t *testing.T) {

	var p api_models.IPoint = &api_models.Point{X: 1, Y: 2}
	var p2 = make([]api_models.Point, 2)
	fmt.Println("-------------------------\n类型检查：")

	// 命令行参数
	fmt.Printf("  命令行参数/摄氏温度: %s\n", temperature)

	// type assertion (*指针类型)
	if p0, ok := p.(*api_models.Point); ok {
		fmt.Printf("  类型断言: %p  %p\n", &p, p0)
	}
	// interface{} 接受任意类型的变量, 不同动态类型的变量不可比较, 只能与nil比较
	var w io.Writer // zeroValue=nil, 接受实现接口: Write(p []byte) 类型的变量, 下面的动态值决定了接收者类型(*T)的不同
	fmt.Printf("  接口w io.Writer(type)：%T, (value)：%[1]v \n", w)
	w = os.Stdout
	fmt.Printf("  接口w os.Stdout(type)：%T, (value)：%[1]v \n", w)
	w = new(bytes.Buffer)
	fmt.Println("  接口w new(bytes.Buffer)(type)：", reflect.TypeOf(w), ", (value)：", w) // %T: reflect.TypeOf(w)

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

	fmt.Println(`  格式化p：v +v T #v make(Slice::Point)`)
	fmt.Printf("  格式化p：%v %+v %T %#v [%d]Point\n", p, p, p, p, cap(p2))
	fmt.Printf("  格式化i：%c %8.1f %8.2f %8x\n", 65, 12.5, 12.509, 54349)

	// 斐波那契数列
	new(Fibonacci).FibonacciToDo(20, 2*time.Second, func(s []int) {
		fmt.Printf("  斐波那契数列: %v", s)
	})
}

// 检查字符串是文字字面值时才是 UTF8 文本
func TestStringType(t *testing.T) {
	var s1 = SS{"1", "2"}
	var s2 = make([]string, 2)
	var s3 = [...]string{"1", "2", "3", "4", "5"}
	var s4 = s3[1:4:5] // 切片: [low:high:max]

	fmt.Println(s1, s2, s3,
		s4,      // "2", "3", "4"
		len(s4), // 4 - 1 len: high-low
		cap(s4), // 5 - 1 cap: max-low
		//cap(s1) == cap(s2),
		utf8.ValidString("ABC") == true,
		utf8.ValidString("A\\xfeC") == true,
		utf8.ValidString("A\xfeC") == false,
		utf8.RuneCountInString("é") == 2, // 两个 rune 的组合
		len("é") == 3, len("é") == len("\u0301"))
}

// 类型检查 指针
func TestPointerType(t *testing.T) {

	// array int
	a := [4]int{0, 1, 2, 3}
	a0 := unsafe.Pointer(&a[0])
	a3 := (*int)(unsafe.Pointer(uintptr(a0) + 3*unsafe.Sizeof(a[0]))) // 指针 偏移 Offset
	*(a3) = 4
	fmt.Println("  指针：array int: a =", a) // [0 1 2 4]

	// struct Person
	type Person struct {
		name   string
		age    int
		gender byte
	}
	who := Person{"John Mono", 30, 0}
	p := unsafe.Pointer(&who)                                                   // 指针 类似 C 语言的 void* 与其他语言的指针,相互转换的桥梁
	name := (*string)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(who.name)))   // 指针 偏移 member: name
	age := (*int)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(who.age)))        // 指针 偏移 member: age
	gender := (*byte)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(who.gender))) // 指针 偏移 member: gender
	*name = "Alice"
	*age = 28
	*gender = 1
	fmt.Printf("  指针：struct Person: a = %v\n", who) // {Alice 28 1}

}

// 二维数组
func TestTwoImensionalArrays(t *testing.T) {

	cols, rows := 2, 4

	fmt.Printf("  二维数组：TwoImensionalArrays(%d,%d)\n", cols, rows)

	raw := make([]int, cols*rows)

	for i := range raw { // range slice array's index
		raw[i] = i + 1
	}

	fmt.Printf("  raw: %+v , %p\n", raw, &raw[0])

	tbl := make([][]int, rows)

	for i := range tbl {
		tbl[i] = raw[i*cols : i*cols+cols] // range slice array
	}

	fmt.Printf("  tbl: %+v , %p\n", tbl, &tbl[0][0])
}
