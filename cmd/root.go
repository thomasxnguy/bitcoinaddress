package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "btcaddress ",
	Short: "BitcoinAddress web application",
	Long:  `BitcoinAddress web application`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Cobra supports persistent flags, which, if defined here,
	// will be global for the application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ßconfig.json)")
	viper.BindPFlag("db_debug", RootCmd.PersistentFlags().Lookup("db_debug"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name "config.json".
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("config.json")
		viper.SetConfigType("json")
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
