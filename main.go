package main

import (
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"os"
)

type Options struct {
	Verbose  []bool `required:"false" short:"v" long:"verbose" description:"Show verbose debug information"`
	IpAddr   string `required:"true" short:"a" long:"address" description:"IP Address of the Asus WRT Router"`
	Port     uint   `required:"true" long:"port" description:"Port of the Asus WRT Router"`
	Username string `required:"true" short:"u" long:"username" description:"Username of the account on the Asus WRT Router"`
	Password string `required:"true" short:"p" long:"password" description:"Password of the account on the Asus WRT Router"`
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

	log.Infof("IP Address:\t%s\n", opts.IpAddr)
	log.Infof("Port:\t%d\n", opts.Port)
	log.Infof("Username:\t%s\n", opts.Username)
	log.Infof("Password:\t%s\n", opts.Password)

	var asusWrt = AsusWrt{
		ipAddr:   opts.IpAddr,
		port:     opts.Port,
		username: opts.Username,
		password: opts.Password,
	}

	log.Infof("AsusWRT Client: %+v\n", asusWrt)

	asusWrt.login()
}
