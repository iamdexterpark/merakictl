package network

import (
	"github.com/ddexterpark/merakictl/api"
	"github.com/ddexterpark/merakictl/api/general/networks/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

// clients
var clients = &cobra.Command{
	Use:   "clients",
	Short: "List the clients that have used this network in the timespan.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		net := cmd.Flag("network").Value.String()
		metadata := api.GetNetworkClients(net, "", "",
			"", "", "", "")
		shell.Display(metadata, "clients", cmd.Flags())
	},
}

// alertconfig - Return The Alert Configuration For This Network
var alertconfig = &cobra.Command{
	Use:   "alertconfig",
	Short: "Return The Alert Configuration For This Network.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		net := cmd.Flag("network").Value.String()
		metadata := configure.GetNetworkAlertConfig(net)
		shell.Display(metadata, "clients", cmd.Flags())
	},
}

