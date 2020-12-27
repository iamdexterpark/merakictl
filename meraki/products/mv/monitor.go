package mv

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/camera/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var GetLiveAnalytics = &cobra.Command{
	Use:   "LiveAnalytics",
	Short: "Returns live state from mv of analytics zones.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := monitor.GetAnalyticZones(serial)
		shell.Display(metadata, "LiveAnalytics", cmd.Flags())
	},
}

var GetAnalyticsOverview = &cobra.Command{
	Use:   "AnalyticsOverview",
	Short: "Returns an overview of aggregate analytics data for a timespan.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		objectType, _ := cmd.Flags().GetString("objectType")
		metadata := monitor.GetAnalyticsOverview(serial, t0, t1, timespan, objectType)
		shell.Display(metadata, "AnalyticsOverview", cmd.Flags())
	},
}

var GetRecentAnalytics = &cobra.Command{
	Use:   "RecentAnalytics",
	Short: "Returns most recent record for analytics zones.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		objectType, _ := cmd.Flags().GetString("objectType")
		metadata := monitor.GetAnalyticsRecent(serial, objectType)
		shell.Display(metadata, "RecentAnalytics", cmd.Flags())
	},
}

var GetAnalyticsZonesHistory = &cobra.Command{
	Use:   "AnalyticsZonesHistory",
	Short: "Return historical records for analytic zones.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[1]
		}

		zoneId := args[0]
		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		resolution, _ := cmd.Flags().GetString("resolution")
		objectType, _ := cmd.Flags().GetString("objectType")
		metadata := monitor.GetAnalyticsZoneHistory(serial,
			zoneId, t0, t1, timespan, resolution, objectType)
		shell.Display(metadata, "AnalyticsZonesHistory", cmd.Flags())
	},
}

var GetAnalyticsZones = &cobra.Command{
	Use:   "AnalyticsZones",
	Short: "Returns all configured analytic zones for this mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := monitor.GetAnalyticZones(serial)
		shell.Display(metadata, "AnalyticsZones", cmd.Flags())
	},
}