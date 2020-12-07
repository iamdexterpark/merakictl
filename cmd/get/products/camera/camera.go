package camera

import "github.com/spf13/cobra"

//MVCmd - Root for all organization CLI commands.
var MVCmd = &cobra.Command{
	Use:   "mv",
	Short: "",
	Long:  "Entrypoint for camera subcommands.",
}

// init - Entrypoint for DeviceCmd sub-commands.
func init() {
	MVCmd.AddCommand(qualityandretention)
	MVCmd.AddCommand(qualityretentionprofiles)
	MVCmd.AddCommand(qualityretentionprofile)
	MVCmd.AddCommand(schedules)
	MVCmd.AddCommand(objectdetectionmodels)
	MVCmd.AddCommand(sense)
	MVCmd.AddCommand(videosettings)
	MVCmd.AddCommand(videolink)
	MVCmd.AddCommand(liveanalytics)
	MVCmd.AddCommand(analyticsoverview)
	MVCmd.AddCommand(recentanalytics)
	MVCmd.AddCommand(analyticszoneshistory)
	MVCmd.AddCommand(analyticszones)
}