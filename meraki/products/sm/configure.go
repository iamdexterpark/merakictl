package sm

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/sm/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var GetApnsCert = &cobra.Command{
	Use:   "ApnsCert",
	Short: "Get the organization's APNS certificate.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
	orgId = args[0]
}
		metadata := configure.GetAPNSCertificate(orgId)
		shell.Display(metadata, "ApnsCert", cmd.Flags())
	},
}

var GetBypassActivationLockAttempts = &cobra.Command{
Use:   "BypassActivationLockAttempts",
Short: "Bypass activation lock attempt status.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}

	attemptId := args[0]

metadata := configure.GetBypassActivationLockAttempts(networkId, attemptId)
shell.Display(metadata, "BypassActivationLockAttempts", cmd.Flags())
},
}


var GetCerts = &cobra.Command{
Use:   "Certs",
Short: "List the certs on a device",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}

	deviceId := args[0]
metadata := configure.GetDeviceCerts(networkId, deviceId)
shell.Display(metadata, "Certs", cmd.Flags())
},
}


var GetDeviceProfiles = &cobra.Command{
Use:   "DeviceProfiles",
Short: "Get the profiles associated with a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceProfiles(networkId, deviceId)
shell.Display(metadata, "DeviceProfiles", cmd.Flags())
},
}

var GetNetworkAdapters = &cobra.Command{
Use:   "NetworkAdapters",
Short: "List the network adapters of a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetNetworkAdapters(networkId, deviceId)
shell.Display(metadata, "NetworkAdapters", cmd.Flags())
},
}


var GetRestrictions = &cobra.Command{
Use:   "Restrictions",
Short: "List the restrictions on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceRestrictions(networkId, deviceId)
shell.Display(metadata, "Restrictions", cmd.Flags())
},
}


var GetSecurityCenters = &cobra.Command{
Use:   "SecurityCenters",
Short: "List the security centers on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceSecurityCenters(networkId, deviceId)
shell.Display(metadata, "SecurityCenters", cmd.Flags())
},
}


var GetDeviceSoftware = &cobra.Command{
Use:   "DeviceSoftware",
Short: "Get a list of softwares associated with a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceSoftwares(networkId, deviceId)
shell.Display(metadata, "DeviceSoftware", cmd.Flags())
},
}


var GetWlanLists = &cobra.Command{
Use:   "WlanLists",
Short: "List the saved SSID names on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceWlanLists(networkId, deviceId)
shell.Display(metadata, "WlanLists", cmd.Flags())
},
}


var GetDevices = &cobra.Command{
Use:   "Devices",
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
shell.Display(metadata, "Devices", cmd.Flags())
},
}


var GetProfiles = &cobra.Command{
Use:   "Profiles",
Short: "List all profiles in a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetProfiles(networkId)
shell.Display(metadata, "Profiles", cmd.Flags())
},
}


var GetTargetGroups = &cobra.Command{
Use:   "TargetGroups",
Short: "List the target groups in this network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}

	withDetails, _ := cmd.Flags().GetString("withDetails")

metadata := configure.GetTargetGroups(networkId, withDetails)
shell.Display(metadata, "TargetGroups", cmd.Flags())
},
}

var GetTargetGroup = &cobra.Command{
	Use:   "TargetGroup",
	Short: "Return a target group.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		targetGroupId := args[0]
		withDetails, _ := cmd.Flags().GetString("withDetails")

		metadata := configure.GetTargetGroup(networkId, targetGroupId, withDetails)
		shell.Display(metadata, "TargetGroup", cmd.Flags())
	},
}

var GetUserSoftware = &cobra.Command{
Use:   "UserSoftware",
Short: "Get a list of softwares associated with a user.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	userId  := args[0]

metadata := configure.GetSmUserSoftware(networkId, userId)
shell.Display(metadata, "UserSoftware", cmd.Flags())
},
}


var GetUsers = &cobra.Command{
Use:   "Users",
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
shell.Display(metadata, "Users", cmd.Flags())
},
}


var GetVPPAccounts = &cobra.Command{
Use:   "VPPAccounts",
Short: "Get a hash containing the unparsed token of the VPP account with the given ID.",
Run: func(cmd *cobra.Command, args []string) {
orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
	orgId = args[0]
}
metadata := configure.GetVPPAccounts(orgId)
shell.Display(metadata, "VPPAccounts", cmd.Flags())
},
}

var GetVPPAccount = &cobra.Command{
	Use:   "VPPAccount",
	Short: "List the VPP accounts in the organization.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[1]
		}

		vppAccountId := args[0]
		metadata := configure.GetVPPAccount(orgId, vppAccountId)
		shell.Display(metadata, "VPPAccount", cmd.Flags())
	},
}