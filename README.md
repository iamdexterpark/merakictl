
# merakictl
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/ddexterpark/merakictl)

*Multi-language Documentation: [English](README.md), [日本語](README.ja.md), [简体中文](README.zh-cn.md).*

A Community developed command line tool for the Meraki Dashboard API. 
For the Vendor supported python SDK please see the wonderful: [Meraki Dashboard API Python Library](https://github.com/meraki/dashboard-api-python)

## Download Merakictl 
This cli does not require any prior programming experience or knowledge of Go Lang to use.

[Download Merakictl](https://github.com/ddexterpark/merakictl/releases)

## (Optional)  Compile CLI From Source Code
Merakictl compiles into a cross-platform (Linux/Mac/Win) static binary.
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

#### Compile Static Binary 
```shell script
    # Linux/MacOS
    cd /Users/{$USERNAME}/go/src/github.com/ddexterpark/merakictl
    go build main.go
    
    # Move executable to preferred location and declare in $PATH
    mv main /usr/local/sbin/merakictl
    export PATH=/usr/local/bin:/usr/local/sbin:"$PATH"
```

```shell script
    # Windows 64-bit
    env GOOS=windows GOARCH=amd64 go build main.go
    
    # Windows 32-bit
    env GOOS=windows GOARCH=386 go build main.go
```
Now this binary should be available in your shell from any location by calling `merakictl`.

## Initial Setup

#### AutoCompletion
To enable tab-based autocompletion use the following command with your preferred shell as input.
```shell script

    merakictl completion [bash|zsh|fish|powershell]
```

#### Set Environment Variables

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

    
## Syntax

[FUll CLI CMD GUIDE](https://github.com/ddexterpark/merakictl/cmd/README.md)

```shell script
    merakictl [COMMAND] [SUBCOMMAND] [TARGET]  [flags]
```

 
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
appliance | mx | |
switch | ms | |
wireless | mr | |
gateway | mg | |
camera | mv | |
systems | sm | |
insight | in | |

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


### Importing and Exporting Data
I highly recommended that you make a habit of exporting the current configuration from the dashboard before running create, 
update or delete commands.

Supported formats for importing and exporting data are JSON & YAML.

IMPORTANT: In most cases the API commands preform overwrite operations, meaning **any dashboard config not captured in the yaml file is at risk of being replaced.**


#### Usage Example
 ```shell script
    # export list of orgs to yaml file
     merakictl get org list -e 
 ```  

#### organizations.yaml
```yaml
---
- id: "100000000000000000"
  name: DextersLab
  samlConsumerUrl: https://n1.meraki.com/saml/login/A0bcDefg/hijO-kLmnoPq
  samlConsumerUrls:
  - https://n1.meraki.com/saml/login/A0bcDefg/hijO-kLmnoPq
  url: https://n1.meraki.com/o/A0bcDefg/manage/organization/overview
- id: "671599294431625609"
  name: DextersProd
  samlConsumerUrl: https://n2.meraki.com/saml/login/R1stUvwx/yzy2-xWvutSr
  samlConsumerUrls:
  - https://n2.meraki.com/saml/login/R1stUvwx/yzy2-xWvutSr
  url: https://n2.meraki.com/o/R1stUvwx/manage/organization/overview
```



### Disclaimer

Please use merakictl responsibly. 
It can be tempting to immediately try out a new tool to solve some production issues, 
but it is always better to use a test environment first and develop a change plan based on empirical evidence.

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
