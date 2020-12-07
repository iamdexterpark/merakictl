package camera

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/camera/monitor"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var liveanalytics = &cobra.Command{
	Use:   "liveanalytics",
	Short: "Returns live state from camera of analytics zones.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := monitor.GetLiveState(serial)
		shell.Display(metadata, "liveanalytics", cmd.Flags())
	},
}

var analyticsoverview = &cobra.Command{
	Use:   "analyticsoverview",
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
		metadata := monitor.GetAggregateAnalytics(serial, t0, t1, timespan, objectType)
		shell.Display(metadata, "analyticsoverview", cmd.Flags())
	},
}

var recentanalytics = &cobra.Command{
	Use:   "recentanalytics",
	Short: "Returns most recent record for analytics zones.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		objectType, _ := cmd.Flags().GetString("objectType")
		metadata := monitor.GetAnalyticsZoneRecord(serial, objectType)
		shell.Display(metadata, "recentanalytics", cmd.Flags())
	},
}

var analyticszoneshistory = &cobra.Command{
	Use:   "analyticszoneshistory",
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
		metadata := monitor.GetAnalyticsZoneHistoricalRecords(serial,
			zoneId, t0, t1, timespan, resolution, objectType)
		shell.Display(metadata, "analyticszoneshistory", cmd.Flags())
	},
}

var analyticszones = &cobra.Command{
	Use:   "analyticszones",
	Short: "Returns all configured analytic zones for this camera.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := monitor.GetAnalyticZones(serial)
		shell.Display(metadata, "analyticszones", cmd.Flags())
	},
}