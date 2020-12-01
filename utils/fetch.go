package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/alexsasharegan/dotenv"
)

func getSessionCookie() string {
	err := dotenv.Load()
	Check(err)

	return os.Getenv("SESSION_COOKIE")
}

func getCachePath(day int) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to generate package path.")
	}

	root := filepath.Dir(filepath.Dir(filename))
	return filepath.Join(root, fmt.Sprintf("day%02d", day), "input.txt")
}

func fetch(year int, day int) []byte {
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

// GetInput gets input for problem and caches it
func GetInput(year int, day int) (input []byte) {
	cachePath := getCachePath(day)
	if _, stat := os.Stat(cachePath); os.IsNotExist(stat) {
		input = fetch(year, day)
		err := ioutil.WriteFile(cachePath, input, 0644)
		Check(err)
	} else {
		input, _ = ioutil.ReadFile(cachePath)
	}
	return
}
