package go_type

// To base type: map
func (q Q) V() map[string]interface{} {
	return q
}

// map 是否存在key?
func (q Q) Contains(key string) bool {
	if _, ok := q[key]; ok {
		return true
	}
	return false
}

// map 删除key
func (q Q) Del(key string) {
	delete(q, key)
}

// clone 复制map
func (q Q) Clone() Q {
	var r = Q{}
	for k, v := range q {
		r[k] = v
	}
	return r
}

// slice 拆分为切片
func (q Q) Slice() (s []interface{}) {
	s = make([]interface{}, len(q)*2)
	for k, v := range q {
		s = append(s, k)
		s = append(s, v)
	}
	return
}
