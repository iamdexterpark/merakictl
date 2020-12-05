package device

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/devices/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

// clients
var clients = &cobra.Command{
	Use:   "clients",
	Short: "Return A Devices Clients.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		serial := cmd.Flag("device").Value.String()
		metadata := monitor.GetDeviceClients(serial)
		shell.Display(metadata, "deviceClients", cmd.Flags())
	},
}

// lldpCdp - List LLDP and CDP information for a device
var lldpCdp = &cobra.Command{
	Use:   "lldpCdp",
	Short: "List LLDP and CDP information for a device.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		serial := cmd.Flag("device").Value.String()
		metadata := monitor.GetlldpCdp(serial)
		shell.Display(metadata, "lldpCdp", cmd.Flags())
	},
}

// uplink - Get the uplink loss percentage and latency in milliseconds for a wired network device.
var uplink = &cobra.Command{
	Use:   "uplink",
	Short: "Get the uplink loss percentage and latency in milliseconds for a wired network device.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		serial := cmd.Flag("device").Value.String()
		metadata := monitor.GetUplinkLoss(serial, args[0], args[1],
			args[2], args[3], args[4], args[5])
		shell.Display(metadata, "uplinkLoss", cmd.Flags())
	},
}
