# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/blob/master/README.md)*

*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/tree/master/docs/commands/README.md)*


## Insight  
 
 All Insight level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | monitoredMediaServers | merakictl get insight monitoredMediaServers {organizationId} | | List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.
  GET  | monitoredMediaServer | merakictl get insight monitoredMediaServer {monitoredMediaServerId} {organizationId} | | Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.
  
  