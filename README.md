
# merakictl
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/ddexterpark/merakictl)

*Multi-language Documentation: [English](README.md), [日本語](README.ja.md), [简体中文](README.zh-cn.md).*

A Community developed Golang SDK and command line tool for the Meraki Dashboard API. 
For a Vendor supported API interface please see the wonderful: [Meraki Dashboard API Python Library](https://github.com/meraki/dashboard-api-python)

## Download Merakictl 
This SDK can be used as a CLI-based Application without any prior programming experience. 
It compiles into a cross-platform (Linux/Mac/Win) static binary. 

Currently, we are in Alpha version: [Download Merakictl](https://github.com/ddexterpark/merakictl/releases)

## Compile CLI From Source Code (Optional) 

#### Installation
Install the [Go](http://golang.org) programming language.

#### Set PATH
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

#### Download Project

```bash
go get github.com/ddexterpark/merakictl
```

#### Compile CLI (Optional) 
```shell script
    # Linux/MacOS
    cd /Users/{$USERNAME}/go/src/github.com/ddexterpark/merakictl
    go build main.go
    mv main /usr/local/sbin/merakictl
    export PATH=/usr/local/bin:/usr/local/sbin:"$PATH"

    # Windows 64-bit
    env GOOS=windows GOARCH=amd64 go build main.go
    
    # Windows 32-bit
    env GOOS=windows GOARCH=386 go build main.go
```

At this point the CLI should be available by calling `merakictl` from any locatio.
    
## Environment Variables

At minimum, to use this tool you will need to set an environmental variable for the API key. 
There are also some optional env vars that you can set to customize your API calls:

#### Required

**MERAKI_API_TOKEN**
```shell script
      Bash -
      export MERAKI_API_TOKEN=1234567890987654321
      echo $MERAKI_API_TOKEN 
      
      PowerShell -
            setx MERAKI_API_TOKEN "1234567890987654321"
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

#### Auto Completion
    
## Syntax

```shell script
    merakictl [COMMAND] [SUBCOMMAND] [TARGET]  [flags]
```

 
#### [COMMAND]
 
Operation | Syntax | Description |
--- | --- | ---
| **claim** | merakictl claim [order/serial/licence] [TARGET] | Claim orders, licences, serial numbers into dashboard. |
| **create** | merakictl create [COMMAND] [SUBCOMMAND] [TARGET] [flags] | Creates (POST) new resources. |
| **get** | merakictl get [COMMAND] [SUBCOMMAND] [TARGET] [flags]| Operation for displaying (GET) api resources. |
| **update** | merakictl update [COMMAND] [SUBCOMMAND] [TARGET]  [flags]| Updates (PUT) targeted resources. |
| **remove** | merakictl remove [COMMAND] [SUBCOMMAND] [TARGET]  [flags]| Destructive (DELETE) API call for removing resources from the Dashboard. |
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
| Global | **--export** | -e | Global flag for extracting config from the Meraki API via get commands. |
| Global | **--input** | | Global flag for using a yaml file for passing config to the Dashboard API |
| Global | **--diff** | | Global flag for diffing config file with dashboard config |
| Global | **--verbose** | -v | Global flag to display the http request & response for troubleshooting. |
| Global | **--organization** | -o | Global flag for Organization id. |
| Global | **--network** | -n | Global flag for Network id. |
| Global | **--hostname** |  | Global flag for Device Hostname. |

### Data Model
Please see the examples directory for full reference guide. 
I highly recommended that you export the current configuration from the dashboard before running create, 
update or delete commands.

In most cases the API commands preform overwrite operations, **any dashboard config not captured in the yaml file is at risk of being replaced.**


#### Example Cmd
 ```shell script
    # export list of orgs to yaml file
     merakictl get org list -e 
 ```  

#### organizations.yaml
```yaml
---
- id: "100000000000000000"
  name: MyOrganization
  samlConsumerUrl: https://n1.meraki.com/saml/login/A0bcDefg/hijO-kLmnoPq
  samlConsumerUrls:
  - https://n1.meraki.com/saml/login/A0bcDefg/hijO-kLmnoPq
  url: https://n1.meraki.com/o/A0bcDefg/manage/organization/overview
- id: "671599294431625609"
  name: YourOrganization
  samlConsumerUrl: https://n2.meraki.com/saml/login/R1stUvwx/yzy2-xWvutSr
  samlConsumerUrls:
  - https://n2.meraki.com/saml/login/R1stUvwx/yzy2-xWvutSr
  url: https://n2.meraki.com/o/R1stUvwx/manage/organization/overview
```



### Disclaimer

Merakictl is an extremely Powerful tool. With the Dashboard rate-limit, you have a theoretical potential to make:


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
The worst possible position is to have a change fail and not have a tested, reliable rollback plan.
- **Exponential Change Schedule** Don't do everything at once. Start with a single network, 
monitor it, give it time to operate normally, then, if nothing is wrong schedule the next 5 networks, then 10, 25, 50, 100, etc..
- **Failure threshold** What percentage of failed changes are acceptable in a batch of networks 
before all scheduled changes are canceled? Typically, 1-5% is acceptable depending on your scale. 
Anything over that needs a root cause analysis, and modification of your plan. 
