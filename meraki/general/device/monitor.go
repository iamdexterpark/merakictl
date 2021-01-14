package device

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/devices/monitor"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

var GetClients = &cobra.Command{
	Use:   "clients",
	Short: "Return A Devices Clients",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")
		metadata := monitor.GetClients(serial, t0, timespan)
		shell.Display(metadata, "Clients", cmd.Flags())
	},
}

var GetLLdpCdp = &cobra.Command{
	Use:   "lldpCdp",
	Short: "List LLDP and CDP information for a device",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := monitor.GetLLdpCdp(serial)
		shell.Display(metadata, "LLdpCdp", cmd.Flags())
	},
}


var GetLossAndLatencyHistory = &cobra.Command{
	Use:   "lossAndLatencyHistory",
	Short: "Get the uplink loss percentage and latency in milliseconds for a wired network device",
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

		metadata := monitor.GetLossAndLatencyHistory(serial, t0, t1, timespan, resolution, uplink, ip)
		shell.Display(metadata, "Uplink", cmd.Flags())
	},
}
