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
- [x] Meraki API V0 Support (Partial)
- [x] Meraki API V1 Support
- [x] Deprecation Policy
- [x] Sunset Policy
- [x] HTTP Response Status Code Handler
- [x] Unmarshal JSON Data
- [x] Marshal JSON Data
- [x] Marshal YAML Data

### Beta Release 2021 Q1
- [x] Pagination Support
- [x] verbose debug flag
- [x] Diff Flag (GET: compare config file against current API state)
- [x] org/net/device flags use exact match to obfuscate Ids with the actual names
- [x] 100% Get Command CLI Coverage
- [x] 100% Networks CLI Coverage
- [x] 100% Devices CLI Coverage
- [x] 100% Organizations CLI Coverage
- [x] Claim Assets into Org
- [x] 100% Insight CLI Coverage
- [x] 100% Cellular Gateway CLI Coverage
- [x] 100% Camera CLI Coverage

### Backlog
- [ ] Meraki Filter by Tags Flag 
- [ ] Proxy Support (certificate path, server, port)
- [ ] Hash passwords/secrets for dynamic encryption/decryption using ENV VAR to store seed secret (Encourage users not to store clear text creds in exported yaml/json config)

- [ ] 100% Appliance CLI Coverage
- [ ] 100% SM CLI Coverage
- [ ] 100% Switch CLI Coverage
- [ ] 100% Wireless CLI Coverage