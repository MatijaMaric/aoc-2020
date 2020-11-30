package utils

import "strconv"

// Check panics if error is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// ToInt converts string to integer
func ToInt(text string) int {
	x, err := strconv.Atoi(text)
	Check(err)

	return x
}
