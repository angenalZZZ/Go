package go_type

///////////类型声明和定义///////////

// 字符串切片: 类型 string slice
type stringS []string

///////////函数声明和定义///////////

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
