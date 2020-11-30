package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/alexsasharegan/dotenv"
)

func getSessionCookie() string {
	err := dotenv.Load()
	Check(err)

	return os.Getenv("SESSION_COOKIE")
}

// Fetch input data for problem
func Fetch(year int, day int) string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	Check(err)

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", getSessionCookie()
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	Check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	return strings.TrimSpace(string(body))
}

// FetchLines fetches input data and return lines
func FetchLines(year int, day int) []string {
	body := Fetch(year, day)

	return strings.Split(body, "\n")
}

// FetchNumbers fetches input data and return lines as numbers
func FetchNumbers(year int, day int) []int {
	lines := FetchLines(year, day)
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i] = ToInt(line)
	}

	return numbers
}

// ToInt converts string to integer
func ToInt(text string) int {
	x, err := strconv.Atoi(text)
	Check(err)

	return x
}
