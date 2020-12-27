package sm

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/sm/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var GetCellularUsageHistory = &cobra.Command{
	Use:   "CellularUsageHistory",
	Short: "Return the client's daily cellular data usage history. Usage data is in kilobytes.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		deviceId := args[0]
		metadata := monitor.GetCellularUsageHistory(networkId, deviceId)
		shell.Display(metadata, "CellularUsageHistory", cmd.Flags())
	},
}


var GetConnectivity = &cobra.Command{
Use:   "Connectivity",
Short: "",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	deviceId := args[0]
	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := monitor.GetConnectivity(networkId, deviceId,
	perPage, startingAfter, endingBefore)
shell.Display(metadata, "Connectivity", cmd.Flags())
},
}


var GetDesktopLogs = &cobra.Command{
Use:   "DesktopLogs",
Short: "Return historical records of various Systems Manager network connection details for desktop devices.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	deviceId := args[0]
	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := monitor.GetDesktopLogs(networkId, deviceId, perPage, startingAfter, endingBefore)
shell.Display(metadata, "DesktopLogs", cmd.Flags())
},
}


var GetDeviceCommandLogs = &cobra.Command{
Use:   "DeviceCommandLogs",
Short: "Return historical records of commands sent to Systems Manager devices. ",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}

	deviceId := args[0]
	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := monitor.GetPerformanceHistory(networkId,
	deviceId, perPage, startingAfter, endingBefore)
shell.Display(metadata, "DeviceCommandLogs", cmd.Flags())
},
}


var GetPerformanceHistory = &cobra.Command{
Use:   "PerformanceHistory",
Short: "Return historical records of various Systems Manager client metrics for desktop devices.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
networkId = args[1]
}
	deviceId := args[0]
	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := monitor.GetPerformanceHistory(networkId,
	deviceId, perPage, startingAfter, endingBefore)
shell.Display(metadata, "PerformanceHistory", cmd.Flags())
},
}

