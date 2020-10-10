package device

import (
	"github.com/spf13/cobra"
)


// DeviceCmd - Root for all organization CLI commands.
var DeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "",
	Long: "Entrypoint for Device subcommands.",
}

// init - Entrypoint for DeviceCmd sub-commands.
func init() {
	DeviceCmd.AddCommand(management)
}

