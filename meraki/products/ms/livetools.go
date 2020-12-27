package ms

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/switch/livetools"
	shell "github.com/ddexterpark/merakictl/utilities"
	"github.com/spf13/cobra"
)

var PostPortCycle = &cobra.Command{
	Use:   "CyclePorts",
	Short: "Cycle a set of switch ports.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, serial := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		if serial == "" {
			serial = args[1]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := livetools.PostCyclePorts(networkId, serial, format)
		shell.Display(metadata, "CyclePorts", cmd.Flags())
	},
}