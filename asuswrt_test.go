package main

import (
	"fmt"
	"testing"
)

type MockAsusWrtClient struct {
	DoLogin func() error
}

func (m *MockAsusWrtClient) Login() error {
	fmt.Println("Hello, World!")
	return nil
}

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

	if err := asusWrt.Login(); err != nil {
		t.Errorf("Error connecting to the AsusWRT Device: %s\n", err)
	}
}
