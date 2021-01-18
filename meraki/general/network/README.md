# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/README.md)*
*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/tree/master/meraki/README.md)*

## Network

All Network level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | details | merakictl get network details {networkId} | | Return a network.
 GET  | alertSettings | merakictl get network alertSettings {networkId} | | Return the alert configuration for this network.
 GET  | clientPolicy | merakictl get network clientPolicy {clientId} {networkId} | | Return the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | splashAuthorizationStatus | merakictl get network splashAuthorizationStatus {clientId} {networkId} | | Return the splash authorization for a client, for each SSID they've associated with through splash. Only enabled SSIDs with Click-through splash enabled will be included. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | devices | merakictl get network devices {networkId} | | List the devices in a network.
 GET  | firmwareUpgrades | merakictl get network firmwareUpgrades {networkId} | | Get current maintenance window for a network.
 GET  | floorPlans | merakictl get network floorPlans {networkId} | | List the floor plans that belong to your network.
 GET  | floorPlan | merakictl get network floorPlan {floorPlanId} {networkId} | | Find a floor plan by ID.
 GET  | groupPolicies | merakictl get network groupPolicies {networkId} | | List the group policies in a network.
 GET  | groupPolicy | merakictl get network groupPolicy {groupPolicyId} {networkId} | | Display a group policy.
 GET  | merakiAuthUsers | merakictl get network merakiAuthUsers {networkId} | | List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network).
 GET  | merakiAuthUser | merakictl get network merakiAuthUser {merakiAuthUserId} {networkId} | | Return the Meraki Auth splash guest, RADIUS, or client VPN user.
 GET  | mqttBrokers | merakictl get network mqttBrokers {networkId} | | List the MQTT brokers for this network.
 GET  | mqttBroker | merakictl get network mqttBroker {mqttBrokerId} {networkId} | | Return an MQTT broker.
 GET  | netflow | merakictl get network netflow {networkId} | | Return the NetFlow traffic reporting settings for a network.
 GET  | channelUtilization | merakictl get network channelUtilization {networkId} | --t0 --t1 --timespan --resolution --perPage --startingAfter --endingBefore | Get the channel utilization over each radio for all APs in a network.
 GET  | piiKeys | merakictl get network piiKeys {networkId} | --username --email --mac --serial --imei --bluetoothMac | List the keys required to access Personally Identifiable Information (PII) for a given identifier. Exactly one identifier will be accepted. If the organization contains org-wide Systems Manager users matching the key provided then there will be an entry with the key "0" containing the applicable keys.
 GET  | piiRequests | merakictl get network piiRequests {networkId} | | List the PII requests for this network or organization.
 GET  | piiRequest | merakictl get network piiRequest {requestId} {networkId} | | Return a PII request.
 GET  | smDevices | merakictl get network smDevices {networkId} | --username --email --mac --serial --imei --bluetoothMac  | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier. These device IDs can be used with the Systems Manager API endpoints to retrieve device details. Exactly one identifier will be accepted.
 GET  | smOwners | merakictl get network smOwners {networkId} | --username --email --mac --serial --imei --bluetoothMac  | Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier. These owner IDs can be used with the Systems Manager API endpoints to retrieve owner details. Exactly one identifier will be accepted.
 GET  | settings | merakictl get network settings {networkId} | | Return the settings for a network.
 GET  | snmp | merakictl get network snmp {networkId} | | Return the SNMP settings for a network.
 GET  | syslog | merakictl get network syslog {networkId} | | List the syslog servers for a network.
 GET  | trafficAnalysis | merakictl get network trafficAnalysis {networkId} | | Return the traffic analysis settings for a network.
 GET  | trafficShaping | merakictl get network trafficShaping {networkId} | | Returns the application categories for traffic shaping rules.
 GET  | dscpTaggingOptions | merakictl get network dscpTaggingOptions {networkId} | | Returns the available DSCP tagging options for your traffic shaping rules.
 GET  | httpServers | merakictl get network httpServers {networkId} | | List the HTTP servers for a network.
 GET  | httpServer | merakictl get network httpServer {httpServerId} {networkId} | | Return an HTTP server for a network.
 GET  | webhookTest | merakictl get network webhookTest {weebhookTestId} {networkId} | | Return the status of a webhook test for a network.
 GET  | bluetoothclients | merakictl get network bluetoothClients {networkId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --includeConnectivityHistory | List the Bluetooth clients seen by APs in this network.
 GET  | bluetoothClient | merakictl get network bluetoothClient {bluetoothClientId} {networkId} | --includeConnectivityHistory --connectivityHistoryTimespan | Return a Bluetooth client. Bluetooth clients can be identified by their ID or their MAC.
 GET  | clientTrafficHistory | merakictl get network clientTrafficHistory {clientId} {networkId} | --perPage --startingAfter --endingBefore | Return the client's network traffic data over time. Usage data is in kilobytes. This endpoint requires detailed traffic analysis to be enabled on the Network-wide > General page. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | clientUsageHistory | merakictl get network clientUsageHistory {clientId} {networkId} | | Return the client's daily usage history. Usage data is in kilobytes. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | networkClients | merakictl get network networkClients {networkId} | --t0 --timespan --perPage --startingAfter --endingBefore | List the clients that have used this network in the timespan.
 GET  | clientIdentifier | merakictl get network clientIdentifier {clientId} {networkId} | | Return the client associated with the given identifier. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | environmentalEvents | merakictl get network environmentalEvents {networkId} | --includedEventTypes[] --excludedEventTypes[] --sensorSerial --gatewaySerial --perPage --startingAfter --endingBefore | List the environmental events for the network.
 GET  | events | merakictl get network events {networkId} | --productType --includedEventTypes[] --excludedEventTypes[] --deviceMac --deviceSerial --deviceName --clientIp --clientMac --clientName --smDeviceMac --smDeviceName --perPage --startingAfter --endingBefore | List the events for the network.
 GET  | splashLoginAttempts | merakictl get network splashLoginAttempts {networkId} | --splashLoginAttempts --ssidNumber --loginIdentifier --timespan | List the splash login attempts for a network.
 GET  | traffic | merakictl get network traffic {networkId} | --t0 --timespan --deviceType | Return the traffic analysis data for this network. Traffic analysis with hostname visibility must be enabled on the network.

