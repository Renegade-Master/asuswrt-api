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

type options struct {
	Verbose  bool   `required:"false" short:"v" long:"verbose"  description:"Show verbose debug information"`
	Url      string `required:"true"  short:"a" long:"address"  description:"URL of the Asus WRT Router"`
	Username string `required:"false"  short:"u" long:"username" description:"Username of the account on the Asus WRT Router" env:"ASUS_USER"`
	Password string `required:"false"  short:"p" long:"password" description:"Password of the account on the Asus WRT Router" env:"ASUS_PASS"`
}

func newHttpClient() *http.Client {
	// Disable Certificate Checking
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transportConfig := &http.Transport{
		DisableCompression: true,
		TLSClientConfig:    tlsConfig,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transportConfig,
	}

	return client
}

func main() {
	var opts options
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

	// Set the Log Level
	if opts.Verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Infof("Flags accepted!\n")

	var asusWrt awrt.WrtClient = &awrt.AsusWrt{
		Client: newHttpClient(),
		Url:    opts.Url,
	}

	log.Debugf("AsusWRT Client: %+v\n", asusWrt)

	log.Infof("Logging in...")
	asusWrt.Login(opts.Username, opts.Password)

	log.Infof("Get Connected Clients...")
	asusWrt.GetConnectedClients()

	log.Infof("Logging out...")
	asusWrt.Logout()
}
