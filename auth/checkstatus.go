package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type AuthenticationResponse struct {
	Authentication struct {
		Authenticated              bool   `json:"authenticated"`
		CanAccessPrimeMusicContent bool   `json:"canAccessPrimeMusicContent"`
		CustomerEmail              string `json:"customerEmail"`
		CustomerID                 string `json:"customerId"`
		CustomerName               string `json:"customerName"`
	} `json:"authentication"`
}

func CheckStatus(client *http.Client, browser string) (bool, error) {
	req, err := http.NewRequest("GET", "https://alexa.amazon.co.uk/api/bootstrap?version=0", nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", browser)

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var authStatus AuthenticationResponse
	if err := json.Unmarshal(body, &authStatus); err != nil {
		return false, err
	}

	return authStatus.Authentication.Authenticated, nil
}

func GetCSRFToken(jar http.CookieJar, url *url.URL) string {
	var csrfToken string
	cookies := jar.Cookies(url)
	for _, cookie := range cookies {
		if cookie.Name == "csrf" {
			csrfToken = cookie.Value
			break
		}
	}
	return csrfToken
}
