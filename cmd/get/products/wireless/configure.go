package wireless

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/wireless/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)



var alternatemgmtinterface = &cobra.Command{
Use:   "alternatemgmtinterface",
Short: "Return alternate management interface and devices with IP assigned.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetAlternateManagementInterface(networkId)
shell.Display(metadata, "alternatemgmtinterface", cmd.Flags())
},
}

var bluetoothsettings = &cobra.Command{
Use:   "bluetoothsettings",
Short: "Return the Bluetooth settings for a network. Bluetooth settings must be enabled on the network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetBluetoothNetworkSettings(networkId)
shell.Display(metadata, "bluetoothsettings", cmd.Flags())
},
}


var bluetoothdevicesettings = &cobra.Command{
Use:   "bluetoothdevicesettings",
Short: "Return the bluetooth settings for a wireless device.",
Run: func(cmd *cobra.Command, args []string) {
_, _, serial := shell.ResolveFlags(cmd.Flags())
if serial == "" {
	serial = args[0]
}
metadata := configure.GetBluetoothDeviceSettings(serial)
shell.Display(metadata, "bluetoothdevicesettings", cmd.Flags())
},
}


var radiosettings = &cobra.Command{
Use:   "radiosettings",
Short: "Return the radio settings of a device.",
Run: func(cmd *cobra.Command, args []string) {
	_, _, serial := shell.ResolveFlags(cmd.Flags())
	if serial == "" {
		serial = args[0]
	}
metadata := configure.GetRadioSettings(serial)
shell.Display(metadata, "radiosettings", cmd.Flags())
},
}


var rfprofiles = &cobra.Command{
Use:   "rfprofiles",
Short: "List the non-basic RF profiles for this network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}

includeTemplateProfiles, _ := cmd.Flags().GetString("includeTemplateProfiles")

metadata := configure.GetRFProfiles(networkId, includeTemplateProfiles)
shell.Display(metadata, "rfprofiles", cmd.Flags())
},
}


var rfprofile = &cobra.Command{
Use:   "rfprofile",
Short: "Return a RF profile.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	rfProfileId := args[0]
metadata := configure.GetRFProfile(networkId, rfProfileId)
shell.Display(metadata, "rfprofile", cmd.Flags())
},
}


var wirelesssettings = &cobra.Command{
Use:   "wirelesssettings",
Short: "Return the wireless settings for a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetSettings(networkId)
shell.Display(metadata, "wirelesssettings", cmd.Flags())
},
}


var l3FirewallRules = &cobra.Command{
Use:   "l3FirewallRules",
Short: "Return the L3 firewall rules for an SSID on an MR network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

number := args[0]
metadata := configure.GetL3FirewallRules(networkId, number)
shell.Display(metadata, "l3FirewallRules", cmd.Flags())
},
}

var l7FirewallRules = &cobra.Command{
	Use:   "l7FirewallRules",
	Short: "Return the L7 firewall rules for an SSID on an MR network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		number := args[0]
		metadata := configure.GetL7FirewallRules(networkId, number)
		shell.Display(metadata, "l7FirewallRules", cmd.Flags())
	},
}

var identitypsks = &cobra.Command{
Use:   "identitypsks",
Short: "List all Identity PSKs in a wireless network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	number := args[0]
metadata := configure.GetIdentityPSKs(networkId, number)
shell.Display(metadata, "identitypsks", cmd.Flags())
},
}

var identitypsk = &cobra.Command{
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
		shell.Display(metadata, "identitypsk", cmd.Flags())
	},
}

var splashsettings = &cobra.Command{
Use:   "splashsettings",
Short: "Display the splash page settings for the given SSID.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}
	number := args[0]

metadata := configure.GetSplashPageSettings(networkId, number)
shell.Display(metadata, "splashsettings", cmd.Flags())
},
}


var trafficshapingrules = &cobra.Command{
Use:   "trafficshapingrules",
Short: "Display the traffic shaping settings for a SSID on an MR network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}
	number := args[0]
metadata := configure.GetTrafficShapingRules(networkId, number)
shell.Display(metadata, "trafficshapingrules", cmd.Flags())
},
}


// ssids
var ssids = &cobra.Command{
	Use:   "ssids",
	Short: "List The MR SSIDs In A Network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetSSIDS(networkId)
		shell.Display(metadata, "ssids", cmd.Flags())
	},
}

// ssid
var ssid = &cobra.Command{
	Use:   "ssid",
	Short: "List A Single MR SSID In A Network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		ssidNumber := args[0]
		metadata := configure.GetSSID(networkId, ssidNumber)
		shell.Display(metadata, "ssid", cmd.Flags())
	},
}