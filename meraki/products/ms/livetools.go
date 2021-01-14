package ms

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/switch/livetools"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

var PostPortCycle = &cobra.Command{
	Use:   "cyclePorts",
	Short: "Cycle a set of switch ports.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, serial := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		var format livetools.Cycle
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := livetools.PostCyclePorts(networkId, serial,  input)
		shell.Display(metadata, "CyclePorts", cmd.Flags())
	},
}