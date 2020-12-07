package camera

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/camera/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var qualityandretention = &cobra.Command{
	Use:   "qualityandretention",
	Short: "Returns quality and retention settings for the given camera.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetQualityAndRetention(serial)
		shell.Display(metadata, "qualityandretention", cmd.Flags())
	},
}

var qualityretentionprofiles = &cobra.Command{
	Use:   "qualityretentionprofiles",
	Short: "List the quality retention profiles for this network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetQualityRetentionProfiles(networkId)
		shell.Display(metadata, "qualityretentionprofiles", cmd.Flags())
	},
}

var qualityretentionprofile = &cobra.Command{
	Use:   "qualityretentionprofile",
	Short: "Retrieve a single quality retention profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		qualityRetentionProfileId:= args[0]

		metadata := configure.GetQualityRetentionProfile(networkId,
			qualityRetentionProfileId)
		shell.Display(metadata, "qualityretentionprofile", cmd.Flags())
	},
}

var schedules = &cobra.Command{
	Use:   "schedules",
	Short: "Returns a list of all camera recording schedules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetCameraRecordingSchedules(networkId)
		shell.Display(metadata, "schedules", cmd.Flags())
	},
}

var objectdetectionmodels = &cobra.Command{
	Use:   "objectdetectionmodels",
	Short: "Returns the MV Sense object detection model list for the given camera.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetObjectDetectionModel(serial)
		shell.Display(metadata, "objectdetectionmodels", cmd.Flags())
	},
}

var sense = &cobra.Command{
	Use:   "sense",
	Short: "Returns sense settings for a given camera.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetSenseSettings(serial)
		shell.Display(metadata, "sense", cmd.Flags())
	},
}


var videosettings = &cobra.Command{
	Use:   "videosettings",
	Short: "Returns video settings for the given camera.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetVideoSettings(serial)
		shell.Display(metadata, "settings", cmd.Flags())
	},
}


var videolink = &cobra.Command{
	Use:   "videolink",
	Short: "Returns video link to the specified camera. If a timestamp is supplied, it links to that timestamp.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		timestamp, _ := cmd.Flags().GetString("timestamp")
		metadata := configure.GetVideoLink(serial,timestamp)
		shell.Display(metadata, "videolink", cmd.Flags())
	},
}

