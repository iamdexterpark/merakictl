package _switch

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/switch/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var packets = &cobra.Command{
	Use:   "packets",
	Short: "Return the packet counters for all the ports of a switch.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")

		metadata := monitor.GetPacketCounters(serial,t0,timespan)
			shell.Display(metadata, "packets", cmd.Flags())
	},
}

var portsstatuses = &cobra.Command{
	Use:   "portsstatuses",
	Short: "Return the status for all the ports of a switch.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		timespan, _ := cmd.Flags().GetString("timespan")

		metadata := monitor.GetPortsStatus(serial,t0,timespan)
		shell.Display(metadata, "portsstatuses", cmd.Flags())
	},
}