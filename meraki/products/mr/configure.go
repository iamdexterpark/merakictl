package mr

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/wireless/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)



var GetAlternateMGMTInterface = &cobra.Command{
Use:   "AlternateMGMTInterface",
Short: "Return alternate management interface and devices with IP assigned.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetAlternateManagementInterface(networkId)
shell.Display(metadata, "AlternateMGMTInterface", cmd.Flags())
},
}

var GetBluetoothSettings = &cobra.Command{
Use:   "BluetoothSettings",
Short: "Return the Bluetooth settings for a network. Bluetooth settings must be enabled on the network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetBluetoothNetworkSettings(networkId)
shell.Display(metadata, "BluetoothSettings", cmd.Flags())
},
}


var GetBluetoothDeviceSettings = &cobra.Command{
Use:   "BluetoothDeviceSettings",
Short: "Return the bluetooth settings for a mr device.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}
metadata := configure.GetBluetoothDeviceSettings(serial)
shell.Display(metadata, "BluetoothDeviceSettings", cmd.Flags())
},
}


var GetRadioSettings = &cobra.Command{
Use:   "RadioSettings",
Short: "Return the radio settings of a device.",
Run: func(cmd *cobra.Command, args []string) {
	_, _, serial := shell.ResolveFlags(cmd.Flags())
	if serial == "" {
		serial = args[0]
	}
metadata := configure.GetRadioSettings(serial)
shell.Display(metadata, "RadioSettings", cmd.Flags())
},
}


var GetRFProfiles = &cobra.Command{
Use:   "RFProfiles",
Short: "List the non-basic RF profiles for this network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

includeTemplateProfiles, _ := cmd.Flags().GetString("includeTemplateProfiles")

metadata := configure.GetRFProfiles(networkId, includeTemplateProfiles)
shell.Display(metadata, "RFProfiles", cmd.Flags())
},
}


var GetRFProfile = &cobra.Command{
Use:   "RFProfile",
Short: "Return a RF profile.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	rfProfileId := args[0]
metadata := configure.GetRFProfile(networkId, rfProfileId)
shell.Display(metadata, "RFProfile", cmd.Flags())
},
}


var GetWirelessSettings = &cobra.Command{
Use:   "WirelessSettings",
Short: "Return the mr settings for a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetSettings(networkId)
shell.Display(metadata, "WirelessSettings", cmd.Flags())
},
}


var GetL3FirewallRules = &cobra.Command{
Use:   "L3FirewallRules",
Short: "Return the L3 firewall rules for an SSID on an MR network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

number := args[0]
metadata := configure.GetL3FirewallRules(networkId, number)
shell.Display(metadata, "L3FirewallRules", cmd.Flags())
},
}

var GetL7FirewallRules = &cobra.Command{
	Use:   "l7FirewallRules",
	Short: "Return the L7 firewall rules for an SSID on an MR network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		number := args[0]
		metadata := configure.GetL7FirewallRules(networkId, number)
		shell.Display(metadata, "L7FirewallRules", cmd.Flags())
	},
}

var GetIdentityPSKs = &cobra.Command{
Use:   "IdentityPSKs",
Short: "List all Identity PSKs in a mr network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	number := args[0]
metadata := configure.GetIdentityPSKs(networkId, number)
shell.Display(metadata, "IdentityPSKs", cmd.Flags())
},
}

var GetIdentityPSK = &cobra.Command{
	Use:   "identitypsk",
	Short: "Return an Identity PSK.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}

		number := args[0]
		identityPskId := args[1]
		metadata := configure.GetIdentityPSK(networkId,
			number, identityPskId)
		shell.Display(metadata, "IdentityPSK", cmd.Flags())
	},
}

var GetSplashSettings = &cobra.Command{
Use:   "SplashSettings",
Short: "Display the splash page settings for the given SSID.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}
	number := args[0]

metadata := configure.GetSplashPageSettings(networkId, number)
shell.Display(metadata, "SplashSettings", cmd.Flags())
},
}


var GetTrafficShapingRules = &cobra.Command{
Use:   "TrafficShapingRules",
Short: "Display the traffic shaping settings for a SSID on an MR network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}
	number := args[0]
metadata := configure.GetTrafficShapingRules(networkId, number)
shell.Display(metadata, "TrafficShapingRules", cmd.Flags())
},
}


var GetSSIDs = &cobra.Command{
	Use:   "SSIDs",
	Short: "List The MR SSIDs In A Network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetSSIDS(networkId)
		shell.Display(metadata, "SSIDs", cmd.Flags())
	},
}

var GetSSID = &cobra.Command{
	Use:   "SSID",
	Short: "List A Single MR SSID In A Network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		ssidNumber := args[0]
		metadata := configure.GetSSID(networkId, ssidNumber)
		shell.Display(metadata, "SSID", cmd.Flags())
	},
}


var PutSSID = &cobra.Command{
	Use:   "SSID",
	Short: "Update The MR SSIDs In A Network.",
	Long:  `merakictl  update mr ssid {SSID#} -n {NetworkId} --input {FILENAME}.yaml`,
	Run: func(cmd *cobra.Command, args []string) {

		// Flags
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		// SSID Number
		ssidNumber := args[0]

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutSSID(networkId, ssidNumber, format)
		shell.Display(metadata, "SSID", cmd.Flags())

	},
}
