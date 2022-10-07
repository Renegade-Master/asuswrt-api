package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type AsusWrt struct {
	ipAddr   string
	port     uint
	username string
	password string
}

func (awrt AsusWrt) login() {
	baseAddr := fmt.Sprintf("%s:%d", awrt.ipAddr, awrt.port)
	auth := fmt.Sprintf("%s:%s", awrt.username, awrt.password)
	loginToken := base64.StdEncoding.EncodeToString([]byte(auth))
	//headers := map[string]string{
	//	"user-agent": "asusrouter-Android-DUTUtil-1.0.0.245",
	//}

	// Disable Certificate Checking
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	payload := url.Values{}
	payload.Add("login_authorization", loginToken)

	r, err := http.PostForm(fmt.Sprintf("https://%s/login.cgi", baseAddr), payload)

	if err != nil {
		log.Errorf("Request Failed: %s", err)
		return
	}

	log.Infof("Request Success!\n%+v", r)

	//try:
	//	r, err := requests.post(url='http://%s/login.cgi'), data=payload, headers=headers).json()
	//except:
	//	return False
	//	if "asus_token" in r:
	//	token = r['asus_token']
	//	self.headers = {
	//		'user-agent': "asusrouter-Android-DUTUtil-1.0.0.245",
	//			'cookie': 'asus_token={}'.format(token)
	//	}
	//	return True
	//	else:
	//	return False
}
