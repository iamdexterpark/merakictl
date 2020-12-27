package get

import (
	"github.com/ddexterpark/merakictl/cmd/utilities"
	"github.com/spf13/cobra"
)

// GetCmd
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Queries EXISTING configurations.",
	Long:  "Base for all HTTP GET Calls to Meraki Dashboard.",
}

func init() {
	GetCmd.AddCommand(utilities.EnvironmentalVariables)
	GetCmd.AddCommand(utilities.Test)
	GetCmd.AddCommand(OrgCmd)
	GetCmd.AddCommand(NetCmd)
	GetCmd.AddCommand(DeviceCmd)
	GetCmd.AddCommand(MXCmd)
	GetCmd.AddCommand(MSCmd)
	GetCmd.AddCommand(MRCmd)
	GetCmd.AddCommand(MGCmd)
	GetCmd.AddCommand(MVCmd)
	GetCmd.AddCommand(SMCmd)
	GetCmd.AddCommand(InsightCmd)
	}
