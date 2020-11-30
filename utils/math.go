package utils

// MaxUint max unsigned integer
const MaxUint = ^uint(0)

// MinUint min unsigned integer
const MinUint = 0

// MaxInt max signed integer
const MaxInt = int(MaxUint >> 1)

// MinInt min signed integer
const MinInt = -MaxInt - 1

// Abs return absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Max returns greater of a and b
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns lesser of a and b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GCD returns greatest common divisor of a and b
func GCD(a, b int) int {
	a, b = Abs(a), Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns least common multiple of numbers
func LCM(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	ans := nums[0] * nums[1] / GCD(nums[0], nums[1])

	for _, num := range nums[2:] {
		ans = LCM(ans, num)
	}

	return ans
}

// Permutations returns all permutations of numbers to n
func Permutations(n int) [][]int {
	if n == 1 {
		return [][]int{{0}}
	}
	perms := Permutations(n - 1)

	interleaved := make([][]int, 0, len(perms)*n)
	for _, perm := range perms {
		for i := 0; i <= len(perm); i++ {
			withN := make([]int, len(perm)+1)
			copy(withN[:i], perm[:i])
			withN[i] = n - 1
			copy(withN[i+1:], perm[i:])
			interleaved = append(interleaved, withN)
		}
	}

	return interleaved
}
