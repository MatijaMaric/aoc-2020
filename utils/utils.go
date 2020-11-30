package utils

// Check panics if error is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
