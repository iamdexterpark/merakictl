package organization

import (
	"github.com/ddexterpark/merakictl/api"
	"github.com/ddexterpark/merakictl/api/general/organizations/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)


var orgList = &cobra.Command{
	Use:   "list",
	Short: "List the organizations that the user has privileges on.",
	Long: pretty.Sprint(api.Organizations{}),
	Run: func(cmd *cobra.Command, args []string) {
		metadata := configure.GetOrganizations()
		shell.Display(metadata, "organizations", cmd.Flags())
	},
}

var single = &cobra.Command{
	Use:   "single",
	Short: "List a specific organization that the user has privileges on.",
	Long: pretty.Sprint(api.Organization{}),
	Run: func(cmd *cobra.Command, args []string) {

		org := cmd.Flag("organization").Value.String()
		metadata := configure.GetOrganization(org)

		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var networks = &cobra.Command{
	Use:   "networks",
	Short: "List the networks that the user has privileges on in an organization.",
	Long:  pretty.Sprint(api.Network{}),
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		metadata := api.GetOrganizationNetworks(org, "",
			"", "", "", "", "10")


		shell.Display(metadata, "networks", cmd.Flags())

	},
}


var devices = &cobra.Command{
	Use:   "devices",
	Short: "List the devices in an organization.",
	Long: pretty.Sprint(api.Devices{}),
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		metadata := api.GetOrganizationDevices(org, "10", "", "")
		shell.Display(metadata, "devices", cmd.Flags())
	},
}


var actionbatchlist = &cobra.Command{
	Use:   "ablist",
	Short: "Return The List Of Action Batches In The Organization.",
	Long:  pretty.Sprint(configure.ActionBatchList{}),
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		metadata := configure.GetActionBatchList(org, "")
		shell.Display(metadata, "actionbatchlist", cmd.Flags())

	},
}

var admins = &cobra.Command{
	Use:   "admins",
	Short: "Return The List Of Action Batches In The Organization.",
	Long:  pretty.Sprint(configure.Admin{}),
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		metadata := configure.GetAdmins(org)
		shell.Display(metadata, "admins", cmd.Flags())

	},
}


var brandingPolicies = &cobra.Command{
	Use:   "brandingPolicies",
	Short: "",
	Long:  pretty.Sprint(configure.BrandingPolicies{}),
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		metadata := configure.GetBrandingPolicies(org)
		shell.Display(metadata, "brandingPolicies", cmd.Flags())

	},
}


var brandingPolicy= &cobra.Command{
	Use:   "brandingPolicy",
	Short: "",
	Long:  pretty.Sprint(configure.BrandingPolicy{}),
	Run: func(cmd *cobra.Command, args []string) {
		org := cmd.Flag("organization").Value.String()
		brandingId := args[0]

		metadata := configure.GetBrandingPolicy(org,brandingId)
		shell.Display(metadata, "brandingPolicy", cmd.Flags())

	},
}

