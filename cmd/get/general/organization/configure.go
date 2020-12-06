package organization

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var list = &cobra.Command{
	Use:   "list",
	Short: "List the organizations that the user has privileges on.",
	Run: func(cmd *cobra.Command, args []string) {
		metadata := configure.GetOrganizations()
		shell.Display(metadata, "organizations", cmd.Flags())
	},
}

var detail = &cobra.Command{
	Use:   "details",
	Short: "List a specific organization that the user has privileges on.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetOrganization(org)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}


var actionbatches = &cobra.Command{
	Use:   "actionbatches",
	Short: "Return The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		status := cmd.Flag("status").Value.String()
		metadata := configure.GetActionBatchList(org, status)
		shell.Display(metadata, "actionbatchlist", cmd.Flags())

	},
}

var actionbatch = &cobra.Command{
	Use:   "actionbatch",
	Short: "Return The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		actionBatchId := args[0]
		metadata := configure.GetActionBatch(org, actionBatchId)
		shell.Display(metadata, "actionbatchlist", cmd.Flags())
	},
}

var admins = &cobra.Command{
	Use:   "admins",
	Short: "Return The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetAdmins(org)
		shell.Display(metadata, "admins", cmd.Flags())

	},
}

var brandingPolicies = &cobra.Command{
	Use:   "brandingPolicies",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetBrandingPolicies(org)
		shell.Display(metadata, "brandingPolicies", cmd.Flags())

	},
}

var brandingPolicy = &cobra.Command{
	Use:   "brandingPolicy",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		brandingId := args[0]
		metadata := configure.GetBrandingPolicy(org, brandingId)
		shell.Display(metadata, "brandingPolicy", cmd.Flags())
	},
}

var configurationtemplates = &cobra.Command{
Use:   "configurationtemplates",
Short: "List The Configuration Templates For This Organization.",
Run: func(cmd *cobra.Command, args []string) {
org, _, _ := shell.ResolveFlags(cmd.Flags())
if org == "" {
org = args[0]
}
metadata := configure.GetConfigurationTemplates(org)
shell.Display(metadata, "configurationtemplates", cmd.Flags())
},
}

var configurationtemplate = &cobra.Command{
Use:   "configurationtemplate",
Short: "Return a Configuration Template For This Organization.",
Run: func(cmd *cobra.Command, args []string) {
org, _, _ := shell.ResolveFlags(cmd.Flags())
if org == "" {
org = args[1]
}
configTemplateId := args[0]
metadata := configure.GetConfigurationTemplate(org, configTemplateId)
shell.Display(metadata, "configurationtemplate", cmd.Flags())
},
}


var devices = &cobra.Command{
	Use:   "devices",
	Short: "List the devices in an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		configurationUpdatedAfter, _ := cmd.Flags().GetString("configurationUpdatedAfter")
		metadata := configure.GetOrganizationDevices(org, perPage, startingAfter, configurationUpdatedAfter)
		shell.Display(metadata, "devices", cmd.Flags())
	},
}


var inventory = &cobra.Command{
	Use:   "inventory",
	Short: "Return The Device Inventory For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		usedState, _ := cmd.Flags().GetString("usedState")
		search, _ := cmd.Flags().GetString("search")

		metadata := configure.GetInventoryDevices(org, perPage, startingAfter, endingBefore, usedState, search)
		shell.Display(metadata, "devices", cmd.Flags())
	},
}



var networks = &cobra.Command{
	Use:   "networks",
	Short: "List the networks that the user has privileges on in an organization.",
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		perPage := cmd.Flag("perPage").Value.String()
		configTemplateId := cmd.Flag("configTemplateId").Value.String()
		tagsFilterType :=  cmd.Flag("tagsFilterType").Value.String()
		startingAfter := cmd.Flag("startingAfter").Value.String()
		endingBefore :=  cmd.Flag("endingBefore").Value.String()
		tags := cmd.Flag("tags").Value.String()

		metadata := configure.GetOrganizationNetworks(org, configTemplateId,
			tagsFilterType, startingAfter, endingBefore, tags, perPage)

		shell.Display(metadata, "networks", cmd.Flags())

	},
}