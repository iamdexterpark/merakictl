# merakictl
A command line tool for interacting with the Meraki Dashboard API.

### Download Application
This application is compiled into a single binary so that it is cross-platform. 

Link TBD...

### Compiling Application from source
```shell script
    go build main.go
    mv main /usr/local/sbin/merakictl
    export PATH=/usr/local/bin:/usr/local/sbin:"$PATH"
```
    
### Environment Variables

At minimum, to use this tool you will need to set an environmental variable for the API key. 
There are also some optional env vars that you can set to customize your API calls:

#### Required
** MERAKI_API_TOKEN**
```shell script
      Bash -
      export MERAKI_API_TOKEN=1234567890987654321
      echo $MERAKI_API_TOKEN 
      
      PowerShell -
            setx MERAKI_API_TOKEN "234567890987654321"
            echo %MERAKI_API_TOKEN%
```
#### Optional
 
**MERAKI_API_URL**
```shell script
        Default = 'https://api.meraki.com/api/'
        China = 'https://api.meraki.cn/api/' 
```

**MERAKI_API_VERSION**

The default version is v1, this tool has limited support for v0 as it is being sunset in 2022. 
Not all endpoints will work in v0.
 
```shell script
    Default = 'v1'
```
    
### Syntax

```shell script
    merakictl [COMMAND] [SUBCOMMAND] [flags]
```

 
#### [COMMAND]
 
Operation	| Syntax | Description |
--- | --- | ---
| **claim** | merakictl claim [SUBCOMMAND] [ | Claim orders, licences, serial numbers into dashboard. |
| **create** | merakictl create [COMMAND] [TARGET] [SUBCOMMAND] [flags] | Creates (POST) new resources. |
| **get** | merakictl get [COMMAND] [TARGET] [SUBCOMMAND] [flags]| Operation for displaying (GET) api resources. |
| **update** | merakictl update [COMMAND] [TARGET] [SUBCOMMAND] [flags]| Updates (PUT) targeted resources. |
| **remove** | merakictl remove [COMMAND] [TARGET] [SUBCOMMAND] [flags]| Destructive (DELETE) API call for removing resources from the Dashboard. |
| **version** | merakictl version | Displays the version and associate release information of Merakictl. |


#### [SUBCOMMAND]

Subcommands mirror the Meraki Dashboard hierarchy: 

Object | Description
--- | --- |
| Organization | Collection of Networks |
| Network | Collection of Devices |
| Device | Meraki Products |


Example syntax:
 ```shell script
     merakictl get org list
     merakictl get org networks
     merakictl get network alert settings
     merakictl get mx l3 firewall rules
     merakictl get ms port
     merakictl get mr ssid
 ```   
In order to invoke these commands you will have to leverage flags to specify the target org/network/device/switchport/ssid/etc...

#### [FLAGS]

Flag Type | Long | Short | Description |
--- | --- | --- | ---
| Global | **--input** | | Global flag for using a yaml file for passing config to the Dashboard API |
| Global | **--export** | -e | Global flag for extracting config from the Meraki API via get commands. |
| Global | **--org** | -o | Global flag for Organizion id, needed for most api calls. |
| Global | **--net** | -n | Global flag for Network id, needed for most device based api calls. |
| Global | **--verbose** | -v | Global flag to display the http request & response for troubleshooting. |

### Data Model
Please see the examples directory for full reference guide. 
I highly recommended that you export the current configuration from the dashboard before running create, 
update or delete commands.

In most cases the API commands preform overwrite operations, **any dashboard config not captured in the yaml file is at risk of being replaced.**


#### Example Cmd
 ```shell script
     merakictl update vpn --input vpnconf.yaml --org 12345678
 ```  

#### vpnconf.yaml
```yaml
---
peers:
  - name: My peer 1
    publicIp: 123.123.123.1
    privateSubnets:
      - 192.168.1.0/24
      - 192.168.128.0/24
    secret: asdf1234
    ipsecPolicies:
      ikeCipherAlgo:
        - tripledes
      ikeAuthAlgo:
        - sha1
      ikeDiffieHellmanGroup:
        - group2
      ikeLifetime: 28800
      childCipherAlgo:
        - aes128
      childAuthAlgo:
        - sha1
      childPfsGroup:
        - disabled
      childLifetime: 28800
    networkTags:
      - all
  - name: My peer 2
    publicIp: 123.123.123.2
    remoteId: miles@meraki.com
    privateSubnets:
      - 192.168.2.0/24
      - 192.168.129.0/24
    secret: asdf56785678567856785678
    networkTags:
      - none
    ipsecPoliciesPreset: default
```



### Disclaimer

Meraki-ctl is an extremely Powerful tool. With the Dashboard rate-limit, you have a theoretical potential to make:


Number of API Calls | 5 | 300 | 180,000 | 1,400,000 |
--- | --- | --- | --- | --- |
**Time to Complete** | **1 second** | **1 minute** | **1 hour** | **8 hours** |


Please be cautious when making production changes. If you are unsure about the tool, look at the code. 

Never run a program that you do not understand. Practice using this tool in a test environment before touching production.

Create a production change plan. Implement every aspect of your plan in a test environment to ensure the highest levels of confidence in your change.   

#### Elements of a great production change plan include:
- **Peer Review** Have someone else review your test plan, ask them to run it in your test environment.
- **Pre-checks**  Capture the state of the network before the change.
- **Post-checks** Capture the state of the network after the change. 
- **Backup Config** Copy the config so that you can re-apply it in the event of a rollback.
- **Rollback Procedure** Do not take this step lightly, things go wrong.
- **Exponential Change Schedule** Don't do everything at once. Start with a single network, 
monitor it, give it time to operate normally, then, if nothing is wrong schedule the next 5 networks, then 10, 25, 50, 100, etc..
- **Failure threshold** What percentage of failed changes are acceptable in a batch of networks 
before all scheduled changes are canceled? Typically, 5% is acceptable, anything over that needs a root cause analysis, and modification of your change plan. 
