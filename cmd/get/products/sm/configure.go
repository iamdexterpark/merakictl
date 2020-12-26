package sm

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/sm/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var apnscert = &cobra.Command{
	Use:   "apnscert",
	Short: "Get the organization's APNS certificate.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
	orgId = args[0]
}
		metadata := configure.GetAPNSCertificate(orgId)
		shell.Display(metadata, "apnscert", cmd.Flags())
	},
}

var bypassactivationlockattempts = &cobra.Command{
Use:   "bypassactivationlockattempts",
Short: "Bypass activation lock attempt status.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}

	attemptId := args[0]

metadata := configure.GetBypassActivationLockAttempts(networkId, attemptId)
shell.Display(metadata, "bypassactivationlockattempts", cmd.Flags())
},
}


var certs = &cobra.Command{
Use:   "certs",
Short: "List the certs on a device",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}

	deviceId := args[0]
metadata := configure.GetDeviceCerts(networkId, deviceId)
shell.Display(metadata, "certs", cmd.Flags())
},
}


var deviceprofiles = &cobra.Command{
Use:   "deviceprofiles",
Short: "Get the profiles associated with a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceProfiles(networkId, deviceId)
shell.Display(metadata, "deviceprofiles", cmd.Flags())
},
}

var networkadapters = &cobra.Command{
Use:   "networkadapters",
Short: "List the network adapters of a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetNetworkAdapters(networkId, deviceId)
shell.Display(metadata, "networkadapters", cmd.Flags())
},
}


var restrictions = &cobra.Command{
Use:   "restrictions",
Short: "List the restrictions on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceRestrictions(networkId, deviceId)
shell.Display(metadata, "restrictions", cmd.Flags())
},
}


var securitycenters = &cobra.Command{
Use:   "securitycenters",
Short: "List the security centers on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceSecurityCenters(networkId, deviceId)
shell.Display(metadata, "securitycenters", cmd.Flags())
},
}


var devicesoftware = &cobra.Command{
Use:   "devicesoftware",
Short: "Get a list of softwares associated with a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceSoftwares(networkId, deviceId)
shell.Display(metadata, "devicesoftware", cmd.Flags())
},
}


var wlanlists = &cobra.Command{
Use:   "wlanlists",
Short: "List the saved SSID names on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceWlanLists(networkId, deviceId)
shell.Display(metadata, "wlanlists", cmd.Flags())
},
}


var devices = &cobra.Command{
Use:   "devices",
Short: "List the devices enrolled in an SM network with various specified fields and filters.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}

	fields, _ := cmd.Flags().GetString("fields")
	wifiMacs, _ := cmd.Flags().GetString("wifiMacs")
	serials, _ := cmd.Flags().GetString("serials")
	ids, _ := cmd.Flags().GetString("ids")
	scope, _ := cmd.Flags().GetString("scope")
	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := configure.GetSMDevices(networkId,
	fields, wifiMacs, serials, ids, scope, perPage,
	startingAfter, endingBefore)
shell.Display(metadata, "devices", cmd.Flags())
},
}


var profiles = &cobra.Command{
Use:   "profiles",
Short: "List all profiles in a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetProfiles(networkId)
shell.Display(metadata, "profiles", cmd.Flags())
},
}


var targetgroups = &cobra.Command{
Use:   "targetgroups",
Short: "List the target groups in this network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}

	withDetails, _ := cmd.Flags().GetString("withDetails")

metadata := configure.GetTargetGroups(networkId, withDetails)
shell.Display(metadata, "targetgroups", cmd.Flags())
},
}

var targetgroup = &cobra.Command{
	Use:   "targetgroup",
	Short: "Return a target group.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		targetGroupId := args[0]
		withDetails, _ := cmd.Flags().GetString("withDetails")

		metadata := configure.GetTargetGroup(networkId,targetGroupId, withDetails)
		shell.Display(metadata, "targetgroup", cmd.Flags())
	},
}

var usersoftware = &cobra.Command{
Use:   "usersoftware",
Short: "Get a list of softwares associated with a user.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	userId  := args[0]

metadata := configure.GetSmUserSoftware(networkId, userId)
shell.Display(metadata, "usersoftware", cmd.Flags())
},
}


var users = &cobra.Command{
Use:   "users",
Short: "List the owners in an SM network with various specified fields and filters.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}

	ids, _ := cmd.Flags().GetString("ids")
	usernames, _ := cmd.Flags().GetString("usernames")
	emails, _ := cmd.Flags().GetString("emails")
	scope, _ := cmd.Flags().GetString("scope")

metadata := configure.GetSmUsers(networkId, ids, usernames, emails, scope)
shell.Display(metadata, "users", cmd.Flags())
},
}


var vppaccounts = &cobra.Command{
Use:   "vppaccounts",
Short: "Get a hash containing the unparsed token of the VPP account with the given ID.",
Run: func(cmd *cobra.Command, args []string) {
orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
	orgId = args[0]
}
metadata := configure.GetVPPAccounts(orgId)
shell.Display(metadata, "vppaccounts", cmd.Flags())
},
}

var vppaccount = &cobra.Command{
	Use:   "vppaccount",
	Short: "List the VPP accounts in the organization.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[1]
		}

		vppAccountId := args[0]
		metadata := configure.GetVPPAccount(orgId, vppAccountId)
		shell.Display(metadata, "vppaccount", cmd.Flags())
	},
}