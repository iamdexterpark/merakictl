package get

import (
	"github.com/ddexterpark/merakictl/cmd/get/general/device"
	"github.com/ddexterpark/merakictl/cmd/get/general/network"
	"github.com/ddexterpark/merakictl/cmd/get/general/organization"
	"github.com/ddexterpark/merakictl/cmd/get/products/wireless"
	"github.com/ddexterpark/merakictl/cmd/get/products/insight"
	"github.com/spf13/cobra"
)

// GetCmd
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Queries EXISTING configurations.",
	Long:  "Base for all HTTP GET Calls to Meraki Dashboard.",
}

func init() {
	GetCmd.AddCommand(test)
	GetCmd.AddCommand(organization.OrgCmd)
	GetCmd.AddCommand(device.DeviceCmd)
	GetCmd.AddCommand(network.NetCmd)
	GetCmd.AddCommand(insight.InsightCmd)
	GetCmd.AddCommand(wireless.MRCmd)
	GetCmd.AddCommand(environmentalvariables)
	}
