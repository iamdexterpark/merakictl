package appliance

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/appliance/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var securityevents = &cobra.Command{
	Use:   "securityevents",
	Short: "List the security events for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		clientId := args[0]

		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		sortOrder, _ := cmd.Flags().GetString("sortOrder")
		metadata := monitor.GetSecurityEvents(networkId, clientId, t0,
			t1, timespan,
			perPage, startingAfter,endingBefore, sortOrder)
			shell.Display(metadata, "securityevents", cmd.Flags())
	},
}


var dhcpsubnets = &cobra.Command{
	Use:   "dhcpsubnets",
	Short: "Return the DHCP subnet information for an appliance.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}


		metadata := monitor.GetDHCP(serial)
		shell.Display(metadata, "dhcpsubnets", cmd.Flags())
	},
}



var performance = &cobra.Command{
Use:   "performance",
Short: "Return the performance score for a single MX. Only primary MX devices supported. If no data is available, a 204 error code is returned.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
serial = args[0]
}


metadata := monitor.GetPerformanceScore(serial)
shell.Display(metadata, "performance", cmd.Flags())
},
}


var networksecurityevents = &cobra.Command{
Use:   "networksecurityevents",
Short: "List the security events for a network.",
Run: func(cmd *cobra.Command, args []string) {
	orgId, _, _ := shell.ResolveFlags(cmd.Flags())
	if orgId == "" {
		orgId = args[0]
	}

	t0, _ := cmd.Flags().GetString("t0")
	t1, _ := cmd.Flags().GetString("t1")
	timespan, _ := cmd.Flags().GetString("timespan")
	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")
	sortOrder, _ := cmd.Flags().GetString("sortOrder")

metadata := monitor.GetNetworkSecurityEvents(orgId, t0,
	t1, timespan,
	perPage, startingAfter,endingBefore, sortOrder)
shell.Display(metadata, "networksecurityevents", cmd.Flags())
},
}



var organizationsecurityevents = &cobra.Command{
Use:   "organizationsecurityevents",
Short: "List the security events for an organization.",
Run: func(cmd *cobra.Command, args []string) {
	orgId, _, _ := shell.ResolveFlags(cmd.Flags())
	if orgId == "" {
		orgId = args[0]
	}
	t0, _ := cmd.Flags().GetString("t0")
	t1, _ := cmd.Flags().GetString("t1")
	timespan, _ := cmd.Flags().GetString("timespan")
	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")
	sortOrder, _ := cmd.Flags().GetString("sortOrder")

metadata := monitor.GetOrganizationSecurityEvents(orgId, t0,
	t1, timespan,
	perPage, startingAfter,endingBefore, sortOrder)
shell.Display(metadata, "organizationsecurityevents", cmd.Flags())
},
}



var uplinkstatuses = &cobra.Command{
Use:   "uplinkstatuses",
Short: "List the uplink status of every Meraki MX and Z series appliances in the organization.",
Run: func(cmd *cobra.Command, args []string) {
	orgId, _, _ := shell.ResolveFlags(cmd.Flags())
	if orgId == "" {
		orgId = args[0]
	}

	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := monitor.GetUplinkStatus(orgId, perPage,
	startingAfter,endingBefore)
shell.Display(metadata, "uplinkstatuses", cmd.Flags())
},
}



var vpnstats = &cobra.Command{
	Use:   "vpnstats",
	Short: "Show VPN history stat for networks in an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[0]
		}
		if networkId == "" {
			networkId = args[1]
		}
		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := monitor.GetVPNHistory(orgId, t0, t1, timespan, perPage,
			startingAfter, endingBefore, networkId)
		shell.Display(metadata, "vpnstats", cmd.Flags())
	},
}


var vpnstatuses = &cobra.Command{
	Use:   "vpnstatuses",
	Short: "Show VPN status for networks in an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[0]
		}
		if networkId == "" {
			networkId = args[1]
		}

		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := monitor.GetVPNStatus(orgId, perPage,
			startingAfter, endingBefore, networkId)
		shell.Display(metadata, "vpnstatuses", cmd.Flags())
	},
}