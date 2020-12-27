package utilities

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"os"
)

var CfgFile string

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name ".merakictl" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(currentdir())
		viper.SetConfigName(".merakictl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
	}
}

// currentdir - get the current working directory
func currentdir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}
