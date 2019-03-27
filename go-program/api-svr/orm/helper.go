package orm

// Query conditions
type Q map[string]interface{}

// To base type: map
func (q Q) V() map[string]interface{} {
	return q
}
