package organization

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/organizations/monitor"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

var GetApiRequestsOverview = &cobra.Command{
	Use:   "apiRequestsOverview",
	Short: "Return an aggregated overview of API requests data.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")


		metadata := monitor.GetAPIRequestSummary(org, t0, t1, timespan)
		shell.Display(metadata, "ApiRequestsOverview", cmd.Flags())
	},
}

var GetApiRequests = &cobra.Command{
	Use:   "apiRequests",
	Short: "List the API requests made by an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		adminId, _ := cmd.Flags().GetString("adminId")
		path, _ := cmd.Flags().GetString("path")
		method, _ := cmd.Flags().GetString("method")
		responseCode, _ := cmd.Flags().GetString("responseCode")
		sourceIp, _ := cmd.Flags().GetString("sourceIp")

		metadata := monitor.GetAPIRequests(org, t0, t1, timespan,
			perPage, startingAfter, endingBefore, adminId, path,
			method, responseCode, sourceIp)
		shell.Display(metadata, "ApiRequests", cmd.Flags())
	},
}

var GetConfigurationChanges = &cobra.Command{
	Use:   "configurationChanges",
	Short: "View the Change Log for your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if networkId == "" {
			networkId = args[1]
		}

		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		adminId, _ := cmd.Flags().GetString("adminId")

		metadata := monitor.GetConfigurationChanges(org, t0, t1, timespan,
			perPage, startingAfter, endingBefore, adminId, networkId)
		shell.Display(metadata, "ConfigurationChanges", cmd.Flags())
	},
}


var GetDeviceStatuses = &cobra.Command{
	Use:   "deviceStatuses",
	Short: "List the status of every Meraki device in the organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")


		metadata := monitor.GetDeviceStatus(org, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "DeviceStatuses", cmd.Flags())
	},
}


var GetUplinksLossAndLatency = &cobra.Command{
	Use:   "uplinksLossAndLatency",
	Short: "Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		uplink, _ := cmd.Flags().GetString("uplink")
		ip, _ := cmd.Flags().GetString("ip")


		metadata := monitor.GetUplinksLossAndLatency(org, t0, t1, timespan, uplink, ip)
		shell.Display(metadata, "UplinksLossAndLatency", cmd.Flags())
	},
}


var GetLicenseOverview = &cobra.Command{
	Use:   "licenseOverview",
	Short: "Return an overview of the license state for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := monitor.GetLicenseOverview(org)
		shell.Display(metadata, "LicenseOverview", cmd.Flags())
	},
}

var GetOpenAPISpec = &cobra.Command{
	Use:   "openAPISpec",
	Short: "Return the OpenAPI 2.0 Specification of the organization's API documentation in JSON.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := monitor.GetOpenapiSpec(org)
		shell.Display(metadata, "OpenAPISpec", cmd.Flags())
	},
}

var GetAlertTypes = &cobra.Command{
	Use:   "alertTypes",
	Short: "Return a list of alert types to be used with managing webhook alerts.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := monitor.GetAlertTypes(org)
		shell.Display(metadata, "AlertTypes", cmd.Flags())
	},
}

var GetWebhookLogs = &cobra.Command{
	Use:   "webhookLogs",
	Short: "Return the log of webhook POSTs sent.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		url, _ := cmd.Flags().GetString("url")

		metadata := monitor.GetWebHookLogs(org, t0, t1, timespan,
			perPage, startingAfter, endingBefore, url)
		shell.Display(metadata, "WebhookLogs", cmd.Flags())
	},
}