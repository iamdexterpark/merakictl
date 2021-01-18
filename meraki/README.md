# Merakictl CLI Reference 


*Main Documentation: [English](https://github.com/ddexterpark/merakictl/blob/master/README.md)*

#### COMMANDS

Long | Short | Syntax | Description |
--- | --- | ---  | ---
show | get | merakictl show [COMMAND] [SUBCOMMAND] [TARGET] [flags]| Operation for displaying (GET) api resources. 
create | post | merakictl create [COMMAND] [SUBCOMMAND] [TARGET] [flags] | Creates (POST) new resources. 
update | put | merakictl update [COMMAND] [SUBCOMMAND] [TARGET]  [flags]| Updates (PUT) targeted resources. 
remove | del, no | merakictl remove [COMMAND] [SUBCOMMAND] [TARGET]  [flags]| Destructive (DELETE) API call for removing resources from the Dashboard. 

#### SUBCOMMANDS

Subcommands mirror the Meraki Dashboard hierarchy.

Long | Short | Syntax | Description
--- | --- | ---  | ---
Organization | org |  | Collection of Networks
Network | net |  | Collection of Devices
Device | sn |  | Meraki Product
appliance | mx | | Meraki MX Security Appliance
switch | ms | | Meraki MS Switches
wireless | mr | | Meraki MR Wirelesss Access  Points
gateway | mg | | Meraki MG Cellualar Gateway
camera | mv | | Meraki MV Cameras
systems | sm | | Meraki SM Systems Management Solution
insight | in | | Meraki Insight Application Telemetry



#### FLAGS
In order to invoke most commands you will have to specify a target resource. 
For convenience, merakictl is capable of name-based resolution. 
This means you do not need to know the randomly generated IDs, 
only the exact name of the organization, network or device.

Flag Type | Long | Short | Description |
--- | --- | --- | ---
| Global | **--organization** | -o | Global flag for resolving names to OrganizationIds. |
| Global | **--network** | -n | Global flag for resolving names to NetworkIds. |
| Global | **--hostname** | -h | Global flag for resolving names to Device serial numbers. |
| Global | **--export** | -e | Global flag for extracting config from the Meraki API via get commands. |
| Global | **--input** | | Global flag for passing yaml/json config files as command input |
| Global | **--diff** | | Global flag for diffing config file with dashboard config |
| Global | **--verbose** | -v | Global flag to display the http request & response for troubleshooting. |


## Organization 
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/general/organization/README.md)

## Network
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/general/network/README.md)

## Device
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/general/device/README.md)

## MX Appliance 
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/products/mx/README.md)
 
## Camera MV
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/products/mv/README.md)
 
## Cellular Gateway mg
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/products/mg/README.md)
  
## Switch MS
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/products/ms/README.md)

## Wireless MR
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/products/mr/README.md)
 
## SM
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/products/sm/README.md)

## Insight
[English](https://github.com/ddexterpark/merakictl/tree/master/meraki/products/insight/README.md)


## Filters

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

