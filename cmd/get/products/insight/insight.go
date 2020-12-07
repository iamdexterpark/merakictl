package insight

import "github.com/spf13/cobra"

// DeviceCmd - Root for all organization CLI commands.
var InsightCmd = &cobra.Command{
	Use:   "insight",
	Short: "",
	Long:  "Entrypoint for Insight subcommands.",
}

// init - Entrypoint for DeviceCmd sub-commands.
func init() {
	InsightCmd.AddCommand(monitoredmediaservers)
	InsightCmd.AddCommand(monitoredmediaserver)

}
