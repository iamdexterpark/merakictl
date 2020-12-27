package mx

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/appliance/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var GetSecurityEvents = &cobra.Command{
	Use:   "SecurityEvents",
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
		metadata := monitor.GetClientSecurityEvents(networkId, clientId, t0,
			t1, timespan,
			perPage, startingAfter, endingBefore, sortOrder)
			shell.Display(metadata, "SecurityEvents", cmd.Flags())
	},
}


var GetDHCPSubnets = &cobra.Command{
	Use:   "DHCPSubnets",
	Short: "Return the DHCP subnet information for an mx.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}


		metadata := monitor.GetDHCP(serial)
		shell.Display(metadata, "DHCPSubnets", cmd.Flags())
	},
}



var GetPerformance = &cobra.Command{
Use:   "Performance",
Short: "Return the performance score for a single MX. Only primary MX devices supported. If no data is available, a 204 error code is returned.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
serial = args[0]
}


metadata := monitor.GetPerformance(serial)
shell.Display(metadata, "Performance", cmd.Flags())
},
}


var GetNetworkSecurityEvents = &cobra.Command{
Use:   "NetworkSecurityEvents",
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
	perPage, startingAfter, endingBefore, sortOrder)
shell.Display(metadata, "NetworkSecurityEvents", cmd.Flags())
},
}



var GetOrganizationSecurityEvents = &cobra.Command{
Use:   "OrganizationSecurityEvents",
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
	perPage, startingAfter, endingBefore, sortOrder)
shell.Display(metadata, "OrganizationSecurityEvents", cmd.Flags())
},
}



var GetUplinkStatuses = &cobra.Command{
Use:   "UplinkStatuses",
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
	startingAfter, endingBefore)
shell.Display(metadata, "UplinkStatuses", cmd.Flags())
},
}



var GetVPNStats = &cobra.Command{
	Use:   "VPNStats",
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

		metadata := monitor.GetVpnStats(orgId, t0, t1, timespan, perPage,
			startingAfter, endingBefore, networkId)
		shell.Display(metadata, "VPNStats", cmd.Flags())
	},
}


var GetVPNStatuses = &cobra.Command{
	Use:   "VPNStatuses",
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
		shell.Display(metadata, "VPNStatuses", cmd.Flags())
	},
}