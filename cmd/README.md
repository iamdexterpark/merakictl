# Merakictl CLI Reference 
  
## Root Commands
 
Command | Syntax | Description |
--- | --- | ---
| **claim** | merakictl claim [order/serial/licence] | Claim orders, licences, serial numbers into dashboard. |
| **create** | merakictl create [COMMAND] [SUBCOMMAND] [FLAGS] | Creates (POST) new resources. |
| **get** | merakictl get [COMMAND] [SUBCOMMAND] [FLAGS]| Operation for displaying (GET) api resources. |
| **update** | merakictl update [COMMAND] [SUBCOMMAND] [FLAGS]| Updates (PUT) targeted resources. |
| **remove** | merakictl remove [COMMAND] [SUBCOMMAND] [FLAGS]| Destructive (DELETE) API call for removing resources from the Dashboard. |
| **version** | merakictl version | Displays the version and associate release information of Merakictl. |


[Full List of Claim Commands](https://github.com/ddexterpark/merakictl/cmd/claim/README.md)

[Full List of Create Commands](https://github.com/ddexterpark/merakictl/cmd/create/README.md)

[Full List of Get Commands](https://github.com/ddexterpark/merakictl/cmd/get/README.md)

[Full List of Remove Commands](https://github.com/ddexterpark/merakictl/cmd/remove/README.md)

[Full List of Update Commands](https://github.com/ddexterpark/merakictl/cmd/update/README.md)

## Subcommands
 
Subcommand | Abbreviated | Description |
--- | --- | --- |
| **organization** | org | Collection of Networks. |
| **network** | net | Collection of Devices. | 
| **device** | sn | Meraki Products. | 
| **appliance** | **mx** | UTM appliances with SD-WAN capability included. | 
| **switch**  | **ms** | Branch and campus access and aggregation switches with remote packet capture and cable testing. |
| **wireless** | **mr** |  Industry-leading indoor and outdoor 802.11ax wireless access points with built-in location analytics. |
| **camera** | **mv** | Indoor and outdoor smart cameras, eliminating traditional external storage such as NVRs. |
| **gateway** | **mg** |  Cellular gateways seamlessly transpose a wireless cellular signal to wired Ethernet, providing primary or failover connectivity. |
| **systems** | **sm** |  A complete enterprise solution for provisioning, monitoring, and securing of mobile devices. |
| **insight** | **in** |  End-to-end visibility into end-users' experience on the network. |


## Flags

Flag Type | Long | Short | Description |
--- | --- | --- | ---
| Global | **--input** | | Global flag for using a yaml file for passing config to the Dashboard API |
| Global | **--export** | -e | Global flag for extracting config from the Meraki API via get commands. |
| Global | **--diff** | | Global flag for diffing a yaml/json file with config in the Dashboard API |
| Global | **--verbose** | -v | Global flag to display the http request & response for troubleshooting. |
| Global | **--organization** | -o | Global flag for matching a string against Organization ids. |
| Global | **--network** | -n | Global flag for matching a string against Network ids. |
| Global | **--hostname** | -h | Global flag for matching a string against device hostnames. |

## Common Filters

Filters are HTTP parameters that shape our API queries. The list of filters below is not exhaustive but provide details around those most commonly use.

Filter | Example | Description |
 --- | --- | --- 
t0 | |  The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
t1 | |  The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
timespan | |  The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
resolution | | The time resolution in seconds for returned data. The valid resolutions are: 600, 1200, 3600, 14400, 86400. The default is 86400.
autoResolution | | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. 
deviceSerial | | The serial of the Meraki device which the list of events will be filtered with
uplink | | Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, cellular. Default will return all uplinks.
ip | | Optional filter for a specific destination IP. Default will return all destination IPs. 
search | |  
configTemplateId | | An optional parameter that is the ID of a config template. Will return all networks bound to that template. 
tags | | An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). 
usedState | |  
tagsFilterType | | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. 
username | | The username of a Systems Manager user
email | | The email of a network user account or a Systems Manager device 
mac | | The MAC of a network client device or a Systems Manager device 
serial | | The serial of a Systems Manager device 
imei | | The IMEI of a Systems Manager device 
bluetoothMac | | The MAC of a Bluetooth client 
includeConnectivityHistory | |  Include the connectivity history for this client
connectivityHistoryTimespan | | The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used.
includedEventTypes[] | |  A list of event types. The returned events will be filtered to only include events with these types.
excludedEventTypes[] | |  A list of event types. The returned events will be filtered to exclude events with these types.
sensorSerial | | The serial of the sensor device which the list of events will be filtered with
gatewaySerial | | The serial of the environmental gateway device which the list of events will be filtered with
loginIdentifier | | The username, email, or phone number used during login 
deviceType | |  
sortOrder | | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order.
networkIds | |  
objectType | | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. 
clientId | | Filter results by network client to return per-device, per-band AP channel utilization metrics inner joined by the queried client's connection history.
apTag | | Filter results by AP tag to return AP channel utilization metrics for devices labeled with the given tag; either :clientId or :deviceSerial must be jointly specified.
band | | Filter results by band (either '2.4' or '5'). 
ssid | | Filter results by SSID
ssidNumber | | An SSID number to include. If not specified, events for all SSIDs will be returned.
vlan | | Filter results by VLAN
fields | | Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. 
wifiMacs | |
serials | |  
ids | | Filter users by id(s). 
scope | | Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags. 
withDetails | | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response 
perPage | | 
startingAfter | | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
endingBefore | | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.


The filters `perPage`, `startingAfter`, and `endingBefore` are part of the Meraki Dashboard APIs pagination strategy. 
They should be considered overrides as Merakictl is designed to automatically handle pagination. 
To learn more see [RFC5988 Web Linking](https://tools.ietf.org/html/rfc5988).

