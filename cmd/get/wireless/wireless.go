package wireless

import "github.com/spf13/cobra"

// MRCmd - Root for all organization CLI commands.
var MRCmd = &cobra.Command{
	Use:   "mr",
	Short: "",
	Long:  "List of available wifi specific API calls exposed via CLI.",
}

// init - Entrypoint for OrgCmd sub-commands.
func init() {
	MRCmd.AddCommand(ssids)
	MRCmd.AddCommand(ssid)
}
