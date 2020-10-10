package device


import (
	"github.com/ddexterpark/merakictl/api/devices/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// management
var management = &cobra.Command{
	Use:   "management",
	Short: "Return The Management Interface Settings For A Device.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		net := cmd.Flag("device").Value.String()
		managementInterface, traceback := configure.GetManagementInterface(net)
		shell.Display(managementInterface, traceback, "ManagementInterface", cmd.Flags())
	},
}
