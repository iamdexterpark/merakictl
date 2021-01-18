# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/blob/master/README.md)*

*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/tree/master/meraki/README.md)*

## Wireless MR
 
 All MR level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | alternateManagementInterface | merakictl get mr alternateManagementInterface {networkId} | | Return alternate management interface and devices with IP assigned.
  GET  | bluetoothNetworkSettings | merakictl get mr bluetoothNetworkSettings {networkId} | | Return the Bluetooth settings for a network. Bluetooth settings must be enabled on the network.
  GET  | bluetoothDeviceSettings | merakictl get mr bluetoothDeviceSettings {serial} | | Return the bluetooth settings for a wireless device.
  GET  | radioSettings | merakictl get mr radioSettings {serial} | | Return the radio settings of a device.
  GET  | rfProfiles | merakictl get mr rfProfiles {networkId} | --includeTemplateProfiles | List the non-basic RF profiles for this network.
  GET  | rfProfile | merakictl get mr rfProfile {rfProfileId} {networkId} | | Return a RF profile.
  GET  | wirelessSettings | merakictl get mr wirelessSettings {networkId} | | Return the wireless settings for a network.
  GET  | l3FirewallRules | merakictl get mr l3FirewallRules {number} {networkId} | | Return the L3 firewall rules for an SSID on an MR network.
  GET  | l7FirewallRules | merakictl get mr l7FirewallRules {number} {networkId} | | Return the L7 firewall rules for an SSID on an MR network.
  GET  | identityPSKs | merakictl get mr identityPSKs {number} {networkId} | | List all Identity PSKs in a wireless network.
  GET  | identityPSK | merakictl get mr identityPSK {number} {identityPskId} {networkId} | | Return an Identity PSK.
  GET  | splashSettings | merakictl get mr splashSettings {number} {networkId}| | Display the splash page settings for the given SSID.
  GET  | trafficShapingRules | merakictl get mr trafficShapingRules {number} {networkId} | | Display the traffic shaping settings for a SSID on an MR network.
  GET  | ssids | merakictl get mr ssids {networkId} | | List the MR SSIDs in a network.
  GET  | ssid | merakictl get mr ssid {number} {networkId}| | Return a single MR SSID.
  GET  | airMarshal | merakictl get mr airmarshal {networkId} | --t0 --timespan | List Air Marshal scan results from a network.
  GET  | channelUtilizationHistory | merakictl get mr channelUtilizationHistory {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band | Return AP channel utilization over time for a device or network client.
  GET  | clientCountHistory | merakictl get mr clientCountHistory {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid | Return wireless client counts over time for a network, device, or network client.
  GET  | connectionStat | merakictl get mr connectionStat {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag | Aggregated connectivity info for a given client on this network. Clients are identified by their MAC.
  GET  | connectionStats  | merakictl get mr connectionStats  {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag  |  Aggregated connectivity info for this network, grouped by clients.
  GET  | connectivityEvents | merakictl get mr connectivityEvents {clientId} {networkId} | --perPage --startingAfter --endingBefore --t0 --t1 --timespan --types --includedSeverities --band --ssidNumber --deviceSerial | List the wireless connectivity events for a client within a network in the timespan.
  GET  | latencyHistory | merakictl get mr latencyHistory {clientId} {networkId} | --t0 --t1 --timespan --resolution | Return the latency history for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP. The latency data is from a sample of 2% of packets and is grouped into 4 traffic categories: background, best effort, video, voice. Within these categories the sampled packet counters are bucketed by latency in milliseconds..
  GET  | latencyStat | merakictl get mr latencyStat {clientId} {networkId}| --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for a given client on this network. Clients are identified by their MAC.
  GET  | latencyStats | merakictl get mr latencyStats {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for this network, grouped by clients.
  GET  | deviceConnectionStats | merakictl get mr deviceConnectionStats {serial} | --t0 --t1 --timespan --band --ssid --vlan --apTag | Aggregated connectivity info for a given AP on this network.
  GET  | networkConnectionStats | merakictl get mr networkConnectionStats {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag | Aggregated connectivity info for this network.
  GET  | dataRateHistory | merakictl get mr dataRateHistory {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid | Return PHY data rates over time for a network, device, or network client.
  GET  | connectionStats | merakictl get mr connectionStats {networkId} |--t0 --t1 --timespan --band --ssid --vlan --apTag | Aggregated connectivity info for this network, grouped by node.
  GET  | deviceLatencyStats | merakictl get mr deviceLatencyStats {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for this network, grouped by node.
  GET  | failedConnections | merakictl get mr failedConnections {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag {serial} {clientId} | List of all failed client connection events on this network in a given time range.
  GET  | latencyHistoryAverage | merakictl get mr latencyHistoryAverage {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid --accessCategory | Return average wireless latency over time for a network, device, or network client.
  GET  | aggregatedLatencies | merakictl get mr aggregatedLatencies {serial} | --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for a given AP on this network.
  GET  | networkLatencyStats | merakictl get mr networkLatencyStats {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for this network.
  GET  | meshStatuses | merakictl get mr meshstatuses {networkId} | --perPage --startingAfter --endingBefore | List wireless mesh statuses for repeaters.
  GET  | signalQualityHistory | merakictl get mr signalQualityHistory {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid | Return signal quality (SNR/RSSI) over time for a device or network client.
  GET  | status | merakictl get mr status {serial} | | Return the SSID statuses of an access point.
  GET  | usageHistory | merakictl get mr usageHistory {networkId} |  --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid | Return AP usage over time for a device or network client.

