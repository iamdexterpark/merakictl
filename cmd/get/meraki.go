package get

import (
	"github.com/ddexterpark/merakictl/meraki/general/device"
	"github.com/ddexterpark/merakictl/meraki/general/network"
	"github.com/ddexterpark/merakictl/meraki/general/organization"
	"github.com/ddexterpark/merakictl/meraki/products/mx"
	"github.com/ddexterpark/merakictl/meraki/products/mv"
	"github.com/ddexterpark/merakictl/meraki/products/mg"
	"github.com/ddexterpark/merakictl/meraki/products/insight"
	"github.com/ddexterpark/merakictl/meraki/products/sm"
	"github.com/ddexterpark/merakictl/meraki/products/ms"
	"github.com/ddexterpark/merakictl/meraki/products/mr"
	"github.com/spf13/cobra"
)

var OrgCmd = &cobra.Command{
	Use:   "organization",
	Aliases: []string{"org"},
	Short: "Root for all organization CLI commands.",
}

func init() {
	OrgCmd.AddCommand(organization.GetList)
	OrgCmd.AddCommand(organization.GetDetail)
	OrgCmd.AddCommand(organization.GetActionBatches)
	OrgCmd.AddCommand(organization.GetActionBatch)
	OrgCmd.AddCommand(organization.GetAdmins)
	OrgCmd.AddCommand(organization.GetBrandingPolicies)
	OrgCmd.AddCommand(organization.GetBrandingPolicy)
	OrgCmd.AddCommand(organization.GetConfigurationTemplates)
	OrgCmd.AddCommand(organization.GetConfigurationTemplate)
	OrgCmd.AddCommand(organization.GetDevices)
	OrgCmd.AddCommand(organization.GetInventoryDevices)
	OrgCmd.AddCommand(organization.GetInventoryDevice)
	OrgCmd.AddCommand(organization.GetLicences)
	OrgCmd.AddCommand(organization.GetLicence)
	OrgCmd.AddCommand(organization.GetLoginSecurity)
	OrgCmd.AddCommand(organization.GetNetworks)
	OrgCmd.AddCommand(organization.GetIDPS)
	OrgCmd.AddCommand(organization.GetIDP)
	OrgCmd.AddCommand(organization.GetSAML)
	OrgCmd.AddCommand(organization.GetSamlRoles)
	OrgCmd.AddCommand(organization.GetSamlRole)
	OrgCmd.AddCommand(organization.GetSNMP)
	OrgCmd.AddCommand(organization.GetApiRequestsOverview)
	OrgCmd.AddCommand(organization.GetApiRequests)
	OrgCmd.AddCommand(organization.GetConfigurationChanges)
	OrgCmd.AddCommand(organization.GetDeviceStatuses)
	OrgCmd.AddCommand(organization.GetUplinksLossAndLatency)
	OrgCmd.AddCommand(organization.GetLicenseOverview)
	OrgCmd.AddCommand(organization.GetOpenAPISpec)
	OrgCmd.AddCommand(organization.GetAlertTypes)
	OrgCmd.AddCommand(organization.GetWebhookLogs)
}

var NetCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net"},
	Short:   "Entrypoint for general network subcommands.",
}

func init() {
	NetCmd.AddCommand(network.GetAlertConfig)
	NetCmd.AddCommand(network.GetClients)
	NetCmd.AddCommand(network.GetClientPolicy)
	NetCmd.AddCommand(network.GetClientSplashAuthorization)
	NetCmd.AddCommand(network.GetDevices)
	NetCmd.AddCommand(network.GetFirmwareUpgrades)
	NetCmd.AddCommand(network.GetFloorplans)
	NetCmd.AddCommand(network.GetFloorplan)
	NetCmd.AddCommand(network.GetGroupPolicies)
	NetCmd.AddCommand(network.GetGroupPolicy)
	NetCmd.AddCommand(network.GetMerakiAuthUsers)
	NetCmd.AddCommand(network.GetMerakiAuthUser)
	NetCmd.AddCommand(network.GetMQTTBrokers)
	NetCmd.AddCommand(network.GetMQTTBroker)
	NetCmd.AddCommand(network.GetNetflow)
	NetCmd.AddCommand(network.GetChannelUtilization)
	NetCmd.AddCommand(network.GetPIIRequests)
	NetCmd.AddCommand(network.GetPIIRequest)
	NetCmd.AddCommand(network.GetSMDevices)
	NetCmd.AddCommand(network.GetSMOwners)
	NetCmd.AddCommand(network.GetSettings)
	NetCmd.AddCommand(network.GetSNMP)
	NetCmd.AddCommand(network.GetSyslog)
	NetCmd.AddCommand(network.GetTrafficAnalysis)
	NetCmd.AddCommand(network.GetTrafficShaping)
	NetCmd.AddCommand(network.GetDSCP)
	NetCmd.AddCommand(network.GetHTTPServers)
	NetCmd.AddCommand(network.GetHTTPServer)
	NetCmd.AddCommand(network.GetWebhookTest)
	NetCmd.AddCommand(network.GetBluetoothClients)
	NetCmd.AddCommand(network.GetBluetoothClient)
	NetCmd.AddCommand(network.GetClientTrafficHistory)
	NetCmd.AddCommand(network.GetClientUsageHistory)
	NetCmd.AddCommand(network.GetNetworkClients)
	NetCmd.AddCommand(network.GetClientIdentifier)
	NetCmd.AddCommand(network.GetEnvironmentalEvents)
	NetCmd.AddCommand(network.GetEvents)
	NetCmd.AddCommand(network.GetSplashLoginAttempts)
	NetCmd.AddCommand(network.GetTraffic)
}

var DeviceCmd = &cobra.Command{
	Use:   "device",
	Aliases: []string{"dev"},
	Short: "Root for general device CLI commands.",
	Long:  "Entrypoint for Device subcommands.",
}

func init() {
	DeviceCmd.AddCommand(device.GetManagementInterface)
	DeviceCmd.AddCommand(device.GetDevice)
	DeviceCmd.AddCommand(device.GetClients)
	DeviceCmd.AddCommand(device.GetLLDPcdp)
	DeviceCmd.AddCommand(device.GetUplink)
}

var MXCmd = &cobra.Command{
	Use:   "appliance",
	Aliases: []string{"mx"},
	Short: "Root for all mx appliance CLI commands.",
	Long:  "Entrypoint for mx subcommands.",
}

func init() {
	MXCmd.AddCommand(mx.GetConnectivityMonitoringDestinations)
	MXCmd.AddCommand(mx.GetContentFilteringCategories)
	MXCmd.AddCommand(mx.GetContentFiltering)
	MXCmd.AddCommand(mx.GetCellularFirewallRules)
	MXCmd.AddCommand(mx.GetFirewalledServices)
	MXCmd.AddCommand(mx.GetFirewalledService)
	MXCmd.AddCommand(mx.GetInboundFirewallRules)
	MXCmd.AddCommand(mx.GetL3FirewallRules)
	MXCmd.AddCommand(mx.GetL7ApplicationCategories)
	MXCmd.AddCommand(mx.GetL7FirewallRules)
	MXCmd.AddCommand(mx.GetOneToManyNatRules)
	MXCmd.AddCommand(mx.GetOneToOneNatRules)
	MXCmd.AddCommand(mx.GetPortForwardingRules)
	MXCmd.AddCommand(mx.GetPorts)
	MXCmd.AddCommand(mx.GetPort)
	MXCmd.AddCommand(mx.GetNetworkIntrusion)
	MXCmd.AddCommand(mx.GetOrganizationIntrusion)
	MXCmd.AddCommand(mx.GetMalware)
	MXCmd.AddCommand(mx.GetSettings)
	MXCmd.AddCommand(mx.GetSingleLan)
	MXCmd.AddCommand(mx.GetStaticRoutes)
	MXCmd.AddCommand(mx.GetStaticRoute)
	MXCmd.AddCommand(mx.GetCustomPerformanceClasses)
	MXCmd.AddCommand(mx.GetCustomPerformanceClass)
	MXCmd.AddCommand(mx.GetTrafficShapingRules)
	MXCmd.AddCommand(mx.GetUplinkBandwidth)
	MXCmd.AddCommand(mx.GetUplinkSelection)
	MXCmd.AddCommand(mx.GetTrafficShaping)
	MXCmd.AddCommand(mx.GetVLANSettings)
	MXCmd.AddCommand(mx.GetVLANs)
	MXCmd.AddCommand(mx.GetVLAN)
	MXCmd.AddCommand(mx.GetSiteToSiteVPN)
	MXCmd.AddCommand(mx.GetThirdPartyVPNPeers)
	MXCmd.AddCommand(mx.GetVPNFirewallRules)
	MXCmd.AddCommand(mx.GetWarmspare)
	MXCmd.AddCommand(mx.GetSecurityEvents)
	MXCmd.AddCommand(mx.GetDHCPSubnets)
	MXCmd.AddCommand(mx.GetPerformance)
	MXCmd.AddCommand(mx.GetNetworkSecurityEvents)
	MXCmd.AddCommand(mx.GetOrganizationSecurityEvents)
	MXCmd.AddCommand(mx.GetUplinkStatuses)
	MXCmd.AddCommand(mx.GetVPNStats)
	MXCmd.AddCommand(mx.GetVPNStatuses)
}

var MSCmd = &cobra.Command{
	Use:   "switch",
	Aliases: []string{"ms"},
	Short: "Root for all Switch CLI commands.",
}

func init() {
	MSCmd.AddCommand(ms.GetAccessControlLists)
	MSCmd.AddCommand(ms.GetAccessPolicies)
	MSCmd.AddCommand(ms.GetAccessPolicy)
	MSCmd.AddCommand(ms.GetSwitchPortsProfiles)
	MSCmd.AddCommand(ms.GetSwitchPortProfile)
	MSCmd.AddCommand(ms.GetSwitchProfiles)
	MSCmd.AddCommand(ms.GetDHCPServerPolicy)
	MSCmd.AddCommand(ms.GetDSCP)
	MSCmd.AddCommand(ms.GetLinkAggregations)
	MSCmd.AddCommand(ms.GetMTU)
	MSCmd.AddCommand(ms.GetPortSchedules)
	MSCmd.AddCommand(ms.GetSwitchPorts)
	MSCmd.AddCommand(ms.GetSwitchPort)
	MSCmd.AddCommand(ms.GetQoSRuleOrder)
	MSCmd.AddCommand(ms.GetQoSRules)
	MSCmd.AddCommand(ms.GetQoSRule)
	MSCmd.AddCommand(ms.GetDHCP)
	MSCmd.AddCommand(ms.GetL3Interfaces)
	MSCmd.AddCommand(ms.GetL3Interface)
	MSCmd.AddCommand(ms.GetRendezvousPoints)
	MSCmd.AddCommand(ms.GetRendezvousPoint)
	MSCmd.AddCommand(ms.GetMulticast)
	MSCmd.AddCommand(ms.GetOSPF)
	MSCmd.AddCommand(ms.GetStaticRoutes)
	MSCmd.AddCommand(ms.GetStaticRoute)
	MSCmd.AddCommand(ms.GetSettings)
	MSCmd.AddCommand(ms.GetStackDHCP)
	MSCmd.AddCommand(ms.GetStackL3Interfaces)
	MSCmd.AddCommand(ms.GetStackL3Interface)
	MSCmd.AddCommand(ms.GetStackStaticRoutes)
	MSCmd.AddCommand(ms.GetStackStaticRoute)
	MSCmd.AddCommand(ms.GetSwitchStacks)
	MSCmd.AddCommand(ms.GetSwitchStack)
	MSCmd.AddCommand(ms.GetStormControl)
	MSCmd.AddCommand(ms.GetSTP)
	MSCmd.AddCommand(ms.GetWarmSpare)
	MSCmd.AddCommand(ms.GetPackets)
	MSCmd.AddCommand(ms.GetPortsStatuses)
}

var MRCmd = &cobra.Command{
	Use:   "wireless",
	Aliases: []string{"mr"},
	Short: "Root for all wireless MR CLI commands.",
}

func init() {
	MRCmd.AddCommand(mr.GetAlternateMGMTInterface)
	MRCmd.AddCommand(mr.GetBluetoothSettings)
	MRCmd.AddCommand(mr.GetBluetoothDeviceSettings)
	MRCmd.AddCommand(mr.GetRadioSettings)
	MRCmd.AddCommand(mr.GetRFProfiles)
	MRCmd.AddCommand(mr.GetRFProfile)
	MRCmd.AddCommand(mr.GetWirelessSettings)
	MRCmd.AddCommand(mr.GetL3FirewallRules)
	MRCmd.AddCommand(mr.GetL7FirewallRules)
	MRCmd.AddCommand(mr.GetIdentityPSKs)
	MRCmd.AddCommand(mr.GetIdentityPSK)
	MRCmd.AddCommand(mr.GetSplashSettings)
	MRCmd.AddCommand(mr.GetTrafficShapingRules)
	MRCmd.AddCommand(mr.GetSSIDs)
	MRCmd.AddCommand(mr.GetSSID)
	MRCmd.AddCommand(mr.GetAirMarshal)
	MRCmd.AddCommand(mr.GetChannelUtilizationHistory)
	MRCmd.AddCommand(mr.GetClientCountHistory)
	MRCmd.AddCommand(mr.GetConnectionStat)
	MRCmd.AddCommand(mr.GetConnectionStats)
	MRCmd.AddCommand(mr.GetConnectivityEvents)
	MRCmd.AddCommand(mr.GetLatencyHistory)
	MRCmd.AddCommand(mr.GetLatencyStat)
	MRCmd.AddCommand(mr.GetLatencyStats)
	MRCmd.AddCommand(mr.GetDeviceConnectionStats)
	MRCmd.AddCommand(mr.GetNetworkConnectionStats)
	MRCmd.AddCommand(mr.GetDataRateHistory)
	MRCmd.AddCommand(mr.GetConnectionStats)
	MRCmd.AddCommand(mr.GetLatencyStats)
	MRCmd.AddCommand(mr.GetFailedConnections)
	MRCmd.AddCommand(mr.GetLatencyHistory)
	MRCmd.AddCommand(mr.GetLatencyHistoryAverage)
	MRCmd.AddCommand(mr.GetAggregatedLatencies)
	MRCmd.AddCommand(mr.GetDeviceLatencyStats)
	MRCmd.AddCommand(mr.GetNetworkLatencyStats)
	MRCmd.AddCommand(mr.GetMeshStatuses)
	MRCmd.AddCommand(mr.GetSignalQualityHistory)
	MRCmd.AddCommand(mr.GetStatus)
	MRCmd.AddCommand(mr.GetUsageHistory)
}

var MGCmd = &cobra.Command{
	Use:   "gateway",
	Aliases: []string{"mg"},
	Short: "Entrypoint for mg cellular gateway subcommands.",
}

func init() {
	MGCmd.AddCommand(mg.GetConnectivityMonitor)
	MGCmd.AddCommand(mg.GetDHCP)
	MGCmd.AddCommand(mg.GetLan)
	MGCmd.AddCommand(mg.GetPortForwardingRules)
	MGCmd.AddCommand(mg.GetSubnetPool)
	MGCmd.AddCommand(mg.GetUplink)
}


var MVCmd = &cobra.Command{
	Use:   "camera",
	Aliases: []string{"mv"},
	Short:  "Entrypoint for mv camera subcommands.",
}

func init() {
	MVCmd.AddCommand(mv.GetQualityAndRetention)
	MVCmd.AddCommand(mv.GetQualityRetentionProfiles)
	MVCmd.AddCommand(mv.GetQualityRetentionProfile)
	MVCmd.AddCommand(mv.GetSchedules)
	MVCmd.AddCommand(mv.GetObjectDetectionModels)
	MVCmd.AddCommand(mv.GetSense)
	MVCmd.AddCommand(mv.GetVideoSettings)
	MVCmd.AddCommand(mv.GetVideoLink)
	MVCmd.AddCommand(mv.GetLiveAnalytics)
	MVCmd.AddCommand(mv.GetAnalyticsOverview)
	MVCmd.AddCommand(mv.GetRecentAnalytics)
	MVCmd.AddCommand(mv.GetAnalyticsZonesHistory)
	MVCmd.AddCommand(mv.GetAnalyticsZones)
}

var SMCmd = &cobra.Command{
	Use:   "systems",
	Aliases: []string{"sm"},
	Short: "Entrypoint for SM subcommands.",
}

func init() {
	SMCmd.AddCommand(sm.GetApnsCert)
	SMCmd.AddCommand(sm.GetBypassActivationLockAttempts)
	SMCmd.AddCommand(sm.GetCerts)
	SMCmd.AddCommand(sm.GetDeviceProfiles)
	SMCmd.AddCommand(sm.GetNetworkAdapters)
	SMCmd.AddCommand(sm.GetRestrictions)
	SMCmd.AddCommand(sm.GetSecurityCenters)
	SMCmd.AddCommand(sm.GetDeviceSoftware)
	SMCmd.AddCommand(sm.GetWlanLists)
	SMCmd.AddCommand(sm.GetDevices)
	SMCmd.AddCommand(sm.GetProfiles)
	SMCmd.AddCommand(sm.GetTargetGroups)
	SMCmd.AddCommand(sm.GetTargetGroup)
	SMCmd.AddCommand(sm.GetDeviceProfiles)
	SMCmd.AddCommand(sm.GetUsers)
	SMCmd.AddCommand(sm.GetUserSoftware)
	SMCmd.AddCommand(sm.GetVPPAccount)
	SMCmd.AddCommand(sm.GetVPPAccounts)
	SMCmd.AddCommand(sm.GetCellularUsageHistory)
	SMCmd.AddCommand(sm.GetConnectivity)
	SMCmd.AddCommand(sm.GetDesktopLogs)
	SMCmd.AddCommand(sm.GetDeviceCommandLogs)
	SMCmd.AddCommand(sm.GetPerformanceHistory)

}

var InsightCmd = &cobra.Command{
	Use:   "insight",
	Aliases: []string{"in"},
	Short: "Entrypoint for Insight subcommands.",
}

func init() {
	InsightCmd.AddCommand(insight.GetMonitoredMediaServers)
	InsightCmd.AddCommand(insight.GetMonitoredMediaServer)

}