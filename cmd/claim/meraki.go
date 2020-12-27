package claim

import (
	"github.com/ddexterpark/merakictl/meraki/general/network"
	"github.com/ddexterpark/merakictl/meraki/general/organization"
	"github.com/spf13/cobra"
)

var OrgCmd = &cobra.Command{
	Use:   "organization",
	Aliases: []string{"org"},
	Short: "Root for all organization CLI commands.",
}

func init() {
	OrgCmd.AddCommand(organization.PostClaim)
	OrgCmd.AddCommand(organization.PostMoveLicenses)
	OrgCmd.AddCommand(organization.PostMoveSeats)
	OrgCmd.AddCommand(organization.PostAssignSeats)
	OrgCmd.AddCommand(organization.PostRenewSeats)
}

var NetCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net"},
	Short:   "Entrypoint for general network subcommands.",
}

func init() {
	NetCmd.AddCommand(network.PostRemoveSerials)
	NetCmd.AddCommand(network.PostClaimSerials)
}
