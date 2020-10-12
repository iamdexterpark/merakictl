package device

import (
	"github.com/ddexterpark/merakictl/api/general/devices/monitor"
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



// lldpCdp - List LLDP and CDP information for a device
var lldpCdp = &cobra.Command{
	Use:   "lldpCdp",
	Short: "List LLDP and CDP information for a device.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		serial := cmd.Flag("device").Value.String()
		lldpCdp, traceback := monitor.GetlldpCdp(serial)
		shell.Display(lldpCdp, traceback, "lldpCdp", cmd.Flags())
	},
}

// uplink - Get the uplink loss percentage and latency in milliseconds for a wired network device.
var uplink = &cobra.Command{
	Use:   "uplink",
	Short: "Get the uplink loss percentage and latency in milliseconds for a wired network device.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		serial := cmd.Flag("device").Value.String()
		uplink, traceback := monitor.GetUplinkLoss(serial, args[0], args[1],
			args[2], args[3], args[4], args[5])
		shell.Display(uplink, traceback, "uplinkLoss", cmd.Flags())
	},
}