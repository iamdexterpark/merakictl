package device

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/devices/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
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

var GetDetails = &cobra.Command{
	Use:   "Details",
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