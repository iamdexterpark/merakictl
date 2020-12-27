package update

import (
	"github.com/ddexterpark/merakictl/meraki/general/device"
	"github.com/ddexterpark/merakictl/meraki/general/network"
	"github.com/ddexterpark/merakictl/meraki/general/organization"
	"github.com/ddexterpark/merakictl/meraki/products/insight"
	"github.com/ddexterpark/merakictl/meraki/products/mg"
	"github.com/ddexterpark/merakictl/meraki/products/mr"
	"github.com/ddexterpark/merakictl/meraki/products/ms"
	"github.com/ddexterpark/merakictl/meraki/products/mv"
	"github.com/ddexterpark/merakictl/meraki/products/mx"
	"github.com/ddexterpark/merakictl/meraki/products/sm"
	"github.com/spf13/cobra"
)

var OrgCmd = &cobra.Command{
	Use:   "organization",
	Aliases: []string{"org"},
	Short: "Root for all organization CLI commands.",
}

func init() {
	OrgCmd.AddCommand(organization.PutOrganization)
	OrgCmd.AddCommand(organization.PutBrandingPolicyPriorities)
	OrgCmd.AddCommand(organization.PutBrandingPolicy)
	OrgCmd.AddCommand(organization.PutSNMP)
	OrgCmd.AddCommand(organization.PutActionBatch)
	OrgCmd.AddCommand(organization.PutAdmins)
	OrgCmd.AddCommand(organization.PutConfigurationTemplate)
	OrgCmd.AddCommand(organization.PutIDP)
	OrgCmd.AddCommand(organization.PutLicence)
	OrgCmd.AddCommand(organization.PutLoginSecurity)
	OrgCmd.AddCommand(organization.PutSAML)
	OrgCmd.AddCommand(organization.PutSamlRole)
}

var NetCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net"},
	Short:   "Entrypoint for general network subcommands.",
}

func init() {
	NetCmd.AddCommand(network.PutAlertConfig)
	NetCmd.AddCommand(network.PutClientSplashAuthorization)
	NetCmd.AddCommand(network.PutFirmwareUpgrades)
	NetCmd.AddCommand(network.PutFloorplan)
	NetCmd.AddCommand(network.PutGroupPolicy)
	NetCmd.AddCommand(network.PutHTTPServer)
	NetCmd.AddCommand(network.PutMerakiAuthUser)
	NetCmd.AddCommand(network.PutMQTTBroker)
	NetCmd.AddCommand(network.PutNetflow)
	NetCmd.AddCommand(network.PutNetwork)
	NetCmd.AddCommand(network.PutSettings)
	NetCmd.AddCommand(network.PutSNMP)
	NetCmd.AddCommand(network.PutSyslog)
	NetCmd.AddCommand(network.PutTrafficAnalysis)
}

var DeviceCmd = &cobra.Command{
	Use:   "device",
	Aliases: []string{"dev"},
	Short: "Root for general device CLI commands.",
	Long:  "Entrypoint for Device subcommands.",
}

func init() {
	DeviceCmd.AddCommand(device.PutDevice)
	DeviceCmd.AddCommand(device.PutManagementInterface)
}

var MXCmd = &cobra.Command{
	Use:   "appliance",
	Aliases: []string{"mx"},
	Short: "Root for all mx appliance CLI commands.",
	Long:  "Entrypoint for mx subcommands.",
}

func init() {
	MXCmd.AddCommand(mx.PutBGP)
	MXCmd.AddCommand(mx.PutPortForwardingRules)
	MXCmd.AddCommand(mx.PutPort)
	MXCmd.AddCommand(mx.PutCellularFirewallRules)
	MXCmd.AddCommand(mx.PutConnectivityMonitoringDestinations)
	MXCmd.AddCommand(mx.PutContentFiltering)
	MXCmd.AddCommand(mx.PutCustomPerformanceClass)
	MXCmd.AddCommand(mx.PutFirewalledService)
	MXCmd.AddCommand(mx.PutInboundFirewallRules)
	MXCmd.AddCommand(mx.PutL3FirewallRules)
	MXCmd.AddCommand(mx.PutL7FirewallRules)
	MXCmd.AddCommand(mx.PutMalware)
	MXCmd.AddCommand(mx.PutNetworkIntrusion)
	MXCmd.AddCommand(mx.PutOneToManyNatRules)
	MXCmd.AddCommand(mx.PutOneToOneNatRules)
	MXCmd.AddCommand(mx.PutOrganizationIntrusion)
	MXCmd.AddCommand(mx.PutSingleLan)
	MXCmd.AddCommand(mx.PutSiteToSiteVPN)
	MXCmd.AddCommand(mx.PutStaticRoute)
	MXCmd.AddCommand(mx.PutThirdPartyVPNPeers)
	MXCmd.AddCommand(mx.PutTrafficShaping)
	MXCmd.AddCommand(mx.PutTrafficShapingRules)
	MXCmd.AddCommand(mx.PutUplinkBandwidth)
	MXCmd.AddCommand(mx.PutUplinkSelection)
	MXCmd.AddCommand(mx.PutVLAN)
	MXCmd.AddCommand(mx.PutVLANSettings)
	MXCmd.AddCommand(mx.PutVPNFirewallRules)
	MXCmd.AddCommand(mx.PutWarmspare)
}

var MSCmd = &cobra.Command{
	Use:   "switch",
	Aliases: []string{"ms"},
	Short: "Root for all Switch CLI commands.",
}

func init() {
	MSCmd.AddCommand(ms.PutAccessControlLists)
	MSCmd.AddCommand(ms.PutStaticRoute)
	MSCmd.AddCommand(ms.PutDHCP)
	MSCmd.AddCommand(ms.PutSettings)
	MSCmd.AddCommand(ms.PutAccessPolicy)
	MSCmd.AddCommand(ms.PutDHCPServerPolicy)
	MSCmd.AddCommand(ms.PutL3Interface)
	MSCmd.AddCommand(ms.PutLinkAggregations)
	MSCmd.AddCommand(ms.PutMTU)
	MSCmd.AddCommand(ms.PutMulticast)
	MSCmd.AddCommand(ms.PutOSPF)
	MSCmd.AddCommand(ms.PutPortSchedules)
	MSCmd.AddCommand(ms.PutQoSRule)
	MSCmd.AddCommand(ms.PutQoSRuleOrder)
	MSCmd.AddCommand(ms.PutRendezvousPoint)
	MSCmd.AddCommand(ms.PutStackDHCP)
	MSCmd.AddCommand(ms.PutStackL3Interface)
	MSCmd.AddCommand(ms.PutStackStaticRoute)
	MSCmd.AddCommand(ms.PutStormControl)
	MSCmd.AddCommand(ms.PutSTP)
	MSCmd.AddCommand(ms.PutSwitchPort)
	MSCmd.AddCommand(ms.PutSwitchPortProfile)
	MSCmd.AddCommand(ms.PutWarmSpare)
}

var MRCmd = &cobra.Command{
	Use:   "wireless",
	Aliases: []string{"mr"},
	Short: "Root for all wireless MR CLI commands.",
}

func init() {
	MRCmd.AddCommand(mr.PutSSID)
	MRCmd.AddCommand(mr.PutTrafficShapingRules)
	MRCmd.AddCommand(mr.PutL7FirewallRules)
	MRCmd.AddCommand(mr.PutL3FirewallRules)
	MRCmd.AddCommand(mr.PutAlternateMGMTInterface)
	MRCmd.AddCommand(mr.PutBluetoothDeviceSettings)
	MRCmd.AddCommand(mr.PutBluetoothNetworkSettings)
	MRCmd.AddCommand(mr.PutIdentityPSK)
	MRCmd.AddCommand(mr.PutRadioSettings)
	MRCmd.AddCommand(mr.PutRFProfile)
	MRCmd.AddCommand(mr.PutSplashSettings)
	MRCmd.AddCommand(mr.PutWirelessSettings)
}

var MGCmd = &cobra.Command{
	Use:   "gateway",
	Aliases: []string{"mg"},
	Short: "Entrypoint for mg cellular gateway subcommands.",
}

func init() {
	MGCmd.AddCommand(mg.PutConnectivityMonitor)
	MGCmd.AddCommand(mg.PutDHCP)
	MGCmd.AddCommand(mg.PutLan)
	MGCmd.AddCommand(mg.PutPortForwardingRules)
	MGCmd.AddCommand(mg.PutSubnetPool)
	MGCmd.AddCommand(mg.PutUplink)
}


var MVCmd = &cobra.Command{
	Use:   "camera",
	Aliases: []string{"mv"},
	Short:  "Entrypoint for mv camera subcommands.",
}

func init() {
	MVCmd.AddCommand(mv.PutQualityAndRetention)
	MVCmd.AddCommand(mv.PutQualityRetentionProfile)
	MVCmd.AddCommand(mv.PutSense)
	MVCmd.AddCommand(mv.PutVideoSettings)
}

var SMCmd = &cobra.Command{
	Use:   "systems",
	Aliases: []string{"sm"},
	Short: "Entrypoint for SM subcommands.",
}

func init() {
	SMCmd.AddCommand(sm.PutFields)
	SMCmd.AddCommand(sm.PutTargetGroup)
}

var InsightCmd = &cobra.Command{
	Use:   "insight",
	Aliases: []string{"in"},
	Short: "Entrypoint for Insight subcommands.",
}

func init() {
	InsightCmd.AddCommand(insight.PutMonitoredMediaServer)

}
