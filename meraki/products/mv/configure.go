package mv

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/camera/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

var GetQualityAndRetention = &cobra.Command{
	Use:   "qualityAndRetention",
	Short: "Returns quality and retention settings for the given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetQualityAndRetention(serial)
		shell.Display(metadata, "QualityAndRetention", cmd.Flags())
	},
}

var PutQualityAndRetention = &cobra.Command{
	Use:   "qualityAndRetention",
	Short: "Returns quality and retention settings for the given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.QualityAndRetention
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutQualityAndRetention(serial,  input)
		shell.Display(metadata, "QualityAndRetention", cmd.Flags())
	},
}

var GetQualityRetentionProfiles = &cobra.Command{
	Use:   "qualityRetentionProfiles",
	Short: "List the quality retention profiles for this network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetQualityRetentionProfiles(networkId)
		shell.Display(metadata, "QualityRetentionProfiles", cmd.Flags())
	},
}

var GetQualityRetentionProfile = &cobra.Command{
	Use:   "qualityRetentionProfile",
	Short: "Retrieve a single quality retention profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		qualityRetentionProfileId:= args[0]

		metadata := configure.GetQualityRetentionProfile(networkId,
			qualityRetentionProfileId)
		shell.Display(metadata, "QualityRetentionProfile", cmd.Flags())
	},
}

var DelQualityRetentionProfile = &cobra.Command{
	Use:   "qualityRetentionProfile",
	Short: "Retrieve a single quality retention profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		qualityRetentionProfileId:= args[0]

		metadata := configure.DelQualityRetentionProfile(networkId,
			qualityRetentionProfileId)
		shell.Display(metadata, "QualityRetentionProfile", cmd.Flags())
	},
}

var PutQualityRetentionProfile = &cobra.Command{
	Use:   "qualityRetentionProfile",
	Short: "Retrieve a single quality retention profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		qualityRetentionProfileId:= args[0]
		var format configure.QualityRetentionProfile
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutQualityRetentionProfile(networkId,
			qualityRetentionProfileId,  input)
		shell.Display(metadata, "QualityRetentionProfile", cmd.Flags())
	},
}

var PostQualityRetentionProfile = &cobra.Command{
	Use:   "qualityRetentionProfile",
	Short: "Retrieve a single quality retention profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.QualityRetentionProfile
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostQualityRetentionProfiles(networkId, input)
		shell.Display(metadata, "QualityRetentionProfile", cmd.Flags())
	},
}

var GetSchedules = &cobra.Command{
	Use:   "schedules",
	Short: "Returns a list of all mv recording schedules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetCameraSchedules(networkId)
		shell.Display(metadata, "Schedules", cmd.Flags())
	},
}

var GetObjectDetectionModels = &cobra.Command{
	Use:   "objectDetectionModels",
	Short: "Returns the MV Sense object detection model list for the given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetObjectDetectionModel(serial)
		shell.Display(metadata, "ObjectDetectionModels", cmd.Flags())
	},
}

var GetSense = &cobra.Command{
	Use:   "sense",
	Short: "Returns sense settings for a given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetSense(serial)
		shell.Display(metadata, "Sense", cmd.Flags())
	},
}

var PutSense = &cobra.Command{
	Use:   "sense",
	Short: "Returns sense settings for a given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.Sense
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutSense(serial,  input)
		shell.Display(metadata, "Sense", cmd.Flags())
	},
}

var GetVideoSettings = &cobra.Command{
	Use:   "videoSettings",
	Short: "Returns video settings for the given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetVideoSettings(serial)
		shell.Display(metadata, "VideoSettings", cmd.Flags())
	},
}

var PutVideoSettings = &cobra.Command{
	Use:   "videoSettings",
	Short: "Returns video settings for the given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		var format configure.VideoSettings
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutVideoSettings(serial,  input)
		shell.Display(metadata, "VideoSettings", cmd.Flags())
	},
}

var GetVideoLink = &cobra.Command{
	Use:   "videoLink",
	Short: "Returns video link to the specified mv. If a timestamp is supplied, it links to that timestamp.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		metadata := configure.GetVideoLink(serial)
		shell.Display(metadata, "VideoLink", cmd.Flags())
	},
}

