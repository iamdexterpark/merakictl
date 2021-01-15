# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/README.md)*
*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/meraki/README.md)*

## Organization

All Organization level API calls. 
 POST |  | merakictl | |
 GET  |  | merakictl | |
 DEL  |  | merakictl | |
 PUT  |  | merakictl | |
 
 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | list | merakictl get organization list | | List the organizations that the user has privileges on.
 GET  | details | merakictl get organization info {organizationId} | | List a specific organization that the user has privileges on.
 
 POST | actionBatch | merakictl post organization actionBatch | | Create an action batch.
 GET  | actionBatches | merakictl get organization actionBatches {organizationId} | | Return The List Of Action Batches In The Organization.
 GET  | actionBatch | merakictl get organization actionBatch {actionBatchId} {organizationId} | | Return A Single Action Batch.
 DEL  | actionBatch | merakictl del organization actionBatch | merakictl del organization actionBatch {actionBatchId} {organizationId} | | Delete an action batch.
 PUT  | actionBatch | merakictl put organization actionBatch | merakictl put organization actionBatch {actionBatchId} {organizationId} | | Update an action batch.
 

 
 GET  | admins | merakictl get organization admins {organizationId} | | List The Dashboard Administrators In This Organization.
 GET  | brandingPolicies | merakictl get organization brandingpolicies {organizationId} | | Return The Branding Policy IDs Of An Organization.
 GET  | brandingPolicy | merakictl get organization brandingPolicy {brandingPolicyId} {organizationId} | | Return The Branding Policies Of An Organization.
 GET  | configurationTemplates | merakictl get organization configurationTemplates {organizationId} | | List The Configuration Templates For This Organization.
 GET  | configurationTemplate | merakictl get organization configurationTemplate {configTemplateId} {organizationId} | | Return a Configuration Template For This Organization.
 GET  | devices | merakictl get organization devices {organizationId} | --perPage --startingAfter --endingBefore --configurationUpdatedAfter |  List the devices in an organization.
 GET  | inventoryDevices | merakictl get organization inventoryDevices {organizationId}  |  --perPage --startingAfter --endingBefore --usedState --search | Return The Device Inventory For An Organization.
 GET  | inventoryDevice | merakictl get organization inventoryDevice {deviceId} {organizationId} |  | Return a single device from the inventory of an organizatio.
 GET  | licences | merakictl get organization licences {organizationId} | --perPage --startingAfter --endingBefore --deviceSerial --networkId --state | List The Licenses For An Organization.
 GET  | licence | merakictl get organization licence {licenceId} {organizationId} | | List A Single License For An Organization.
 GET  | loginSecurity | merakictl get organization loginSecurity {organizationId} | | Returns The Login Security Settings For An Organization.
 GET  | networks | merakictl get organization networks {organizationId} | --configTemplateId --tags --tagsFilterType --perPage --startingAfter --endingBefore | List the networks that the user has privileges on in an organization.
 GET  | idps | merakictl get organization idps {organizationId} | | List the SAML IdPs in your organization.
 GET  | idp | merakictl get organization idp {ldpId} {organizationId} | | List a SAML IdP in your organization.
 GET  | saml | merakictl get organization saml {organizationId} | | Returns the SAML SSO enabled settings for an organization.
 GET  | samlRoles | merakictl get organization samlRoles {organizationId} | | List the SAML roles for this organization.
 GET  | samlRole | merakictl get organization samlRole {samalRoleId} {organizationId} | | List a single SAML role for this organization.
 GET  | snmp | merakictl get organization snmp {organizationId} | | Return the SNMP settings for an organization.
 GET  | apiRequestsOverview | merakictl get organization apiRequestsOverview {organizationId} | --t0 --t1 --timespan | Return an aggregated overview of API requests data.
 GET  | apiRequests | merakictl get organization apiRequests {organizationId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --adminId --path --method --responseCode --sourceIp | List the API requests made by an organization
 GET  | configurationChanges | merakictl get organization configurationChanges {organizationId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --networkId --adminId | View the Change Log for your organization.
 GET  | deviceStatuses | merakictl get organization deviceStatuses {organizationId}  | --perPage --startingAfter --endingBefore | List the status of every Meraki device in the organization.
 GET  | uplinksLossAndLatency | merakictl get organization uplinksLossAndLatency {organizationId} | --t0 --t1 --timespan --uplink --ip | Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago.
 GET  | licenseOverview | merakictl get organization licenseOverview {organizationId} | | Return an overview of the license state for an organization.
 GET  | openApiSpec | merakictl get organization openapiSpec {organizationId} | | Return the OpenAPI 2.0 Specification of the organization's API documentation in JSON.
 GET  | alertTypes | merakictl get organization alertTypes {organizationId} | | Return a list of alert types to be used with managing webhook alerts.
 GET  | webhookLogs | merakictl get organization webhookLogs {organizationId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --url | Return the log of webhook POSTs sent.
 