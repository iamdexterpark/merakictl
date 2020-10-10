package wireless

import (
	"github.com/ddexterpark/merakictl/api/wireless"
	"github.com/ddexterpark/merakictl/shell"
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

		session, traceback := wireless.UpdateSSID(networkId,ssidNumber,format)
		shell.Display(session, traceback, "ssid", cmd.Flags())

	},
}
