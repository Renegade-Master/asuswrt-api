package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("Hello, World!")
}

func TestCreateAsusWrtClient(t *testing.T) {
	var asusWrt = AsusWrt{
		ipAddr:   "127.0.0.1",
		port:     9999,
		username: "test_user",
		password: "test_pass",
	}

	fmt.Printf("%+v\n", asusWrt)

	if err := asusWrt.login(); err != nil {
		t.Errorf("Error connecting to the AsusWRT Device: %s\n", err)
	}
}
