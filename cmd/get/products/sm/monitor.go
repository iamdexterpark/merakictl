package sm

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/sm/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var cellularusagehistory = &cobra.Command{
	Use:   "cellularusagehistory",
	Short: "Return the client's daily cellular data usage history. Usage data is in kilobytes.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		deviceId := args[0]
		metadata := monitor.GetClientCellularData(networkId, deviceId)
		shell.Display(metadata, "cellularusagehistory", cmd.Flags())
	},
}


var connectivity = &cobra.Command{
Use:   "connectivity",
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

metadata := monitor.GetHistoricalConnectivityData(networkId, deviceId,
	perPage, startingAfter, endingBefore)
shell.Display(metadata, "connectivity", cmd.Flags())
},
}


var desktoplogs = &cobra.Command{
Use:   "desktoplogs",
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

metadata := monitor.GetDesktopDevicesHistoricalRecords(networkId, deviceId, perPage, startingAfter, endingBefore)
shell.Display(metadata, "desktoplogs", cmd.Flags())
},
}


var devicecommandlogs = &cobra.Command{
Use:   "devicecommandlogs",
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

metadata := monitor.GetCommandHistoricalRecords(networkId,
	deviceId, perPage, startingAfter, endingBefore)
shell.Display(metadata, "devicecommandlogs", cmd.Flags())
},
}


var performancehistory = &cobra.Command{
Use:   "performancehistory",
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

metadata := monitor.GetClientMetricsHistoricalRecords(networkId,
	deviceId, perPage, startingAfter, endingBefore)
shell.Display(metadata, "performancehistory", cmd.Flags())
},
}

