package go_type

import "unicode/utf8"

///////////类型声明和定义///////////

// 字符串切片: 类型 string slice
type stringS []string

///////////函数声明和定义///////////

// 检查字符串是文字字面值时才是 UTF8 文本
func CheckValidString() {
	println(utf8.ValidString("ABC") == true,
		utf8.ValidString("A\\xfeC") == true,
		utf8.ValidString("A\xfeC") == false,
		utf8.RuneCountInString("é") == 2, // 两个 rune 的组合
		len("é") == 3, len("é") == len("\u0301"))
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
