package device

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/devices/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

var GetManagementInterface = &cobra.Command{
	Use:   "managementInterface",
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
	Use:   "managementInterface",
	Short: "Update the management interface settings for a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.ManagementInterface
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutManagementInterface(serial,  input)
		shell.Display(metadata, "ManagementInterface", cmd.Flags())
	},
}

var GetDevice = &cobra.Command{
	Use:   "device",
	Short: "Return A Single Device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetDevice(serial)
		shell.Display(metadata, "Device", cmd.Flags())
	},
}

var PutDevice = &cobra.Command{
	Use:   "device",
	Short: "Update the attributes of a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.Device
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutDevice(serial,  input)
		shell.Display(metadata, "Device", cmd.Flags())
	},
}