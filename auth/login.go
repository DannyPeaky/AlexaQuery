package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dannypeaky/alexaquery/storage"
)

type CookieData struct {
	Response struct {
		Tokens struct {
			Cookies map[string][]struct {
				Expires  string `json:"Expires"`
				HttpOnly bool   `json:"HttpOnly"`
				Name     string `json:"Name"`
				Path     string `json:"Path"`
				Secure   bool   `json:"Secure"`
				Value    string `json:"Value"`
			} `json:"cookies"`
		} `json:"tokens"`
	} `json:"response"`
}

func Login(client *http.Client, refreshToken string, clientUrl *url.URL, cookieFilePath string) {

	// This is the data we will send in our post request
	data := url.Values{}
	data.Set("app_name", "Amazon Alexa")
	data.Set("requested_token_type", "auth_cookies")
	data.Set("domain", "www.amazon.co.uk")
	data.Set("source_token_type", "refresh_token")
	data.Set("source_token", refreshToken)

	req, _ := http.NewRequest("POST", "https://api.amazon.co.uk/ap/exchangetoken/cookies", strings.NewReader(data.Encode()))
	req.Header.Add("x-amzn-identity-auth-domain", "api.amazon.co.uk")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	body, _ := io.ReadAll(resp.Body)

	// Unmarshal JSON response to CookieData
	var cookieData CookieData
	json.Unmarshal(body, &cookieData)

	var cookiesToSave []*http.Cookie
	for _, cookies := range cookieData.Response.Tokens.Cookies {
		for _, cookie := range cookies {
			// Parse the expiration date
			expiration, _ := time.Parse(time.RFC1123, cookie.Expires)
			// Create a new http.Cookie and add it to the jar
			httpCookie := &http.Cookie{
				Name:     cookie.Name,
				Value:    cookie.Value,
				Path:     cookie.Path,
				Domain:   ".amazon.co.uk",
				Expires:  expiration,
				Secure:   cookie.Secure,
				HttpOnly: cookie.HttpOnly,
			}
			cookiesToSave = append(cookiesToSave, httpCookie)
		}
	}
	client.Jar.SetCookies(clientUrl, cookiesToSave)

	// Attempt to get CSRF cookie
	csrfURLs := []string{
		"https://alexa.amazon.co.uk/api/language",
		"https://alexa.amazon.co.uk/templates/oobe/d-device-pick.handlebars",
		"https://alexa.amazon.co.uk/api/devices-v2/device?cached=false",
	}

	// Headers to be used in the request
	headers := map[string]string{
		"DNT":        "1",
		"Connection": "keep-alive",
		"Referer":    "https://alexa.amazon.co.uk/spa/index.html",
		"Origin":     "https://alexa.amazon.co.uk",
	}

	csrfCookieExists := false

	for _, csrfURL := range csrfURLs {
		req, _ := http.NewRequest("GET", csrfURL, nil)
		// Add the headers to the request
		for key, value := range headers {
			req.Header.Add(key, value)
		}
		client.Do(req)

		// Check if csrf cookie exists
		for _, cookie := range client.Jar.Cookies(req.URL) {
			if cookie.Name == "csrf" {
				csrfCookieExists = true
				break
			}
		}

		if csrfCookieExists {
			break
		}
	}

	if !csrfCookieExists {
		fmt.Println("ERROR: no CSRF cookie received")
		os.Exit(1)
	}

	err = storage.SaveCookiesToJSON(client.Jar, clientUrl, cookieFilePath)
	if err != nil {
		fmt.Println("ERROR: failed to save cookies to file:", err)
		os.Exit(1)
	}
}
