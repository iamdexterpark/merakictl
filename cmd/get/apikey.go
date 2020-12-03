package get

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

// apikey -> api.Token()
var apikey = &cobra.Command{
	Use:   "apikey",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("API KEY is : %s \n", api.Token())

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