package go_type

import (
	"fmt"
	"unicode/utf8"
)

///////////函数声明和定义///////////

// 检查字符串是文字字面值时才是 UTF8 文本
func CheckValidString() {

	var s1 = stringS{"1", "2"}
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

// bytes 转换为 []string
func BytesToStrings(buf interface{}) (s []string) {
	if bu, OK := buf.([]interface{}); OK {
		i := 0
		s = make([]string, len(bu))
		for _, b := range bu {
			if v, OK := b.([]byte); OK {
				s[i] = string(v)
				i++
			}
		}
	}
	return
}

// string 类型的值是常量，不可直接更改；遍历请用: range []byte(s)
func Set(s string, i int, v rune) string {
	if len(s) <= i || utf8.ValidRune(v) == false {
		return s
	}
	b := []rune(s)
	b[i] = v // rune: uint8, uint16, uint32 [unicode 编码: 1, 2, 4 个字节]
	return string(b)
}

// 字符串切片/扩展方法：字符串添加并过滤
func (s stringS) AppendWithFilter(f func(string) bool) stringS {
	r := make(stringS, 0)
	for _, t := range s {
		if f(t) {
			r = append(r, t)
		}
	}
	return r
}

// 字符串切片/扩展方法：映射处理
func (s stringS) Map(f func(string) string) stringS {
	r := make(stringS, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}
