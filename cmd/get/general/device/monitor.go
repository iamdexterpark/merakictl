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
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")
		metadata := monitor.GetDeviceClients(serial, t0, timespan)
		shell.Display(metadata, "deviceClients", cmd.Flags())
	},
}

// lldpCdp - List LLDP and CDP information for a device
var lldpCdp = &cobra.Command{
	Use:   "lldpCdp",
	Short: "List LLDP and CDP information for a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := monitor.GetlldpCdp(serial)
		shell.Display(metadata, "lldpCdp", cmd.Flags())
	},
}

// uplink - Get the uplink loss percentage and latency in milliseconds for a wired network device.
var uplink = &cobra.Command{
	Use:   "uplink",
	Short: "Get the uplink loss percentage and latency in milliseconds for a wired network device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		resolution, _ := cmd.Flags().GetString("resolution")
		uplink, _ := cmd.Flags().GetString("uplink")
		ip, _ := cmd.Flags().GetString("ip")

		metadata := monitor.GetUplinkLoss(serial,t0, t1, timespan, resolution, uplink, ip)
		shell.Display(metadata, "uplinkLoss", cmd.Flags())
	},
}
