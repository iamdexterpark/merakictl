package organization

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)


var GetList = &cobra.Command{
	Use:   "List",
	Short: "List the organizations that the user has privileges on.",
	Run: func(cmd *cobra.Command, args []string) {
		metadata := configure.GetOrganizations()
		shell.Display(metadata, "organizations", cmd.Flags())
	},
}

var GetDetail = &cobra.Command{
	Use:   "Details",
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


var DelOrganization = &cobra.Command{
	Use:   "fromList",
	Short: "Delete an existing organization",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.DelOrganization(org)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var PutOrganization = &cobra.Command{
	Use:   "org",
	Short: "Update an existing organization",
	Long:  `update org {NAME} -o {ORG_ID}`,
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var name = args[0]

		metadata := configure.PutOrganization(org, name)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var GetActionBatches = &cobra.Command{
	Use:   "ActionBatches",
	Short: "Return The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		status := cmd.Flag("status").Value.String()
		metadata := configure.GetActionBatches(org, status)
		shell.Display(metadata, "ActionBatches", cmd.Flags())

	},
}

var GetActionBatch = &cobra.Command{
	Use:   "ActionBatch",
	Short: "Return The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		actionBatchId := args[0]
		metadata := configure.GetActionBatch(org, actionBatchId)
		shell.Display(metadata, "ActionBatch", cmd.Flags())
	},
}

var GetAdmins = &cobra.Command{
	Use:   "Admins",
	Short: "Return The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetAdmins(org)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var GetBrandingPolicies = &cobra.Command{
	Use:   "BrandingPolicies",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetBrandingPolicies(org)
		shell.Display(metadata, "BrandingPolicies", cmd.Flags())

	},
}

var GetBrandingPolicy = &cobra.Command{
	Use:   "BrandingPolicy",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		brandingId := args[0]
		metadata := configure.GetBrandingPolicy(org, brandingId)
		shell.Display(metadata, "BrandingPolicy", cmd.Flags())
	},
}


var GetConfigurationTemplates = &cobra.Command{
Use:   "ConfigurationTemplates",
Short: "List The Configuration Templates For This Organization.",
Run: func(cmd *cobra.Command, args []string) {
org, _, _ := shell.ResolveFlags(cmd.Flags())
if org == "" {
org = args[0]
}
metadata := configure.GetConfigTemplates(org)
shell.Display(metadata, "ConfigurationTemplates", cmd.Flags())
},
}

var GetConfigurationTemplate = &cobra.Command{
Use:   "ConfigurationTemplate",
Short: "Return a Configuration Template For This Organization.",
Run: func(cmd *cobra.Command, args []string) {
org, _, _ := shell.ResolveFlags(cmd.Flags())
if org == "" {
org = args[1]
}
configTemplateId := args[0]
metadata := configure.GetConfigTemplate(org, configTemplateId)
shell.Display(metadata, "ConfigurationTemplate", cmd.Flags())
},
}


var GetDevices = &cobra.Command{
	Use:   "Devices",
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
		shell.Display(metadata, "Devices", cmd.Flags())
	},
}


var GetInventoryDevices = &cobra.Command{
	Use:   "InventoryDevices",
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
		shell.Display(metadata, "InventoryDevices", cmd.Flags())
	},
}

var GetInventoryDevice = &cobra.Command{
	Use:   "InventoryDevice",
	Short: "Return The Device Inventory For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		deviceId := args[0]

		metadata := configure.GetInventoryDevice(org, deviceId)
		shell.Display(metadata, "InventoryDevice", cmd.Flags())
	},
}

var GetLicences = &cobra.Command{
	Use:   "Licences",
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
			shell.Display(metadata, "Licences", cmd.Flags())
	},
}

var GetLicence = &cobra.Command{
	Use:   "Licence",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		licenceId, _ := cmd.Flags().GetString("licenceId")
		metadata := configure.GetLicense(org, licenceId)
		shell.Display(metadata, "Licence", cmd.Flags())
	},
}

var GetLoginSecurity = &cobra.Command{
	Use:   "loginSecurity",
	Short: "Returns The Login Security Settings For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetLoginSecurity(org)
		shell.Display(metadata, "LoginSecurity", cmd.Flags())
	},
}


var GetNetworks = &cobra.Command{
	Use:   "Networks",
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

		shell.Display(metadata, "Networks", cmd.Flags())

	},
}

var GetIDPS = &cobra.Command{
	Use:   "IDPS",
	Short: "List the SAML IdPs in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetIDPS(org)
		shell.Display(metadata, "IDPS", cmd.Flags())
	},
}

var GetIDP = &cobra.Command{
	Use:   "IDP",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		ldp := args[0]
		metadata := configure.GetIDP(org, ldp)
		shell.Display(metadata, "IDP", cmd.Flags())
	},
}

var GetSAML = &cobra.Command{
	Use:   "SAML",
	Short: "Returns the SAML SSO enabled settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetSAML(org)
		shell.Display(metadata, "SAML", cmd.Flags())
	},
}


var GetSamlRoles = &cobra.Command{
	Use:   "SamlRoles",
	Short: "List the SAML roles for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetSAMLRoles(org)
		shell.Display(metadata, "SamlRoles", cmd.Flags())
	},
}

var GetSamlRole = &cobra.Command{
	Use:   "SamlRole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		samlRoleId := args[0]
		metadata := configure.GetSAMLRole(org, samlRoleId)
		shell.Display(metadata, "SamlRole", cmd.Flags())
	},
}

var GetSNMP = &cobra.Command{
	Use:   "SNMP",
	Short: "Return the SNMP settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetSNMP(org)
		shell.Display(metadata, "SNMP", cmd.Flags())
	},
}
