package pkg

// MustNotError panic if err.
func MustNotError(err error) {
	if err != nil {
		panic(err)
	}
}
