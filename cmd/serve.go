package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thomasxnguy/bitcoinaddress/api"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Long:  `Starts a http server and serves the api`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := api.NewServer()
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// flags and configuration settings.
	viper.SetDefault("port", "localhost:3000")
	viper.SetDefault("log_level", "debug")
}
