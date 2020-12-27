package create

import (
	"github.com/ddexterpark/merakictl/meraki/general/device"
	"github.com/ddexterpark/merakictl/meraki/general/network"
	"github.com/ddexterpark/merakictl/meraki/general/organization"
	"github.com/ddexterpark/merakictl/meraki/products/insight"

	//"github.com/ddexterpark/merakictl/meraki/products/sm"
	//"github.com/ddexterpark/merakictl/meraki/products/ms"
	//"github.com/ddexterpark/merakictl/meraki/products/mr"
	"github.com/spf13/cobra"
)

var OrgCmd = &cobra.Command{
	Use:   "organization",
	Aliases: []string{"org"},
	Short: "Root for all organization CLI commands.",
}

func init() {
	OrgCmd.AddCommand(organization.PostActionBatch)
	OrgCmd.AddCommand(organization.PostClone)
	OrgCmd.AddCommand(organization.PostCombineNetworks)
	OrgCmd.AddCommand(organization.PostOrganization)
	OrgCmd.AddCommand(organization.PostAdmins)
	OrgCmd.AddCommand(organization.PostBrandingPolicy)
	OrgCmd.AddCommand(organization.PostConfigurationTemplate)
	OrgCmd.AddCommand(organization.PostIDP)
	OrgCmd.AddCommand(organization.PostNetworks)
	OrgCmd.AddCommand(organization.PostSamlRole)
}

var NetCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net"},
	Short:   "Entrypoint for general network subcommands.",
}

func init() {
	NetCmd.AddCommand(network.PostBindNetwork)
	NetCmd.AddCommand(network.PostClaimSerials)
	NetCmd.AddCommand(network.PostFloorplan)
	NetCmd.AddCommand(network.PostGroupPolicy)
	NetCmd.AddCommand(network.PostHTTPServer)
	NetCmd.AddCommand(network.PostMerakiAuthUser)
	NetCmd.AddCommand(network.PostMQTTBroker)
	NetCmd.AddCommand(network.PostPIIRequest)
	NetCmd.AddCommand(network.PostProvisionClient)
	NetCmd.AddCommand(network.PostRemoveSerials)
	NetCmd.AddCommand(network.PostSplitNetwork)
	NetCmd.AddCommand(network.PostUnBindNetwork)
	NetCmd.AddCommand(network.PostWebhookTest)
}

var DeviceCmd = &cobra.Command{
	Use:   "device",
	Aliases: []string{"dev"},
	Short: "Root for general device CLI commands.",
	Long:  "Entrypoint for Device subcommands.",
}

func init() {
	DeviceCmd.AddCommand(device.PostBlinkLEDs)
	DeviceCmd.AddCommand(device.PostReboot)
}

var MXCmd = &cobra.Command{
	Use:   "appliance",
	Aliases: []string{"mx"},
	Short: "Root for all mx appliance CLI commands.",
	Long:  "Entrypoint for mx subcommands.",
}

func init() {
	MXCmd.AddCommand()
}

var MSCmd = &cobra.Command{
	Use:   "switch",
	Aliases: []string{"ms"},
	Short: "Root for all Switch CLI commands.",
}

func init() {
	MSCmd.AddCommand()
}

var MRCmd = &cobra.Command{
	Use:   "wireless",
	Aliases: []string{"mr"},
	Short: "Root for all wireless MR CLI commands.",
}

func init() {
	MRCmd.AddCommand()
}

var MGCmd = &cobra.Command{
	Use:   "gateway",
	Aliases: []string{"mg"},
	Short: "Entrypoint for mg cellular gateway subcommands.",
}

func init() {
	MGCmd.AddCommand()
}


var MVCmd = &cobra.Command{
	Use:   "camera",
	Aliases: []string{"mv"},
	Short:  "Entrypoint for mv camera subcommands.",
}

func init() {
	MVCmd.AddCommand()
}

var SMCmd = &cobra.Command{
	Use:   "systems",
	Aliases: []string{"sm"},
	Short: "Entrypoint for SM subcommands.",
}

func init() {
	SMCmd.AddCommand()
}

var InsightCmd = &cobra.Command{
	Use:   "insight",
	Aliases: []string{"in"},
	Short: "Entrypoint for Insight subcommands.",
}

func init() {
	InsightCmd.AddCommand(insight.PostMonitoredMediaServer)

}