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
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fmt.Println("Hello, World!")
}

func TestCreateClient(t *testing.T) {
	var asusWrt = AsusWrt{
		Client: MockHttpClient(),
		url:    "https://127.0.0.1:8443",
	}

	fmt.Printf("%+v\n", asusWrt)
}

func TestLogin(t *testing.T) {
	expected := "Set-Cookie: asus_token=dGVzdGVzdCBzdHJpbmcKCBzdHJpbmcK; HttpOnly;" +
		"Connection: close"

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	var asusWrt = AsusWrt{
		Client: MockHttpClient(),
		url:    svr.URL,
	}

	if err := asusWrt.Login("test_user", "test_pass"); err != nil {
		t.Errorf("Error connecting to the AsusWRT Device: %s\n", err)
	}
}

func TestLoout(t *testing.T) {
	expected := "{" +
		"\"error_status\":\"1\"" +
		"}"

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	var asusWrt = AsusWrt{
		Client: MockHttpClient(),
		url:    svr.URL,
	}

	if err := asusWrt.Logout(); err != nil {
		t.Errorf("Error connecting to the AsusWRT Device: %s\n", err)
	}
}

func MockHttpClient() *http.Client {
	// Disable Certificate Checking
	tlsConfig := tls.Config{InsecureSkipVerify: true}

	client := http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tlsConfig},
	}

	return &client
}
