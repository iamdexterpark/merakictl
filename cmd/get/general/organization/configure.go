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
		metadata := configure.GetActionBatches(org, status)
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
metadata := configure.GetConfigTemplates(org)
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
metadata := configure.GetConfigTemplate(org, configTemplateId)
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

var licences = &cobra.Command{
	Use:   "licences",
	Short: "List The Licenses For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if networkId == "" {
			networkId = args[1]
		}

		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		deviceSerial, _ := cmd.Flags().GetString("deviceSerial")
		state, _ := cmd.Flags().GetString("state")

		metadata := configure.GetLicenses(org, perPage,
			startingAfter, endingBefore, deviceSerial, networkId, state)
			shell.Display(metadata, "licences", cmd.Flags())
	},
}

var licence = &cobra.Command{
	Use:   "licence",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		licenceId, _ := cmd.Flags().GetString("licenceId")
		metadata := configure.GetLicense(org, licenceId)
		shell.Display(metadata, "licences", cmd.Flags())
	},
}

var securitysettings = &cobra.Command{
	Use:   "securitysettings",
	Short: "Returns The Login Security Settings For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetLoginSecurity(org)
		shell.Display(metadata, "securitysettings", cmd.Flags())
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

		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		configTemplateId, _ := cmd.Flags().GetString("configTemplateId")
		tagsFilterType, _ := cmd.Flags().GetString("tagsFilterType")
		tags, _ := cmd.Flags().GetString("tags")

		metadata := configure.GetNetworks(org, configTemplateId,
			tagsFilterType, startingAfter, endingBefore, tags, perPage)

		shell.Display(metadata, "networks", cmd.Flags())

	},
}

var ldps = &cobra.Command{
	Use:   "ldps",
	Short: "List the SAML IdPs in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetIDPS(org)
		shell.Display(metadata, "ldps", cmd.Flags())
	},
}

var ldp = &cobra.Command{
	Use:   "ldp",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		ldp := args[0]
		metadata := configure.GetIDP(org, ldp)
		shell.Display(metadata, "ldps", cmd.Flags())
	},
}

var saml = &cobra.Command{
	Use:   "saml",
	Short: "Returns the SAML SSO enabled settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetSAML(org)
		shell.Display(metadata, "saml", cmd.Flags())
	},
}


var samlroles = &cobra.Command{
	Use:   "samlroles",
	Short: "List the SAML roles for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetSAMLRoles(org)
		shell.Display(metadata, "samlroles", cmd.Flags())
	},
}

var samlrole = &cobra.Command{
	Use:   "samlrole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		samlRoleId := args[0]
		metadata := configure.GetSAMLRole(org, samlRoleId)
		shell.Display(metadata, "samlrole", cmd.Flags())
	},
}

var snmp = &cobra.Command{
	Use:   "snmp",
	Short: "Return the SNMP settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetSNMP(org)
		shell.Display(metadata, "snmp", cmd.Flags())
	},
}
