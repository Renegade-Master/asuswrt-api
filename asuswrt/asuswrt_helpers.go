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
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

func sendRequest[
	T *bytes.Reader | *strings.Reader](
	client *AsusWrt, method string, path string, payload T, useragent string) (*http.Response, error) {

	reqPath := fmt.Sprintf("%s/%s", client.Url, path)

	if request, err := http.NewRequest(method, reqPath, io.Reader(payload)); err != nil {
		log.Errorf("Failed to create Request: %s", err)
		return nil, err
	} else {
		request.Header.Set("User-Agent", useragent)
		request.Header.Set("Accept-Language", "en-US,en;q=0.5")
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		var session *sessions.Session
		if session, err = getSession(client, request, SessionName); err != nil {
			log.Errorf("Request Failed: %s", err)
			return nil, err
		}

		log.Debugf("Request: [%+v]", request)
		log.Debugf("Request Session: [%+v]", session)

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

func getSession(client *AsusWrt, request *http.Request, sessionName string) (*sessions.Session, error) {
	if session, err := client.store.Get(request, sessionName); err != nil {
		log.Errorf("Failed to retrieve Session information: %s", err)
		return nil, err
	} else {
		return session, nil
	}
}
