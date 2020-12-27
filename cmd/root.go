/*
Copyright © 2020 Dexter Park dDexterPark@icloud.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/ddexterpark/merakictl/cmd/create"
	"github.com/ddexterpark/merakictl/cmd/get"
	"github.com/ddexterpark/merakictl/cmd/misc"
	"github.com/ddexterpark/merakictl/cmd/remove"
	"github.com/ddexterpark/merakictl/cmd/update"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "merakictl",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the format command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var cfgFile string

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
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

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(misc.CompletionCmd)
	rootCmd.AddCommand(misc.Version)
	//rootCmd.AddCommand(utilities.Upgrade)
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(update.UpdateCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(remove.DeleteCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "input", "", "input file (default is $HOME/.merakictl.yaml)")
	rootCmd.PersistentFlags().BoolP("diff", "d", false, "diff config file with dashboard API config")
	rootCmd.PersistentFlags().BoolP("export", "e", false, "Export config to Yaml")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "Export config to JSON")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Display HTTP Request/Response for traceback")

	// flags for resolving names to dashboard Ids
	rootCmd.PersistentFlags().StringP("organization", "o", "", "The target organization")
	rootCmd.PersistentFlags().StringP("network", "n", "", "The target Network")
	rootCmd.PersistentFlags().StringP("hostname", "s", "", "The target devices hostname")

	// Flags for URL Parameters need to be moved. Hard to read autocompletion help menu.
		rootCmd.PersistentFlags().StringP("perPage", "", "10", "")
		rootCmd.PersistentFlags().StringP("startingAfter", "", "", "")
		rootCmd.PersistentFlags().StringP("endingBefore", "", "", "")
		rootCmd.PersistentFlags().StringP("configurationUpdatedAfter", "", "", "")
		rootCmd.PersistentFlags().StringP("t0", "", "", "")
		rootCmd.PersistentFlags().StringP("t1", "", "", "")
		rootCmd.PersistentFlags().StringP("timespan", "", "", "")
		rootCmd.PersistentFlags().StringP("resolution", "", "", "")
		rootCmd.PersistentFlags().StringP("autoResolution", "", "", "")
		rootCmd.PersistentFlags().StringP("deviceSerial", "", "", "")
		rootCmd.PersistentFlags().StringP("uplink", "", "", "")
		rootCmd.PersistentFlags().StringP("ip", "", "", "")
		rootCmd.PersistentFlags().StringP("search", "", "", "")
		rootCmd.PersistentFlags().StringP("configTemplateId", "", "", "")
		rootCmd.PersistentFlags().StringP("tags", "", "", "")
		rootCmd.PersistentFlags().StringP("usedState", "", "", "")
		rootCmd.PersistentFlags().StringP("tagsFilterType", "", "", "")
		rootCmd.PersistentFlags().StringP("username", "", "", "")
		rootCmd.PersistentFlags().StringP("email", "", "", "")
		rootCmd.PersistentFlags().StringP("mac", "", "", "")
		rootCmd.PersistentFlags().StringP("serial", "", "", "")
		rootCmd.PersistentFlags().StringP("imei", "", "", "")
		rootCmd.PersistentFlags().StringP("bluetoothMac", "", "", "")
		rootCmd.PersistentFlags().StringP("includeConnectivityHistory", "", "", "")
		rootCmd.PersistentFlags().StringP("connectivityHistoryTimespan", "", "", "")
		rootCmd.PersistentFlags().StringP("includedEventTypes", "", "", "")
		rootCmd.PersistentFlags().StringP("excludedEventTypes", "", "", "")
		rootCmd.PersistentFlags().StringP("sensorSerial", "", "", "")
		rootCmd.PersistentFlags().StringP("gatewaySerial", "", "", "")
		rootCmd.PersistentFlags().StringP("loginIdentifier", "", "", "")
		rootCmd.PersistentFlags().StringP("deviceType", "", "", "")
		rootCmd.PersistentFlags().StringP("sortOrder", "", "", "")
		rootCmd.PersistentFlags().StringP("networkIds", "", "", "")
		rootCmd.PersistentFlags().StringP("objectType", "", "", "")
		rootCmd.PersistentFlags().StringP("clientId", "", "", "")
		rootCmd.PersistentFlags().StringP("apTag", "", "", "")
		rootCmd.PersistentFlags().StringP("band", "", "", "")
		rootCmd.PersistentFlags().StringP("ssid", "", "", "")
		rootCmd.PersistentFlags().StringP("ssidNumber", "", "", "")
		rootCmd.PersistentFlags().StringP("vlan", "", "", "")
		rootCmd.PersistentFlags().StringP("fields", "", "", "")
		rootCmd.PersistentFlags().StringP("wifiMacs", "", "", "")
		rootCmd.PersistentFlags().StringP("serials", "", "", "")
		rootCmd.PersistentFlags().StringP("ids", "", "", "")
		rootCmd.PersistentFlags().StringP("scope", "", "", "")
		rootCmd.PersistentFlags().StringP("withDetails", "", "", "")
		rootCmd.PersistentFlags().StringP("licenceId", "", "", "")
}
