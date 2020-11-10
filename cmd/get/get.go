package get

import (
	"github.com/ddexterpark/merakictl/cmd/get/device"
	"github.com/ddexterpark/merakictl/cmd/get/network"
	"github.com/ddexterpark/merakictl/cmd/get/organization"
	"github.com/ddexterpark/merakictl/cmd/get/wireless"
	"github.com/spf13/cobra"
)

// GetCmd
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Queries EXISTING configurations.",
	Long:  "Base for all HTTP GET Calls to Meraki Dashboard.",
}

func init() {
	GetCmd.AddCommand(organization.OrgCmd)
	GetCmd.AddCommand(device.DeviceCmd)
	GetCmd.AddCommand(network.NetCmd)
	GetCmd.AddCommand(wireless.MRCmd)
	GetCmd.AddCommand(apikey)
}
