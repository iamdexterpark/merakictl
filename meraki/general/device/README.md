# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/README.md)*
*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/meraki/README.md)*

## Device

All Device level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | managementInterface | merakictl get device managementInterface {serial} | | Return the management interface settings for a device.
 GET  | device | merakictl get device details {serial} |  | Return a single device.
 GET  | clients | merakictl get device clients {serial} | --t0 --timespan | List the clients of a device, up to a maximum of a month ago. The usage of each client is returned in kilobytes. If the device is a switch, the switchport is returned; otherwise the switchport field is null.
 GET  | lldpcdp | merakictl get device lldpcdp {serial} | | List LLDP and CDP information for a device.
 GET  | uplinkloss | merakictl get device uplinkloss {serial} | --t0 --t1 --timespan --resolution --uplink --ip | Get the uplink loss percentage and latency in milliseconds for a wired network device.
 
 