package ms

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/switch/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var GetPackets = &cobra.Command{
	Use:   "Packets",
	Short: "Return the packet counters for all the ports of a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")

		metadata := monitor.GetPacketCounters(serial, t0, timespan)
			shell.Display(metadata, "Packets", cmd.Flags())
	},
}

var GetPortsStatuses = &cobra.Command{
	Use:   "PortsStatuses",
	Short: "Return the status for all the ports of a ms.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")

		metadata := monitor.GetStatuses(serial, t0, timespan)
		shell.Display(metadata, "PortsStatuses", cmd.Flags())
	},
}