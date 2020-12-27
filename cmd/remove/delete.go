package remove

import (
	"github.com/spf13/cobra"
)

// DeleteCmd
var DeleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Long:  "Base for all HTTP DEL Calls to Meraki Dashboard.",
}

func init() {
	DeleteCmd.AddCommand(OrgCmd)
	DeleteCmd.AddCommand(NetCmd)
	DeleteCmd.AddCommand(DeviceCmd)
	DeleteCmd.AddCommand(MXCmd)
	DeleteCmd.AddCommand(MSCmd)
	DeleteCmd.AddCommand(MRCmd)
	DeleteCmd.AddCommand(MGCmd)
	DeleteCmd.AddCommand(MVCmd)
	DeleteCmd.AddCommand(SMCmd)
	DeleteCmd.AddCommand(InsightCmd)
}
