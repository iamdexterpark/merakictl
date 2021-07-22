package delete

import (
	"github.com/ddexterpark/merakictl/meraki/general/network"
	"github.com/ddexterpark/merakictl/meraki/general/organization"
	"github.com/ddexterpark/merakictl/meraki/products/insight"
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
	OrgCmd.AddCommand(organization.DelActionBatch)
	OrgCmd.AddCommand(organization.DelAdmins)
	OrgCmd.AddCommand(organization.DelBrandingPolicy)
	OrgCmd.AddCommand(organization.DelConfigurationTemplate)
	OrgCmd.AddCommand(organization.DelIDP)
	OrgCmd.AddCommand(organization.DelOrganization)
	OrgCmd.AddCommand(organization.DelSamlRole)
}

var NetCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net"},
	Short:   "Entrypoint for general network subcommands.",
}

func init() {
	NetCmd.AddCommand(network.DelFloorPlan)
	NetCmd.AddCommand(network.DelGroupPolicy)
	NetCmd.AddCommand(network.DelHTTPServer)
	NetCmd.AddCommand(network.DelMerakiAuthUser)
	NetCmd.AddCommand(network.DelMQTTBroker)
	NetCmd.AddCommand(network.DelNetwork)
	NetCmd.AddCommand(network.DelPIIRequest)
}

var DeviceCmd = &cobra.Command{
	Use:   "device",
	Aliases: []string{"dev"},
	Short: "Root for general device CLI commands.",
	Long:  "Entrypoint for Device subcommands.",
}

func init() {
	//DeviceCmd.AddCommand()
}

var MXCmd = &cobra.Command{
	Use:   "appliance",
	Aliases: []string{"mx"},
	Short: "Root for all mx appliance CLI commands.",
	Long:  "Entrypoint for mx subcommands.",
}

func init() {
	MXCmd.AddCommand(mx.DelCustomPerformanceClass)
	MXCmd.AddCommand(mx.DelStaticRoute)
	MXCmd.AddCommand(mx.DelVLAN)
}

var MSCmd = &cobra.Command{
	Use:   "switch",
	Aliases: []string{"ms"},
	Short: "Root for all Switch CLI commands.",
}

func init() {
	MSCmd.AddCommand(ms.DelAccessPolicy)
	MSCmd.AddCommand(ms.DelL3Interface)
	MSCmd.AddCommand(ms.DelLinkAggregations)
	MSCmd.AddCommand(ms.DelPortSchedules)
	MSCmd.AddCommand(ms.DelQoSRule)
	MSCmd.AddCommand(ms.DelRendezvousPoint)
	MSCmd.AddCommand(ms.DelStackL3Interface)
	MSCmd.AddCommand(ms.DelStackStaticRoute)
	MSCmd.AddCommand(ms.DelStaticRoute)
	MSCmd.AddCommand(ms.DelSwitchStack)
}

var MRCmd = &cobra.Command{
	Use:   "wireless",
	Aliases: []string{"mr"},
	Short: "Root for all wireless MR CLI commands.",
}

func init() {
	MRCmd.AddCommand(mr.DelIdentityPSK)
	MRCmd.AddCommand(mr.DelRFProfile)
}

var MGCmd = &cobra.Command{
	Use:   "gateway",
	Aliases: []string{"mg"},
	Short: "Entrypoint for mg cellular gateway subcommands.",
}

func init() {
	// MGCmd.AddCommand(mg.)
}


var MVCmd = &cobra.Command{
	Use:   "camera",
	Aliases: []string{"mv"},
	Short:  "Entrypoint for mv camera subcommands.",
}

func init() {
	MVCmd.AddCommand(mv.DelQualityRetentionProfile)
}

var SMCmd = &cobra.Command{
	Use:   "systems",
	Aliases: []string{"sm"},
	Short: "Entrypoint for SM subcommands.",
}

func init() {
	SMCmd.AddCommand(sm.DelTargetGroup)
}

var InsightCmd = &cobra.Command{
	Use:   "insight",
	Aliases: []string{"in"},
	Short: "Entrypoint for Insight subcommands.",
}

func init() {
	InsightCmd.AddCommand(insight.DelMonitoredMediaServer)
}