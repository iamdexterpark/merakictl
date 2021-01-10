package post

import (
	"github.com/spf13/cobra"
)

// CreateCmd
var CreateCmd = &cobra.Command{
	Use:   "create",
	Aliases: []string{"post", "set"},
	Short: "",
	Long:  "Base for all HTTP POST Calls to Meraki Dashboard.",
}

func init() {
	CreateCmd.AddCommand(OrgCmd)
	CreateCmd.AddCommand(NetCmd)
	CreateCmd.AddCommand(DeviceCmd)
	CreateCmd.AddCommand(MXCmd)
	CreateCmd.AddCommand(MSCmd)
	CreateCmd.AddCommand(MRCmd)
	CreateCmd.AddCommand(MGCmd)
	CreateCmd.AddCommand(MVCmd)
	CreateCmd.AddCommand(SMCmd)
	CreateCmd.AddCommand(InsightCmd)
}
