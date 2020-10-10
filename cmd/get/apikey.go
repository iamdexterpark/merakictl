package get

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"github.com/spf13/cobra"
)

// apikey -> api.Token()
var apikey = &cobra.Command{
	Use:   "apikey",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("API KEY is : %s \n",api.Token())

	},
}
