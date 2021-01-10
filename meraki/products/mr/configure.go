package mr

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/wireless/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)



var GetAlternateManagementInterface = &cobra.Command{
Use:   "alternateManagementInterface",
Short: "Return alternate management interface and devices with IP assigned.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetAlternateManagementInterface(networkId)
shell.Display(metadata, "AlternateManagementInterface", cmd.Flags())
},
}

var PutAlternateManagementInterface = &cobra.Command{
	Use:   "alternateManagementInterface",
	Short: "Return alternate management interface and devices with IP assigned.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.AlternateManagementInterface
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutAlternateManagementInterface(networkId,  input)
		shell.Display(metadata, "AlternateManagementInterface", cmd.Flags())
	},
}

var GetBluetoothNetworkSettings = &cobra.Command{
Use:   "bluetoothNetworkSettings",
Short: "Return the Bluetooth settings for a network. Bluetooth settings must be enabled on the network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[0]
}
metadata := configure.GetBluetoothNetworkSettings(networkId)
shell.Display(metadata, "BluetoothNetworkSettings", cmd.Flags())
},
}

var PutBluetoothNetworkSettings = &cobra.Command{
	Use:   "bluetoothNetworkSettings",
	Short: "Return the Bluetooth settings for a network. Bluetooth settings must be enabled on the network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.BluetoothNetworkSettings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutBluetoothNetworkSettings(networkId,  input)
		shell.Display(metadata, "BluetoothNetworkSettings", cmd.Flags())
	},
}

var GetBluetoothDeviceSettings = &cobra.Command{
Use:   "bluetoothDeviceSettings",
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

var PutBluetoothDeviceSettings = &cobra.Command{
	Use:   "bluetoothDeviceSettings",
	Short: "Return the bluetooth settings for a mr device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.BluetoothDeviceSettings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutBluetoothDeviceSettings(serial,  input)
		shell.Display(metadata, "BluetoothDeviceSettings", cmd.Flags())
	},
}

var GetRadioSettings = &cobra.Command{
Use:   "radioSettings",
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

var PutRadioSettings = &cobra.Command{
	Use:   "radioSettings",
	Short: "Return the radio settings of a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.RadioSettings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutRadioSettings(serial,  input)
		shell.Display(metadata, "RadioSettings", cmd.Flags())
	},
}

var GetRFProfiles = &cobra.Command{
Use:   "rfProfiles",
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
Use:   "rfProfile",
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

var DelRFProfile = &cobra.Command{
	Use:   "rfProfile",
	Short: "Return a RF profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		rfProfileId := args[0]
		metadata := configure.DelRFProfile(networkId, rfProfileId)
		shell.Display(metadata, "RFProfile", cmd.Flags())
	},
}

var PutRFProfile = &cobra.Command{
	Use:   "rfProfile",
	Short: "Return a RF profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		rfProfileId := args[0]
		var format configure.RFProfile
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutRFProfile(networkId, rfProfileId,  input)
		shell.Display(metadata, "RFProfile", cmd.Flags())
	},
}

var PostRFProfile = &cobra.Command{
	Use:   "rfProfile",
	Short: "Return a RF profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.RFProfile
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostRFProfile(networkId,  input)
		shell.Display(metadata, "RFProfile", cmd.Flags())
	},
}

var GetWirelessSettings = &cobra.Command{
Use:   "wirelessSettings",
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

var PutWirelessSettings = &cobra.Command{
	Use:   "wirelessSettings",
	Short: "Return the mr settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Settings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSettings(networkId,  input)
		shell.Display(metadata, "WirelessSettings", cmd.Flags())
	},
}

var GetL3FirewallRules = &cobra.Command{
Use:   "l3FirewallRules",
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

var PutL3FirewallRules = &cobra.Command{
	Use:   "l3FirewallRules",
	Short: "Return the L3 firewall rules for an SSID on an MR network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		number := args[0]
		var format configure.L3FirewallRules
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutL3FirewallRules(networkId, number,  input)
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

var PutL7FirewallRules = &cobra.Command{
	Use:   "l7FirewallRules",
	Short: "Return the L7 firewall rules for an SSID on an MR network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		number := args[0]
		var format configure.L7FirewallRules
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutL7FirewallRules(networkId, number,  input)
		shell.Display(metadata, "L7FirewallRules", cmd.Flags())
	},
}

var GetIdentityPSKs = &cobra.Command{
Use:   "identityPSKs",
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
	Use:   "identityPSK",
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

var DelIdentityPSK = &cobra.Command{
	Use:   "identityPSK",
	Short: "Return an Identity PSK.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}

		number := args[0]
		identityPskId := args[1]
		metadata := configure.DelIdentityPSK(networkId,
			number, identityPskId)
		shell.Display(metadata, "IdentityPSK", cmd.Flags())
	},
}

var PutIdentityPSK = &cobra.Command{
	Use:   "identityPSK",
	Short: "Return an Identity PSK.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[2]
		}

		number := args[0]
		identityPskId := args[1]
		var format configure.IdentityPSK
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutIdentityPSK(networkId,
			number, identityPskId,  input)
		shell.Display(metadata, "IdentityPSK", cmd.Flags())
	},
}

var PostIdentityPSK = &cobra.Command{
	Use:   "identityPSK",
	Short: "Return an Identity PSK.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		number := args[0]
		var format configure.IdentityPSK
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostIdentityPSK(networkId,
			number,  input)
		shell.Display(metadata, "IdentityPSK", cmd.Flags())
	},
}

var GetSplashSettings = &cobra.Command{
Use:   "splashSettings",
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

var PutSplashSettings = &cobra.Command{
	Use:   "splashSettings",
	Short: "Display the splash page settings for the given SSID.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		number := args[0]
		var format configure.SplashPageSettings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSplashPageSettings(networkId, number,  input)
		shell.Display(metadata, "SplashSettings", cmd.Flags())
	},
}

var GetTrafficShapingRules = &cobra.Command{
Use:   "trafficShapingRules",
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

var PutTrafficShapingRules = &cobra.Command{
	Use:   "trafficShapingRules",
	Short: "Display the traffic shaping settings for a SSID on an MR network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		number := args[0]
		var format configure.TrafficShapingRules
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutTrafficShapingRules(networkId, number,  input)
		shell.Display(metadata, "TrafficShapingRules", cmd.Flags())
	},
}

var GetSSIDs = &cobra.Command{
	Use:   "ssids",
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
	Use:   "ssid",
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
	Use:   "ssid",
	Short: "Update The MR SSIDs In A Network.",
	Long:  `merakictl  put mr ssid {SSID#} -n {NetworkId} --input {FILENAME}.yaml`,
	Run: func(cmd *cobra.Command, args []string) {

		// Flags
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		// SSID Number
		ssidNumber := args[0]
		var format configure.SSID
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSSID(networkId, ssidNumber,  input)
		shell.Display(metadata, "SSID", cmd.Flags())

	},
}
