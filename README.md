
# merakictl
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/ddexterpark/merakictl)

*Multi-language Documentation: [English](README.md), [日本語](docs/README.ja.md), [简体中文](docs/README.zh-cn.md).*

Merakictl is a community developed command line tool for interfacing with the Meraki Dashboard API. 
It allows you to query and alter network configs from a familiar cli-based environment. 

This tool leverages a community supported Go Lang Library: [dashboard-api-golang](https://github.com/ddexterpark/dashboard-api-golang)

For the vendor supported python Library: [dashboard-api-python](https://github.com/meraki/dashboard-api-python)


## Download Merakictl
This CLI tool does not require any prior programming experience or knowledge of Go Lang to use.
Please check the releases page to download the latest version.
Use this reference chart to select the appropriate version:

### [Download Merakictl](https://github.com/ddexterpark/merakictl/releases)

| $OS     | $ARCH                                                      |
|---------|------------------------------------------------------------|
| Linux   | [386](https://github.com/ddexterpark/merakictl/releases)   |
| Linux   | [amd64](https://github.com/ddexterpark/merakictl/releases) |
| OSX     | [386](https://github.com/ddexterpark/merakictl/releases)   |
| OSX     | [amd64](https://github.com/ddexterpark/merakictl/releases) |
| Windows | [386](https://github.com/ddexterpark/merakictl/releases)   |
| windows | [amd64](https://github.com/ddexterpark/merakictl/releases) |

#### Export Path to merakictl

```shell script
    # Linux/MacOS

    # Move executable to preferred location and append in $PATH
    mv merakictl-$OS-$ARCH /usr/local/sbin/merakictl
    export PATH=/usr/local/bin:/usr/local/sbin:"$PATH"

    # Windows
    # Move executable to executable location (Powershell)
    Move-Item -Path "merakictl-windows-$ARCH.exe" -Destination "C:\Program Files\merakictl.exe"

```
Now this binary should be available in your shell from any location by calling `merakictl`.


## Compile CLI From Source Code (Optional)
The following steps are very optional and intended only for those who wish to compile from source code. 

For everyone else, I preemptively compiled this tool into static binaries for multiple platforms.

Please see the `Download Merakictl` section for details.

#### Install Go Lang (Optional)
Install the [Go](http://golang.org) programming language.

#### Set Go Lang PATH (Optional)
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
#### Import Project (Optional)
```bash
go get github.com/ddexterpark/merakictl
```

#### Compile Static Binary (Optional)
```shell script
    cd /Users/{$USERNAME}/go/src/github.com/ddexterpark/merakictl
    go build main.go

    # Make binary exicutable on *nix systems
    chmod +x main.go

```


## Initial Setup  

#### Set Environment Variables

At minimum, to use this tool you will need to set an environmental variable for the API key. 
There are also some optional env vars that you can set to customize your API calls:

#### Required

**MERAKI_DASHBOARD_API_KEY**
```shell script
      Bash -
      export MERAKI_DASHBOARD_API_KEY=1234567890987654321
      echo $MERAKI_DASHBOARD_API_KEY
      
      PowerShell -
            setx MERAKI_DASHBOARD_API_KEY "1234567890987654321"
            echo %MERAKI_DASHBOARD_API_KEY%
```
[Generate a Meraki Dashboard API Key]( https://documentation.meraki.com/General_Administration/Other_Topics/The_Cisco_Meraki_Dashboard_API)

#### Optional Variables

This variable allows you to change the base URL used when making API calls. 
This is useful for targeting specific shards or regions.

By default, we leverage the mega-proxy, the Meraki API's Global Load-Balancer.
 
**MERAKI_DASHBOARD_API_URL**
```shell script
        Default = "https://api-mp.meraki.com/api/"
        China = "https://api.meraki.cn/api/" 
```

**MERAKI_DASHBOARD_API_VERSION**

The default version is v1, this tool has limited support for v0 as it is being sunset in 2022. 
Not all endpoints will work in v0.
 
```shell script
    Default = 'v1'
```

#### AutoCompletion
To enable tab-based autocompletion use the following command with your preferred shell as input.
```shell script

    merakictl completion [bash|zsh|fish|powershell]
```

## Syntax

The full command guide is available [here](https://github.com/ddexterpark/merakictl/tree/master/docs/commands/README.md)

```shell script
    merakictl [COMMAND] [SUBCOMMAND] [TARGET]  [flags]
```

#### COMMANDS

| Long   | Short   | Syntax                                                    | Description                                                              |
|--------|---------|-----------------------------------------------------------|--------------------------------------------------------------------------|
| show   | get     | merakictl show [COMMAND] [SUBCOMMAND] [TARGET] [flags]    | Operation for displaying (GET) api resources.                            |
| create | post    | merakictl create [COMMAND] [SUBCOMMAND] [TARGET] [flags]  | Creates (POST) new resources.                                            |
| update | put     | merakictl update [COMMAND] [SUBCOMMAND] [TARGET]  [flags] | Updates (PUT) targeted resources.                                        |
| remove | del, no | merakictl remove [COMMAND] [SUBCOMMAND] [TARGET]  [flags] | Destructive (DELETE) API call for removing resources from the Dashboard. |

#### SUBCOMMANDS

Subcommands mirror the Meraki Dashboard hierarchy.

| Long         | Short | Syntax | Description                           |
|--------------|-------|--------|---------------------------------------|
| Organization | org   |        | Collection of Networks                |
| Network      | net   |        | Collection of Devices                 |
| Device       | sn    |        | Meraki Product                        |
| appliance    | mx    |        | Meraki MX Security Appliance          |
| switch       | ms    |        | Meraki MS Switches                    |
| wireless     | mr    |        | Meraki MR Wirelesss Access  Points    |
| gateway      | mg    |        | Meraki MG Cellualar Gateway           |
| camera       | mv    |        | Meraki MV Cameras                     |
| systems      | sm    |        | Meraki SM Systems Management Solution |
| insight      | in    |        | Meraki Insight Application Telemetry  |


#### Usage Example
 ```shell script
    merakictl show organization list
    merakictl show org networks --organization 'DextersLab'
    merakictl show network devices --network 'My Network'  -o 'DextersLab'
    merakictl show mr ssids -n 'My Network'  -o 'DextersLab' --export
    merakictl update mr ssid 0 -n 'My Network'  -o 'DextersLab' --import SSIDS.yaml
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
