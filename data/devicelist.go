package data

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"

	"github.com/dannypeaky/alexaquery/auth"
)

type Device struct {
	AccountName  string `json:"accountName"`
	DeviceType   string `json:"deviceType"`
	SerialNumber string `json:"serialNumber"`
	DeviceFamily string `json:"deviceFamily"`
}

type DeviceList struct {
	Devices []Device
}

func GetDeviceList(client *http.Client) ([]Device, error) {
	req, err := http.NewRequest("GET", "https://alexa.amazon.co.uk/api/devices-v2/device?cached=false", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Referer", "https://alexa.amazon.co.uk/spa/index.html")
	req.Header.Add("Origin", "https://alexa.amazon.co.uk")
	req.Header.Add("csrf", auth.GetCSRFToken(client.Jar, req.URL))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var deviceList DeviceList
	err = json.Unmarshal(body, &deviceList)
	if err != nil {
		return nil, err
	}

	filteredDevices := make([]Device, 0)
	for _, device := range deviceList.Devices {
		if device.DeviceFamily == "ECHO" || device.DeviceFamily == "KNIGHT" || device.DeviceFamily == "ROOK" {
			filteredDevices = append(filteredDevices, device)
		}
	}

	sort.Slice(filteredDevices, func(i, j int) bool {
		return filteredDevices[i].AccountName < filteredDevices[j].AccountName
	})

	// for _, device := range filteredDevices {
	// 	fmt.Printf("%s=%s=%s=%s\n", device.AccountName, device.DeviceType, device.SerialNumber, device.DeviceFamily)
	// }

	return filteredDevices, nil
}
