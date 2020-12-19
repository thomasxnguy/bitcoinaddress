package integrationtests

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thomasxnguy/bitcoinaddress/api"
	"gopkg.in/h2non/baloo.v3"
)

var ch1 = make(chan struct{})

var mockServer = baloo.New("http://localhost:3000")

// init is the global initialization function for the itest package.
func init() {
	// flags and configuration settings.
	viper.SetDefault("port", "localhost:3000")
	viper.SetDefault("log_level", "debug")

	go func() {
		server, err := api.NewServer()
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
		//ch1 <- struct{}{}
	}()
}
