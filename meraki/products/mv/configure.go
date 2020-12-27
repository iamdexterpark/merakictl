package mv

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/camera/configure"
	shell "github.com/ddexterpark/merakictl/utilities"
	"github.com/spf13/cobra"
)

var GetQualityAndRetention = &cobra.Command{
	Use:   "QualityAndRetention",
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
	Use:   "QualityAndRetention",
	Short: "Returns quality and retention settings for the given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutQualityAndRetention(serial, format)
		shell.Display(metadata, "QualityAndRetention", cmd.Flags())
	},
}

var GetQualityRetentionProfiles = &cobra.Command{
	Use:   "QualityRetentionProfiles",
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
	Use:   "QualityRetentionProfile",
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
	Use:   "QualityRetentionProfile",
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
	Use:   "QualityRetentionProfile",
	Short: "Retrieve a single quality retention profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		qualityRetentionProfileId:= args[0]
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutQualityRetentionProfile(networkId,
			qualityRetentionProfileId, format)
		shell.Display(metadata, "QualityRetentionProfile", cmd.Flags())
	},
}

var PostQualityRetentionProfile = &cobra.Command{
	Use:   "QualityRetentionProfile",
	Short: "Retrieve a single quality retention profile.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostQualityRetentionProfiles(networkId,
			 format)
		shell.Display(metadata, "QualityRetentionProfile", cmd.Flags())
	},
}

var GetSchedules = &cobra.Command{
	Use:   "Schedules",
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
	Use:   "ObjectDetectionModels",
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
	Use:   "Sense",
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
	Use:   "Sense",
	Short: "Returns sense settings for a given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutSense(serial, format)
		shell.Display(metadata, "Sense", cmd.Flags())
	},
}

var GetVideoSettings = &cobra.Command{
	Use:   "VideoSettings",
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
	Use:   "VideoSettings",
	Short: "Returns video settings for the given mv.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, serial := shell.ResolveFlags(cmd.Flags())
		if serial == "" {
			serial = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutVideoSettings(serial, format)
		shell.Display(metadata, "VideoSettings", cmd.Flags())
	},
}

var GetVideoLink = &cobra.Command{
	Use:   "VideoLink",
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

