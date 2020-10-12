package organization

import (
	"github.com/ddexterpark/merakictl/api/general/organizations/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)


// orgList
var orgList = &cobra.Command{
	Use:   "list",
	Short: "List the organizations that the user has privileges on.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		organizations, traceback := configure.GetOrganizations()
		shell.Display(organizations, traceback, "organizations", cmd.Flags())
	},
}

// org
var single = &cobra.Command{
	Use:   "single",
	Short: "List a specific organization that the user has privileges on.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		org := cmd.Flag("organization").Value.String()
		organization, traceback := configure.GetOrganization(org)

		shell.Display(organization, traceback, "organization", cmd.Flags())
	},
}

// networks
var networks = &cobra.Command{
	Use:   "networks",
	Short: "List the networks that the user has privileges on in an organization.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		networks, traceback := configure.GetOrganizationNetworks(org, "",
			"", "", "", "", "1000")
		shell.Display(networks, traceback, "networks", cmd.Flags())

	},
}


// devices
var devices = &cobra.Command{
	Use:   "devices",
	Short: "List the devices in an organization.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		devicestatus, traceback := configure.GetOrganizationDevices(org)
		shell.Display(devicestatus, traceback, "devices", cmd.Flags())
	},
}

// action batch list
var actionbatchlist = &cobra.Command{
	Use:   "ablist",
	Short: "Return The List Of Action Batches In The Organization.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		organization, traceback := configure.GetActionBatchList(org)
		shell.Display(organization, traceback, "actionbatchlist", cmd.Flags())

	},
}

// admins
var admins = &cobra.Command{
	Use:   "admins",
	Short: "Return The List Of Action Batches In The Organization.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		organization, traceback := configure.GetAdmins(org)
		shell.Display(organization, traceback, "admins", cmd.Flags())

	},
}

// brandingPolicies
var brandingPolicies = &cobra.Command{
	Use:   "brandingPolicies",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		organization, traceback := configure.GetBrandingPolicies(org)
		shell.Display(organization, traceback, "brandingPolicies", cmd.Flags())

	},
}

// brandingPolicies
var brandingPolicy= &cobra.Command{
	Use:   "brandingPolicy",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		brandingId := args[0]

		organization, traceback := configure.GetBrandingPolicy(org,brandingId)
		shell.Display(organization, traceback, "brandingPolicy", cmd.Flags())

	},
}

