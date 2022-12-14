package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	awrt "github.com/renegade-master/asuswrt-api/asuswrt"
	log "github.com/sirupsen/logrus"
)

type Options struct {
	Verbose  []bool `required:"false" short:"v" long:"verbose" description:"Show verbose debug information"`
	Url      string `required:"true" short:"a" long:"address" description:"URL of the Asus WRT Router"`
	Username string `required:"true" short:"u" long:"username" description:"Username of the account on the Asus WRT Router"`
	Password string `required:"true" short:"p" long:"password" description:"Password of the account on the Asus WRT Router"`
}

func NewHttpClient() *http.Client {
	// Disable Certificate Checking
	tlsConfig := tls.Config{InsecureSkipVerify: true}

	client := http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tlsConfig},
	}

	return &client
}

func main() {
	log.SetLevel(log.DebugLevel)
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				log.Infof("Displaying help\n")
				os.Exit(0)
			}

			log.Errorf("Error: %s\n", flagsErr)
			os.Exit(1)
		default:
			log.Errorf("Error: %s\n", flagsErr)
			os.Exit(1)
		}
	}
	log.Infof("Flags accepted!\n")

	var asusWrt awrt.AsusWrtClient = &awrt.AsusWrt{
		Client: NewHttpClient(),
		Url:    opts.Url,
	}

	log.Debugf("AsusWRT Client: %+v\n", asusWrt)

	log.Infof("Logging in...\n\n")
	asusWrt.Login(opts.Username, opts.Password)

	log.Infof("Get Connected Clients...\n\n")
	asusWrt.GetConnectedClients()

	log.Infof("Logging out...\n\n")
	asusWrt.Logout()
}
