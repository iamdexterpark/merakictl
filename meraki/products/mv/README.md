# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/README.md)*
*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/tree/master/meraki/README.md)*

## Camera MV  
 
 All MV level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
 GET  | qualityAndRetention | merakictl get mv qualityAndRetention {serial} | | Returns quality and retention settings for the given camera.
 GET  | qualityRetentionProfiles | merakictl get mv qualityRetentionProfiles {networkId} | | List the quality retention profiles for this network.
 GET  | qualityRetentionProfile | merakictl get mv qualityRetentionProfile {qualityRetentionProfileId} {networkId} | | Retrieve a single quality retention profile.
 GET  | schedules | merakictl get mv schedules  {networkId} | | Returns a list of all camera recording schedules.
 GET  | objectDetectionModels | merakictl get mv objectDetectionModelss {serial} | | Returns the MV Sense object detection model list for the given camera.
 GET  | sense | merakictl get mv sense {serial} | | Returns sense settings for a given camera.
 GET  | videoSettings | merakictl get mv videoSettings {serial} | | Returns video settings for the given camera.
 GET  | videoLink | merakictl get mv videoLink {serial} | --timestamp | Returns video link to the specified camera. If a timestamp is supplied, it links to that timestamp.
 GET  | liveAnalytics | merakictl get mv liveAnalytics {serial} | | Returns live state from camera of analytics zones.
 GET  | analyticsOverview | merakictl get mv analyticsOverview {serial} | --t0 --t1 --timespan --objectType | Returns an overview of aggregate analytics data for a timespan.
 GET  | recentAnalytics | merakictl get mv recentAnalytics {serial} | --objectType | Returns most recent record for analytics zones.
 GET  | analyticsZonesHistory | merakictl get mv analyticsZonesHistory {zoneId} {serial} | --t0 --t1 --timespan --resolution --objectType | Return historical records for analytic zones.
 GET  | analyticsZones | merakictl get mv analyticsZones {serial} | | Returns all configured analytic zones for this camera.
 
 