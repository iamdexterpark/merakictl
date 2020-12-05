package get

import (
	"github.com/ddexterpark/dashboard-api-golang/api"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

// apikey -> api.Token()
var environmentalvariables = &cobra.Command{
	Use:   "env",
	Short: "Prints out environmental variables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		obfuscated := string(api.Token()[len(api.Token())-6:])
		pretty.Println("API KEY is : *****************"+obfuscated)
		pretty.Println("API Version is :", api.APIversion())
		pretty.Println("API BaseUrl is :", api.BaseUrl())


	},
}

var test = &cobra.Command{
	Use:   "test",
	Short: "test stuff.",
	Long:  pretty.Sprint("test"),
	Run: func(cmd *cobra.Command, args []string) {
		pretty.Println(shell.ResolveFlags(cmd.Flags()))
	},
}