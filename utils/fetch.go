package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alexsasharegan/dotenv"
)

func getSessionCookie() string {
	err := dotenv.Load()
	Check(err)

	return os.Getenv("SESSION_COOKIE")
}

// Fetch input data for problem
func Fetch(year int, day int) []byte {
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

	// return strings.TrimSpace(string(body))
	return body
}
