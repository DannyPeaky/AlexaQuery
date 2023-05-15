package storage

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"time"
)

type SaveCookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	Expires  time.Time `json:"expires"`
	Secure   bool      `json:"secure"`
	HttpOnly bool      `json:"http_only"`
}

type LoadCookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	Expires  string `json:"expires"`
	Secure   bool   `json:"secure"`
	HttpOnly bool   `json:"http_only"`
}

func parseTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// LoadCookiesFromJSON loads cookies from a JSON file into the cookie jar
func LoadCookiesFromJSON(jar http.CookieJar, url *url.URL, filePath string) error {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the JSON data into a slice of LoadCookie structs
	var cookies []LoadCookie
	err = json.NewDecoder(file).Decode(&cookies)
	if err != nil {
		return err
	}

	// Iterate over the cookies and add them to the cookie jar
	for _, c := range cookies {
		// Parse the Expires time string
		expires, err := parseTime(c.Expires)
		if err != nil {
			return err
		}

		// Create a new http.Cookie
		cookie := &http.Cookie{
			Name:     c.Name,
			Value:    c.Value,
			Path:     c.Path,
			Domain:   c.Domain,
			Expires:  expires,
			Secure:   c.Secure,
			HttpOnly: c.HttpOnly,
		}

		// Add the cookie to the cookie jar
		jar.SetCookies(url, []*http.Cookie{cookie})
	}

	return nil
}

func SaveCookiesToJSON(jar http.CookieJar, url *url.URL, filePath string) error {
	// Get all cookies for the specified URL
	cookies := jar.Cookies(url)

	// Create a slice to hold the Cookie structs
	var cookieList []SaveCookie

	// Iterate over all cookies
	for _, cookie := range cookies {
		// Create a new Cookie struct
		c := SaveCookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Path:     "/",
			Domain:   ".amazon.co.uk",
			Expires:  cookie.Expires,
			Secure:   true,
			HttpOnly: cookie.HttpOnly,
		}

		// Append the Cookie struct to the slice
		cookieList = append(cookieList, c)
	}

	// Marshal the slice to JSON
	data, err := json.Marshal(cookieList)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
