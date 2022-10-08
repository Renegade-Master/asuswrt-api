package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// AsusWrtClient Interface for defining methods that should exist in an AsusWrt
type AsusWrtClient interface {
	// Login Connects to the Asus WRT Router and assigns a Token to the internal Client
	Login() error
	Logout() error
}

type AsusWrt struct {
	Client   *http.Client
	ipAddr   string
	port     uint
	username string
	password string
}

func (awrt *AsusWrt) Login() error {
	baseAddr := fmt.Sprintf("%s:%d", awrt.ipAddr, awrt.port)
	auth := fmt.Sprintf("%s:%s", awrt.username, awrt.password)
	loginToken := base64.StdEncoding.EncodeToString([]byte(auth))
	reqUrl := fmt.Sprintf("https://%s/login.cgi", baseAddr)

	jsonBody := []byte(`{"login_authorization": ` + loginToken + `}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, reqUrl, bodyReader)

	r, err := awrt.Client.Do(req)

	if err != nil {
		log.Errorf("Request Failed: %s", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("Error closing the HTTP Request")
		}
	}(r.Body)

	log.Infof("Request Success!\n")

	if r != nil {
		log.Infof("Response: %+v\n", *r)
	}

	if r.Header != nil {
		log.Infof("Header: %+v\n", r.Header)
	}

	if r.Body != nil {
		log.Infof("Body: %+v\n", r.Body)
	}

	return nil
}
