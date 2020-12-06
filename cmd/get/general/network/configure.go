package network

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/networks/configure"
	organizations "github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

// clients
var clients = &cobra.Command{
	Use:   "clients",
	Short: "List the clients that have used this network in the timespan.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		net := cmd.Flag("network").Value.String()
		metadata := organizations.GetNetworkClients(net, "", "",
			"", "", "", "")
		shell.Display(metadata, "clients", cmd.Flags())
	},
}

// alertconfig - Return The Alert Configuration For This Network
var alertconfig = &cobra.Command{
	Use:   "alertconfig",
	Short: "Return The Alert Configuration For This Network.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		net := cmd.Flag("network").Value.String()
		metadata := configure.GetNetworkAlertConfig(net)
		shell.Display(metadata, "clients", cmd.Flags())
	},
}
