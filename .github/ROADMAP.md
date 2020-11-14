## RoadMap

### Alpha Release 2021 Q1 (OCT)
- [x] Basic cli structure (claim, create, get, update, remove)
- [x] Read environmental variables
- [x] Error Handling
- [x] Unit-testing

- [x] Traceback (OS/Merakictl version, env vars, HTTP req/resp)
- [x] Export Config flag
- [x] Import Config flag
- [x] Organization flag

- [x] HTTP GET/POST/PUT/DELETE 
- [x] Authorization
- [x] Exponential Backoff Policy
- [x] Connection Reuse Policy
- [x] Meraki API V0 Support
- [x] Meraki API V1 Support
- [x] Deprecation Policy
- [x] Sunset Policy
- [x] HTTP Response Status Code Handler
- [x] Unmarshal JSON Data
- [x] Marshal JSON Data
- [x] Marshal YAML Data

### Beta Release 2021 Q1 (NOV)
- [x] 100% Get Command API Coverage
- [x] Pagination Support
- [x] verbose debug flag
- [ ] Claim Assets into Org
- [ ] 100% Get Command CLI Coverage

### Backlog
- [ ] org/net/device flags use exact match to obfuscate Ids with the actual names
- [ ] Meraki Tags Flag
- [ ] Diff Flag (GET: compare config file against current API state)
- [ ] Proxy Support (certificate path, server, port)
- [ ] Action Batches (Bulk Operations)
- [ ] External Key Vault (Encourage users not to store creds in exported yaml/json config)
