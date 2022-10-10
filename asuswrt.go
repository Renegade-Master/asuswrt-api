package main

import (
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
)

const (
	Desktop string = "Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"
	Android        = "asusrouter-Android-DUTUtil-1.0.0.3.58-163"
)

// AsusWrtClient Interface for defining methods that should exist in an AsusWrt
type AsusWrtClient interface {
	// Login Connects to the Asus WRT Router and assigns a Token to the internal Client
	Login() error
	Logout() error
	GetWanTraffic() error
}

type AsusWrt struct {
	Client   *http.Client
	url      string
	username string
	password string
}

func (awrt *AsusWrt) Login() error {
	auth := fmt.Sprintf("%s:%s", awrt.username, awrt.password)
	loginToken := base64.StdEncoding.EncodeToString([]byte(auth))

	form := url.Values{}
	form.Add("login_authorization", loginToken)

	response, err := awrt.sendRequest(http.MethodPost, "login.cgi", strings.NewReader(form.Encode()), Android)

	if err != nil {
		log.Errorf("Request Failed: %s", err)
		return err
	}

	log.Infof("Request Success!\n")

	if response != nil {
		log.Debugf("Response: %+v\n", *response)
	}

	if response.Header != nil {
		log.Debugf("Header: %+v\n", response.Header)
	}

	if response.Body != nil {
		log.Debugf("Body: %+v\n", response.Body)
	}

	if response.Cookies() != nil {
		log.Debugf("Cookies:\n")
		for _, cookie := range response.Cookies() {
			log.Debugf("Cookie: [%s]\nValue: [%s]\n", cookie, cookie.Value)
		}
	}

	return nil
}

func (awrt *AsusWrt) Logout() error {
	//TODO implement me
	panic("implement me")
}

func (awrt *AsusWrt) GetWanTraffic() error {
	//TODO implement me
	panic("implement me")
}

func (awrt *AsusWrt) sendRequest(method string, path string, payload *strings.Reader, useragent string) (*http.Response, error) {
	var req *http.Request
	reqPath := fmt.Sprintf("%s/%s", awrt.url, path)

	if r, err := http.NewRequest(method, reqPath, payload); err != nil {
		log.Errorf("Failed to create Request: %s", err)
	} else {
		req = r
	}

	req.Header.Set("User-Agent", useragent)
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	log.Debugf("Request: [%+v]", req)

	if resp, err := awrt.Client.Do(req); err != nil {
		log.Errorf("Request Failed: %s", err)
		return nil, err
	} else {
		return resp, nil
	}
}
