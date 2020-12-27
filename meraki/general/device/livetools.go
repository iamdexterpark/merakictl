package device

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/devices/livetools"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var PostBlinkLEDs = &cobra.Command{
	Use:   "BlinkLEDs",
	Short: "Blink the LEDs on a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := livetools.PostBlinkLeds(serial)
		shell.Display(metadata, "BlinkLEDs", cmd.Flags())
	},
}

var PostReboot = &cobra.Command{
	Use:   "Reboot",
	Short: "Reboot a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := livetools.PostReboot(serial)
		shell.Display(metadata, "Reboot", cmd.Flags())
	},
}