package insight

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/insight/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var monitoredmediaservers = &cobra.Command{
	Use:   "monitoredmediaservers",
	Short: "List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[0]
		}
		metadata := configure.GetMonitoredMediaServers(orgId)
		shell.Display(metadata, "monitoredmediaservers", cmd.Flags())
	},
}

var monitoredmediaserver = &cobra.Command{
	Use:   "monitoredmediaserver",
	Short: "Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[1]
		}
		monitoredMediaServerId := args[0]
		metadata := configure.GetMonitoredMediaServer(orgId, monitoredMediaServerId)
		shell.Display(metadata, "monitoredmediaserver", cmd.Flags())
	},
}
