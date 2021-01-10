package put

import (
	"github.com/spf13/cobra"
)

// UpdateCmd
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Aliases: []string{"put"},
	Short: "Updates configurations.",
	Long:  "Base for all HTTP PUT Calls to Meraki Dashboard.",
}


func init() {
	UpdateCmd.AddCommand(OrgCmd)
	UpdateCmd.AddCommand(NetCmd)
	UpdateCmd.AddCommand(DeviceCmd)
	UpdateCmd.AddCommand(MXCmd)
	UpdateCmd.AddCommand(MSCmd)
	UpdateCmd.AddCommand(MRCmd)
	UpdateCmd.AddCommand(MGCmd)
	UpdateCmd.AddCommand(MVCmd)
	UpdateCmd.AddCommand(SMCmd)
	UpdateCmd.AddCommand(InsightCmd)
}


