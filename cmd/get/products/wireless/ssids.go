package wireless

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/wireless/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

// ssids
var ssids = &cobra.Command{
	Use:   "ssids",
	Short: "List The MR SSIDs In A Network.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		networkId := cmd.Flag("network").Value.String()
		metadata := configure.GetSSIDS(networkId)
		shell.Display(metadata, "ssids", cmd.Flags())
	},
}

// ssid
var ssid = &cobra.Command{
	Use:   "ssid",
	Short: "List A Single MR SSID In A Network.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		networkId := cmd.Flag("network").Value.String()

		ssidNumber := args[0]
		metadata := configure.GetSSID(networkId, ssidNumber)
		shell.Display(metadata, "ssid", cmd.Flags())
	},
}
