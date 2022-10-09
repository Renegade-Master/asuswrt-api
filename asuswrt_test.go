package main

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fmt.Println("Hello, World!")
}

func TestCreateAsusWrtClient(t *testing.T) {
	expected := "Test"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	var asusWrt = AsusWrt{
		Client:   MockHttpClient(),
		url:      svr.URL,
		username: "test_user",
		password: "test_pass",
	}

	fmt.Printf("%+v\n", asusWrt)

	if err := asusWrt.Login(); err != nil {
		t.Errorf("Error connecting to the AsusWRT Device: %s\n", err)
	}
}

func MockHttpClient() *http.Client {
	log.Infof("Running AsusWrt Client init\n")

	// Disable Certificate Checking
	tlsConfig := tls.Config{InsecureSkipVerify: true}

	client := http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tlsConfig},
	}

	return &client
}
