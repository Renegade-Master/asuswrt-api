package main

import (
	"crypto/tls"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type Options struct {
	Verbose  []bool `required:"false" short:"v" long:"verbose" description:"Show verbose debug information"`
	IpAddr   string `required:"true" short:"a" long:"address" description:"IP Address of the Asus WRT Router"`
	Port     uint   `required:"true" long:"port" description:"Port of the Asus WRT Router"`
	Username string `required:"true" short:"u" long:"username" description:"Username of the account on the Asus WRT Router"`
	Password string `required:"true" short:"p" long:"password" description:"Password of the account on the Asus WRT Router"`
}

func NewHttpClient() *http.Client {
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

func main() {
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

	log.Debugf("IP Address:\t%s\n", opts.IpAddr)
	log.Debugf("Port:\t%d\n", opts.Port)
	log.Debugf("Username:\t%s\n", opts.Username)
	log.Debugf("Password:\t%s\n", opts.Password)

	var asusWrt = AsusWrt{
		ipAddr:   opts.IpAddr,
		port:     opts.Port,
		username: opts.Username,
		password: opts.Password,
		Client:   NewHttpClient(),
	}

	log.Debug("AsusWRT Client: %+v\n", asusWrt)

	asusWrt.Login()
}
