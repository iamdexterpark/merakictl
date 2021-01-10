package get

import (
	"github.com/ddexterpark/merakictl/cmd/misc"
	"github.com/spf13/cobra"
)

// ShowCmd
var ShowCmd = &cobra.Command{
	Use:   "show",
	Aliases: []string{"get"},
	Short: "Queries EXISTING configurations.",
	Long:  "Base for all HTTP GET Calls to Meraki Dashboard.",
}

func init() {
	ShowCmd.AddCommand(misc.EnvironmentalVariables)
	ShowCmd.AddCommand(misc.Test)
	ShowCmd.AddCommand(OrgCmd)
	ShowCmd.AddCommand(NetCmd)
	ShowCmd.AddCommand(DeviceCmd)
	ShowCmd.AddCommand(MXCmd)
	ShowCmd.AddCommand(MSCmd)
	ShowCmd.AddCommand(MRCmd)
	ShowCmd.AddCommand(MGCmd)
	ShowCmd.AddCommand(MVCmd)
	ShowCmd.AddCommand(SMCmd)
	ShowCmd.AddCommand(InsightCmd)
	}
