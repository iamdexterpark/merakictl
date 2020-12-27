package device

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/devices/configure"
	shell "github.com/ddexterpark/merakictl/utilities"
	"github.com/spf13/cobra"
)

var GetManagementInterface = &cobra.Command{
	Use:   "ManagementInterface",
	Short: "Return The Management Interface Settings For A Device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetManagementInterface(serial)
		shell.Display(metadata, "ManagementInterface", cmd.Flags())
	},
}

var PutManagementInterface = &cobra.Command{
	Use:   "ManagementInterface",
	Short: "Update the management interface settings for a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutManagementInterface(serial,format)
		shell.Display(metadata, "ManagementInterface", cmd.Flags())
	},
}

var GetDevice = &cobra.Command{
	Use:   "Device",
	Short: "Return A Single Device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetDevice(serial)
		shell.Display(metadata, "device", cmd.Flags())
	},
}

var PutDevice = &cobra.Command{
	Use:   "Device",
	Short: "Update the attributes of a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutDevice(serial, format)
		shell.Display(metadata, "device", cmd.Flags())
	},
}
