/*
 *    Copyright (c) 2022 Renegade-Master [renegade.master.dev@protonmail.com]
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	Android string = "asusrouter-Android-DUTUtil-1.0.0.3.58-163"
	Desktop        = "Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"
)

// AsusWrtClient Interface for defining methods that should exist in an AsusWrt
type AsusWrtClient interface {
	// Login Connects to the Asus WRT Router and assigns a Token to the internal Client
	Login(username string, password string) error
	Logout() error
	GetConnectedClients() error
	GetWanTraffic() error
}

type AsusWrt struct {
	Client *http.Client
	url    string
}

func (awrt *AsusWrt) Login(username string, password string) error {
	auth := fmt.Sprintf("%s:%s", username, password)
	loginToken := base64.StdEncoding.EncodeToString([]byte(auth))

	form := url.Values{}
	form.Add("login_authorization", loginToken)
	encForm := strings.NewReader(form.Encode())

	if response, err := sendRequest(awrt, http.MethodPost, "login.cgi", encForm, Android); err != nil {
		log.Errorf("Request Failed: %s", err)
		return err
	} else {
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
				log.Debugf("\tCookie: [%s]\n\t\tValue: [%s]\n", cookie, cookie.Value)
			}
		}

		return nil
	}
}

func (awrt *AsusWrt) Logout() error {
	form := url.Values{}
	encForm := strings.NewReader(form.Encode())

	if response, err := sendRequest(awrt, http.MethodPost, "Logout.asp", encForm, Android); err != nil {
		log.Errorf("Request Failed: %s", err)
		return err
	} else {
		log.Infof("Request Success!\n")
		log.Debugf("Response: [%+v]", response)

		return nil
	}
}

func (awrt *AsusWrt) GetConnectedClients() error {
	payload := "get_clientlist(appobj);wl_sta_list_2g(appobj);wl_sta_list_5g(appobj);wl_sta_list_5g_2(appobj);nvram_get(custom_clientlist)"

	jsonBody := []byte(`{"hook": ` + payload + `}`)
	encForm := bytes.NewReader(jsonBody)

	if response, err := sendRequest(awrt, http.MethodPost, "appGet.cgi", encForm, Android); err != nil {
		log.Errorf("Request Failed: %s", err)
		return err
	} else {
		log.Infof("Request Success!\n")
		log.Debugf("Response: [%+v]", response)

		return nil
	}
}

func (awrt *AsusWrt) GetWanTraffic() error {
	//TODO implement me
	panic("implement me")
}

func sendRequest[T *bytes.Reader | *strings.Reader](client *AsusWrt, method string, path string, payload T, useragent string) (*http.Response, error) {
	reqPath := fmt.Sprintf("%s/%s", client.url, path)

	if request, err := http.NewRequest(method, reqPath, io.Reader(payload)); err != nil {
		log.Errorf("Failed to create Request: %s", err)
		return nil, err
	} else {
		request.Header.Set("User-Agent", useragent)
		request.Header.Set("Accept-Language", "en-US,en;q=0.5")
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		log.Debugf("Request: [%+v]", request)

		if response, err := client.Client.Do(request); err != nil {
			log.Errorf("Request Failed: %s", err)
			return nil, err
		} else {
			switch response.StatusCode {
			case 400:
				log.Infof("Invalid request. Rejected by Server")
				return response, err
			default:
				return response, nil
			}
		}
	}
}
