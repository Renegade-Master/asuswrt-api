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

package asuswrt

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

const (
	Android string = "asusrouter-Android-DUTUtil-1.0.0.3.58-163"
	Desktop        = "Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"

	SessionName = "main"
)

// WrtClient is an interface for defining methods that should exist in a WrtClient
type WrtClient interface {
	// Login Connects to the Asus WRT Router and assigns a Token to the internal Client.
	Login(username string, password string) error

	// Logout of the session.
	Logout() error

	// GetConnectedClients returns a list of the clients connected to the Network.
	GetConnectedClients() error

	// GetWanTraffic returns the current upload and download speed.
	GetWanTraffic() error
}

// AsusWrt is an implementation of the WrtClient interface
type AsusWrt struct {
	Client *http.Client
	store  *sessions.CookieStore
	Url    string
}

func (awrt *AsusWrt) Login(username string, password string) error {
	awrt.store = sessions.NewCookieStore()
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
