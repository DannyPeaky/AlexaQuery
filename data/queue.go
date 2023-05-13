package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dannypeaky/alexaquery/auth"
)

func GetQueue(client *http.Client, deviceSerialNumber string, deviceType string) (PlayerInfo, error) {
	url := fmt.Sprintf("https://alexa.amazon.co.uk/api/np/player?deviceSerialNumber=%s&deviceType=%s", deviceSerialNumber, deviceType)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PlayerInfo{}, err
	}

	// Set request headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:1.0) bash-script/1.0")
	req.Header.Set("DNT", "1")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Referer", "https://alexa.amazon.co.uk/spa/index.html")
	req.Header.Set("Origin", "https://alexa.amazon.co.uk")
	req.Header.Set("csrf", auth.GetCSRFToken(client.Jar, req.URL))

	resp, err := client.Do(req)
	if err != nil {
		return PlayerInfo{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PlayerInfo{}, err
	}

	var queue Queue
	err = json.Unmarshal(body, &queue)
	if err != nil {
		return PlayerInfo{}, err
	}

	return queue.PlayerInfo, nil
}
