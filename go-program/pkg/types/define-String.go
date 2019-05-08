package types

import (
	"sort"
	"unicode/utf8"
)

// bytes 转换为 []string
func BytesToStrings(buf interface{}) (r sort.StringSlice) {
	if bu, OK := buf.([]interface{}); OK {
		i := 0
		r = make([]string, len(bu))
		for _, b := range bu {
			if v, OK := b.([]byte); OK {
				r[i] = string(v)
				i++
			}
		}
	}
	return
}

// string 类型的值是常量，不可直接更改；遍历请用: range []byte(s)
func SetChar(s string, i int, v rune) string {
	if len(s) <= i || utf8.ValidRune(v) == false {
		return s
	}
	b := []rune(s)
	b[i] = v // rune: uint8, uint16, uint32 [unicode 编码: 1, 2, 4 个字节]
	return string(b)
}

// 字符串切片/扩展方法：字符串添加并过滤
func (s SS) AppendWithFilter(f func(string) bool) (r sort.StringSlice) {
	r = make([]string, 0)
	for _, t := range s {
		if f(t) {
			r = append(r, t)
		}
	}
	return
}

// 字符串切片/扩展方法：映射处理
func (s SS) Map(f func(string) string) (r sort.StringSlice) {
	r = make([]string, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return
}
