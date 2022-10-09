package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// AsusWrtClient Interface for defining methods that should exist in an AsusWrt
type AsusWrtClient interface {
	// Login Connects to the Asus WRT Router and assigns a Token to the internal Client
	Login() error
	Logout() error
	GetWanTraffic() error
}

type AsusWrt struct {
	Client    *http.Client
	url       string
	username  string
	password  string
	useragent string
}

func (awrt *AsusWrt) Login() error {
	auth := fmt.Sprintf("%s:%s", awrt.username, awrt.password)
	loginToken := base64.StdEncoding.EncodeToString([]byte(auth))

	jsonBody := []byte(`{"login_authorization": "` + loginToken + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	r, err := awrt.sendRequest(http.MethodPost, "login.cgi", bodyReader)

	if err != nil {
		log.Errorf("Request Failed: %s", err)
		return err
	}

	log.Infof("Request Success!\n")

	if r != nil {
		log.Debugf("Response: %+v\n", *r)
	}

	if r.Header != nil {
		log.Debugf("Header: %+v\n", r.Header)
	}

	if r.Body != nil {
		log.Debugf("Body: %+v\n", r.Body)
	}

	if r.Cookies() != nil {
		log.Debugf("Cookies:\n")
		for _, cookie := range r.Cookies() {
			log.Debugf("Cookie: [%s]\nValue: [%s]\n", cookie, cookie.Value)
		}
	}

	return nil
}

func (awrt *AsusWrt) sendRequest(method string, path string, payload *bytes.Reader) (*http.Response, error) {
	var req *http.Request
	reqPath := fmt.Sprintf("%s/%s", awrt.url, path)

	if r, err := http.NewRequest(method, reqPath, payload); err != nil {
		log.Errorf("Failed to create Request: %s", err)
	} else {
		req = r
	}

	req.Header.Set("User-Agent", awrt.useragent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	log.Debugf("Request: [%+v]", req)

	if resp, err := awrt.Client.Do(req); err != nil {
		log.Errorf("Request Failed: %s", err)
		return nil, err
	} else {
		return resp, nil
	}
}

//func (awrt *AsusWrt) GetWanTraffic() error {
//
//}
