package go_type

// type define Map
type Map map[string]interface{}

// map 是否存在key?
func (m *Map) Contains(key string) bool {
	if _, ok := (*m)[key]; ok {
		return true
	}
	return false
}
