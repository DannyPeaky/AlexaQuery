package alexaquery

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/dannypeaky/alexaquery/auth"
	"github.com/dannypeaky/alexaquery/data"
	"github.com/dannypeaky/alexaquery/storage"
)

type QueryClient struct {
	Client     *http.Client
	Browser    string
	CookiePath string
	Url        *url.URL
}

func (c *QueryClient) Login(token string) {
	auth.Login(c.Client, token, c.Url, c.CookiePath)
}

func (c *QueryClient) CheckStatus() (bool, error) {
	return auth.CheckStatus(c.Client, c.Browser)
}

func (c *QueryClient) GetCSRFToken() string {
	return auth.GetCSRFToken(c.Client.Jar, c.Url)
}

func (c *QueryClient) LoadCookiesFromJSON() error {
	return storage.LoadCookiesFromJSON(c.Client.Jar, c.Url, c.CookiePath)
}

func (c *QueryClient) SaveCookiesToJSON() error {
	return storage.SaveCookiesToJSON(c.Client.Jar, c.Url, c.CookiePath)
}

func (c *QueryClient) GetDeviceList() ([]data.Device, error) {
	return data.GetDeviceList(c.Client)
}

func (c *QueryClient) GetNotifications() ([]data.Notification, error) {
	return data.GetNotifications(c.Client)
}

func (c *QueryClient) GetQueue(deviceSerialNumber string, deviceType string) (data.PlayerInfo, error) {
	return data.GetQueue(c.Client, deviceSerialNumber, deviceType)
}

func NewQueryClient(cookiePath string) *QueryClient {
	jar, _ := cookiejar.New(nil)
	u, _ := url.Parse("https://amazon.co.uk")
	var client QueryClient = QueryClient{
		Client:     &http.Client{Jar: jar},
		Browser:    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:1.0) bash-script/1.0",
		CookiePath: cookiePath,
		Url:        u,
	}
	return &client
}
