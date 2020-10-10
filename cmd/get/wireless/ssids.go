package wireless

import (
	"github.com/ddexterpark/merakictl/api/wireless"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// ssids
var ssids = &cobra.Command{
	Use:   "ssids",
	Short: "List The MR SSIDs In A Network.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		networkId := cmd.Flag("network").Value.String()
		ssids, traceback := wireless.GetSSIDS(networkId)
		shell.Display(ssids, traceback, "ssids", cmd.Flags())
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
		ssid, traceback := wireless.GetSSID(networkId,ssidNumber)
		shell.Display(ssid, traceback, "ssid", cmd.Flags())
	},
}