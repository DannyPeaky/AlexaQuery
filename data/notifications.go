package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dannypeaky/alexaquery/auth"
)

func GetNotifications(client *http.Client, deviceSerialNumber string, deviceType string) ([]Notification, error) {
	url := fmt.Sprintf("https://alexa.amazon.co.uk/api/notifications?deviceSerialNumber=%s&deviceType=%s", deviceSerialNumber, deviceType)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:1.0) bash-script/1.0")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Referer", "https://alexa.amazon.co.uk/spa/index.html")
	req.Header.Set("Origin", "https://alexa.amazon.co.uk")
	req.Header.Set("csrf", auth.GetCSRFToken(client.Jar, req.URL))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var notifications Notifications
	err = json.Unmarshal(body, &notifications)
	if err != nil {
		return nil, err
	}

	// Print the response body
	return notifications.Notifications, nil
}
