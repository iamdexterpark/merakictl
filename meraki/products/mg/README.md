# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/blob/master/README.md)*

*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/tree/master/meraki/README.md)*

## Cellular Gateway mg  
  
  All MG level API calls.
  
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
 GET  | connectivityMonitor | merakictl get mg connectivityMonitor {networkId} | | Return the connectivity testing destinations for an MG network.
 GET  | dhcp | merakictl get mg dhcp {networkId} | | List common DHCP settings of MGs.
 GET  | lan | merakictl get mg lan {serial} | | Show the LAN Settings of a MG.
 GET  | portForwardingRules | merakictl get mg portForwardingRules {serial} | | Returns the port forwarding rules for a single MG.
 GET  | subnetPool | merakictl get mg subnetPool {networkId} | | Return the subnet pool and mask configured for MGs in the network.
 GET  | uplink | merakictl get mg uplink {networkId} | | Returns the uplink settings for your MG network.
 