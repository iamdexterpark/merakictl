package wireless

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/wireless/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

// ssid
var ssid = &cobra.Command{
	Use:   "ssid",
	Short: "Update The MR SSIDs In A Network.",
	Long:  `merakictl  update mr ssid {SSID#} -n {NetworkId} --input {FILENAME}.yaml`,
	Run: func(cmd *cobra.Command, args []string) {

		// Flags
		networkId := cmd.Flag("network").Value.String()

		// SSID Number
		ssidNumber := args[0]

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.UpdateSSID(networkId, ssidNumber, format)
		shell.Display(metadata, "ssid", cmd.Flags())

	},
}
