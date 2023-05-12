package main

import (
	"alexaquery/auth"
	"alexaquery/data"
	"alexaquery/storage"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type QueryClient struct {
	client     *http.Client
	browser    string
	cookiePath string
	url        *url.URL
}

func (c *QueryClient) Login(token string) {
	auth.Login(c.client, token, c.url, c.cookiePath)
}

func (c *QueryClient) CheckStatus() (string, bool, error) {
	return auth.CheckStatus(c.client, c.browser)
}

func (c *QueryClient) GetCSRFToken() string {
	return auth.GetCSRFToken(c.client.Jar, c.url)
}

func (c *QueryClient) LoadCookiesFromJSON() error {
	return storage.LoadCookiesFromJSON(c.client.Jar, c.url, c.cookiePath)
}

func (c *QueryClient) SaveCookiesToJSON() error {
	return storage.SaveCookiesToJSON(c.client.Jar, c.url, c.cookiePath)
}

func (c *QueryClient) GetDeviceList() error {
	return data.GetDeviceList(c.client)
}

func (c *QueryClient) GetNotifications(deviceSerialNumber string, deviceType string) {
	data.GetNotifications(c.client, deviceSerialNumber, deviceType)
}

func (c *QueryClient) GetQueue(deviceSerialNumber string, deviceType string) {
	data.GetQueue(c.client, deviceSerialNumber, deviceType)
}

func NewQueryClient(cookiePath string) *QueryClient {
	jar, _ := cookiejar.New(nil)
	u, _ := url.Parse("https://amazon.co.uk")
	var client QueryClient = QueryClient{
		client:     &http.Client{Jar: jar},
		browser:    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:1.0) bash-script/1.0",
		cookiePath: cookiePath,
		url:        u,
	}
	err := client.LoadCookiesFromJSON()
	if err != nil {
		fmt.Println(err)
	}
	return &client
}
