/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/ddexterpark/merakictl/cmd/remove"
	"github.com/ddexterpark/merakictl/cmd/update"
	"github.com/spf13/cobra"
	"log"
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



func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(version)
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(update.UpdateCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(remove.DeleteCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "input", "", "input file (default is $HOME/.merakictl.yaml)")
	rootCmd.PersistentFlags().StringP("org", "o", "", "The target organization")
	rootCmd.PersistentFlags().StringP("network", "n", "", "The target Network")
	rootCmd.PersistentFlags().BoolP("export", "e", false, "Export config to Yaml")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "Export config to JSON")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Display HTTP Request/Response for traceback")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}




