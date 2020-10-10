package network

import (
	"github.com/spf13/cobra"
)


// NetCmd - Root for all organization CLI commands.
var NetCmd = &cobra.Command{
	Use:   "network",
	Short: "",
	Long: "Entrypoint for network subcommands.",
}

// init - Entrypoint for NetCmd sub-commands.
func init() {
	NetCmd.AddCommand(clients)
	NetCmd.AddCommand(alertconfig)
}
