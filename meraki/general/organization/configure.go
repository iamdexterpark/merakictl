package organization

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)



var GetActionBatches = &cobra.Command{
	Use:   "actionBatches",
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
	Use:   "actionBatch",
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

var DelActionBatch = &cobra.Command{
	Use:   "actionBatch",
	Short: "Delete Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		actionBatchId := args[0]
		metadata := configure.DelActionBatch(org, actionBatchId)
		shell.Display(metadata, "ActionBatch", cmd.Flags())
	},
}

var PutActionBatch = &cobra.Command{
	Use:   "actionBatch",
	Short: "Update The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		var format configure.ActionBatch
		input, _ := shell.ReadConfigFile(cmd, &format)
		actionBatchId := args[0]
		metadata := configure.PutActionBatch(org, actionBatchId, input)
		shell.Display(metadata, "ActionBatch", cmd.Flags())
	},
}

var PostActionBatch = &cobra.Command{
	Use:   "actionBatch",
	Short: "Create an Action Batch.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.ActionBatch
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostActionBatch(org, input)
		shell.Display(metadata, "ActionBatch", cmd.Flags())
	},
}

var GetAdmins = &cobra.Command{
	Use:   "admins",
	Short: "Display a list of dashboard administrators.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetAdmins(org)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var DelAdmins = &cobra.Command{
	Use:   "admins",
	Short: "Delete a dashboard administrator.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		adminId := args[0]
		metadata := configure.DelAdmin(org, adminId)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var PostAdmins = &cobra.Command{
	Use:   "admins",
	Short: "Create a new dashboard administrator.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		adminId := args[0]
		metadata := configure.PostAdmin(org, adminId)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var PutAdmins = &cobra.Command{
	Use:   "admins",
	Short: "Create a new dashboard administrator.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		adminId := args[0]
		var format configure.Admin
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutAdmin(org, adminId, input)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var GetBrandingPolicies = &cobra.Command{
	Use:   "brandingPolicies",
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

var PutBrandingPolicyPriorities = &cobra.Command{
	Use:   "brandingPolicyPriorities",
	Short: "Update the priority ordering of an organization's branding policies.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.BrandingPolicyPriorities
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutBrandingPolicyPriorities(org, input)
		shell.Display(metadata, "BrandingPolicyPriorities", cmd.Flags())

	},
}

var GetBrandingPolicy = &cobra.Command{
	Use:   "brandingPolicy",
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

var DelBrandingPolicy = &cobra.Command{
	Use:   "brandingPolicy",
	Short: "Delete The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		brandingId := args[0]
		metadata := configure.DelBrandingPolicy(org, brandingId)
		shell.Display(metadata, "BrandingPolicy", cmd.Flags())
	},
}

var PutBrandingPolicy = &cobra.Command{
	Use:   "brandingPolicy",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		brandingId := args[0]
		var format configure.BrandingPolicy
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutBrandingPolicy(org, brandingId, input)
		shell.Display(metadata, "BrandingPolicy", cmd.Flags())
	},
}

var PostBrandingPolicy = &cobra.Command{
	Use:   "brandingPolicy",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.BrandingPolicy
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostBrandingPolicy(org, input)
		shell.Display(metadata, "BrandingPolicy", cmd.Flags())
	},
}

var GetConfigurationTemplates = &cobra.Command{
Use:   "configurationTemplates",
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
Use:   "configurationTemplate",
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

var DelConfigurationTemplate = &cobra.Command{
	Use:   "configurationTemplate",
	Short: "Delete a Configuration Template For This Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		configTemplateId := args[0]
		metadata := configure.DelConfigTemplate(org, configTemplateId)
		shell.Display(metadata, "ConfigurationTemplate", cmd.Flags())
	},
}

var PutConfigurationTemplate = &cobra.Command{
	Use:   "configurationTemplate",
	Short: "Update a Configuration Template For This Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		configTemplateId := args[0]
		var format configure.ConfigTemplate
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutConfigTemplate(org, configTemplateId, input)
		shell.Display(metadata, "ConfigurationTemplate", cmd.Flags())
	},
}

var PostConfigurationTemplate = &cobra.Command{
	Use:   "configurationTemplate",
	Short: "Create a Configuration Template For This Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.ConfigTemplate
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostConfigTemplate(org, input)
		shell.Display(metadata, "ConfigurationTemplate", cmd.Flags())
	},
}

var GetDevices = &cobra.Command{
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
		shell.Display(metadata, "Devices", cmd.Flags())
	},
}


var GetInventoryDevices = &cobra.Command{
	Use:   "inventoryDevices",
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
	Use:   "inventoryDevice",
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
			shell.Display(metadata, "Licences", cmd.Flags())
	},
}

var PostAssignSeats = &cobra.Command{
	Use:   "assignSeats",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, serial := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		var format configure.AssignSeats
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostAssignSeats(org, serial, input)
		shell.Display(metadata, "AssignSeats", cmd.Flags())
	},
}

var PostMoveSeats = &cobra.Command{
	Use:   "moveSeats",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, serial := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		var format configure.MoveSeats
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostMoveSeats(org, serial, input)
		shell.Display(metadata, "MoveSeats", cmd.Flags())
	},
}

var PostMoveLicenses = &cobra.Command{
	Use:   "moveLicenses",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, serial := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		var format configure.MoveLicenses
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostMoveLicenses(org, serial, input)
		shell.Display(metadata, "MoveLicenses", cmd.Flags())
	},
}

var PostRenewSeats = &cobra.Command{
	Use:   "renewSeats",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, serial := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		var format configure.RenewSeats
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostRenewSeats(org, serial, input)
		shell.Display(metadata, "RenewSeats", cmd.Flags())
	},
}

var GetLicence = &cobra.Command{
	Use:   "licence",
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

var PutLicence = &cobra.Command{
	Use:   "licence",
	Short: "Update a license.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		licenceId, _ := cmd.Flags().GetString("licenceId")
		metadata := configure.PutLicense(org, licenceId)
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

var PutLoginSecurity = &cobra.Command{
	Use:   "loginSecurity",
	Short: "Update The Login Security Settings For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.LoginSecurity
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutLoginSecurity(org, input)
		shell.Display(metadata, "LoginSecurity", cmd.Flags())
	},
}

var GetNetworks = &cobra.Command{
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

		shell.Display(metadata, "Networks", cmd.Flags())

	},
}


var PostNetworks = &cobra.Command{
	Use:   "networks",
	Short: "Create a network that the user has privileges on in an organization.",
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.Networks
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostNetworks(org, input)

		shell.Display(metadata, "Networks", cmd.Flags())

	},
}

var PostCombineNetworks = &cobra.Command{
	Use:   "combineNetworks",
	Short: "Combine multiple networks into a single network.",
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.CombineNetworks
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostCombineNetworks(org, input)

		shell.Display(metadata, "CombineNetworks", cmd.Flags())

	},
}

var GetIDPS = &cobra.Command{
	Use:   "idps",
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
	Use:   "idp",
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

var DelIDP = &cobra.Command{
	Use:   "idp",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		ldp := args[0]
		metadata := configure.DelIDP(org, ldp)
		shell.Display(metadata, "IDP", cmd.Flags())
	},
}

var PutIDP = &cobra.Command{
	Use:   "idp",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		ldp := args[0]
		var format configure.IDP
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutIDP(org, ldp, input)
		shell.Display(metadata, "IDP", cmd.Flags())
	},
}

var PostIDP = &cobra.Command{
	Use:   "idp",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.IDP
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostIDP(org, input)
		shell.Display(metadata, "IDP", cmd.Flags())
	},
}

var GetSAML = &cobra.Command{
	Use:   "saml",
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

var PutSAML = &cobra.Command{
	Use:   "saml",
	Short: "Returns the SAML SSO enabled settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.SAML
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSAML(org, input)
		shell.Display(metadata, "SAML", cmd.Flags())
	},
}

var GetSamlRoles = &cobra.Command{
	Use:   "samlRoles",
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
	Use:   "samlRole",
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

var DelSamlRole = &cobra.Command{
	Use:   "samlRole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		samlRoleId := args[0]
		metadata := configure.DelSAMLRole(org, samlRoleId)
		shell.Display(metadata, "SamlRole", cmd.Flags())
	},
}

var PutSamlRole = &cobra.Command{
	Use:   "samlRole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		samlRoleId := args[0]
		var format configure.SAMLRole
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSAMLRole(org, samlRoleId, input)
		shell.Display(metadata, "SamlRole", cmd.Flags())
	},
}

var PostSamlRole = &cobra.Command{
	Use:   "samlRole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.SAMLRole
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostSAMLRole(org, input)
		shell.Display(metadata, "SamlRole", cmd.Flags())
	},
}

var GetSNMP = &cobra.Command{
	Use:   "snmp",
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


var PutSNMP = &cobra.Command{
	Use:   "snmp",
	Short: "Update the SNMP settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.SNMP
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSNMP(org, input)
		shell.Display(metadata, "SNMP", cmd.Flags())
	},
}

var GetList = &cobra.Command{
	Use:   "list",
	Short: "List the organizations that the user has privileges on.",
	Run: func(cmd *cobra.Command, args []string) {
		metadata := configure.GetOrganizations()
		shell.Display(metadata, "organizations", cmd.Flags())
	},
}

var GetDetail = &cobra.Command{
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


var DelOrganization = &cobra.Command{
	Use:   "details",
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
	Use:   "details",
	Short: "Update an existing organization",
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		var name = args[0]

		metadata := configure.PutOrganization(org, name)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var PostOrganization = &cobra.Command{
	Use:   "details",
	Short: "Create an organization",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		metadata := configure.PostOrganization(name)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var PostClaim = &cobra.Command{
	Use:   "claim",
	Short: "Claim a list of devices, licenses, and/or orders into an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		var format configure.Claim
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostClaim(org, input)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var PostClone = &cobra.Command{
	Use:   "clone",
	Short: "Create a new organization by cloning the addressed organization",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		name := args[0]
		var format configure.Clone
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostClone(org, name, input)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}