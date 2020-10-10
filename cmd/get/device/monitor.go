package device

import (
	"github.com/ddexterpark/merakictl/api/devices/monitor"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// clients
var clients = &cobra.Command{
	Use:   "clients",
	Short: "Return A Devices Clients.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		serial := cmd.Flag("device").Value.String()
		deviceClients, traceback := monitor.GetDeviceClients(serial)
		shell.Display(deviceClients, traceback, "deviceClients", cmd.Flags())
	},
}
