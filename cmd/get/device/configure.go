package device


import (
	"github.com/ddexterpark/merakictl/api/general/devices/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// management
var management = &cobra.Command{
	Use:   "management",
	Short: "Return The Management Interface Settings For A Device.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		serial := cmd.Flag("device").Value.String()
		managementInterface, traceback := configure.GetManagementInterface(serial)
		shell.Display(managementInterface, traceback, "ManagementInterface", cmd.Flags())
	},
}

// device
var device = &cobra.Command{
	Use:   "device",
	Short: "Return A Single Device.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		serial := cmd.Flag("device").Value.String()
		device, traceback := configure.GetSingleDevice(serial)
		shell.Display(device, traceback, "device", cmd.Flags())
	},
}