
# merakictl
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/ddexterpark/merakictl)

*多语言文档：[英文](../README.md), [日本语](README.ja.md)，[简体中文](README.zh-cn.md).*

社区开发的用于Meraki Dashboard API的Golang SDK和命令行工具。
有关供应商支持的API接口，请参见精彩内容：[Meraki仪表板API Python库]https://github.com/meraki/dashboard-api-python)

##下载Merakictl
该SDK可以用作基于CLI的应用程序，而无需任何先前的编程经验。
它可以编译成跨平台(Linux / Mac / Win)静态二进制文件。

当前，我们处于Alpha版本：[下载Merakictl](https://github.com/ddexterpark/merakictl/releases)

##从源代码运行

####安装
安装[Go](http://golang.org)编程语言。

####设置路径
``
导出GOPATH = $ HOME / go
导出PATH = $ PATH：$ GOPATH / bin
```

####下载项目

``
去获取github.com/ddexterpark/merakictl
```

####编译CLI(可选)
外壳脚本
    ＃Linux / MacOS
    去建立main.go
    mv main / usr / local / sbin / merakictl
    导出PATH = / usr / local / bin：/ usr / local / sbin：“ $ PATH”

    ＃Windows 64位
    env GOOS = windows GOARCH = amd64编译main.go
    
    ＃Windows 32位
    env GOOS = windows GOARCH = 386去建立main.go
```
    
＃＃ 环境变量

至少，要使用此工具，您将需要为API密钥设置一个环境变量。
您还可以设置一些可选的env var来自定义API调用：

####必填

** MERAKI_API_TOKEN **
外壳脚本
      重击-
      出口MERAKI_API_TOKEN = 1234567890987654321
      回声$ MERAKI_API_TOKEN
      
      电源外壳 -
            setx MERAKI_API_TOKEN“ 1234567890987654321”
            回声％MERAKI_API_TOKEN％
```
＃＃＃＃ 可选的
 
** MERAKI_API_URL **
外壳脚本
        默认='https://api.meraki.com/api/'
        中国='https://api.meraki.cn/api/'
```

** MERAKI_API_VERSION **

默认版本为v1，此工具仅对v0提供有限支持，因为该版本将于2022年停用。
并非所有端点都可以在v0中工作。
 
外壳脚本
    默认='v1'
```
    
＃＃ 句法

外壳脚本
    merakictl [命令] [子命令] [目标] [标志]
```

 
#### [命令]
 
操作|语法|描述
--- | --- | ---
| **声明** | merakictl索赔[订单/序列/许可证] [TARGET] |向仪表板索取订单，许可证，序列号。 |
| **创建** | merakictl创建[命令] [子命令] [目标] [标志] |创建(POST)新资源。 |
| **获取** | merakictl获取[命令] [子命令] [目标] [标志] |显示(GET)api资源的操作。 |
| **更新** | merakictl更新[命令] [子命令] [目标] [标志] |更新(PUT)目标资源。 |
| **删除** | merakictl删除[命令] [子命令] [目标] [标志] |破坏性(DELETE)API调用，用于从仪表板上删除资源。 |
| **版本** | merakictl版本|显示Merakictl的版本和相关发行信息。 |

#### [SUBCOMMAND]

子命令反映了Meraki仪表板层次结构：

对象描述
--- | --- |
|组织机构网络集合|
|网络|设备集合|
|设备| Meraki产品|


语法示例：
 外壳脚本
     merakictl获取组织列表
     merakictl获取组织网络
     merakictl获取网络警报设置
     merakictl获取MX L3防火墙规则
     merakictl获取ms端口
     merakictl得到ssid先生
 ```
为了调用这些命令，您将必须利用标志来指定目标org / network / device / switchport / ssid / etc ...

#### [标志]

标记类型|长|短|描述
--- | --- | --- | ---
|全球| **-导出** | -e |全局标志，用于通过get命令从Meraki API中提取配置。 |
|全球| **-输入** | |使用yaml文件将配置传递到仪表板API的全局标志|
|全球| **-diff ** | |与仪表板配置差异配置文件的全局标志
|全球| **-详细** | -v |全局标志，用于显示http请求和响应以进行故障排除。 |
|全球| **-组织** | -o |组织ID的全局标志。 |
|全球| **-网络** | -n |网络ID的全局标志。 |
|全球| **-主机名** | |设备主机名的全局标志。 |

###数据模型
请参阅示例目录以获取完整的参考指南。
我强烈建议您在运行create之前从仪表板导出当前配置，
更新或删除命令。

在大多数情况下，API命令会执行覆盖操作，** yaml文件中未捕获的任何仪表板配置都有被替换的危险。**


####示例Cmd
 外壳脚本
     merakictl更新vpn-输入vpnconf.yaml --org 12345678
 ```

#### vpnconf.yaml
``yaml
---
同行：
  -姓名：我的同龄人1
    publicIp：123.123.123.1
    私有子网：
      -192.168.1.0/24
      -192.168.128.0/24
    秘密：asdf1234
    ipsecPolicies：
      ikeCipherAlgo：
        -三倍
      ikeAuthAlgo：
        -sha1
      ikeDiffieHellmanGroup：
        -第2组
      寿命：28800
      childCipherAlgo：
        -aes128
      childAuthAlgo：
        -sha1
      childPfsGroup：
        -禁用
      儿童寿命：28800
    networkTags：
      -全部
  -姓名：我的同龄人2
    publicIp：123.123.123.2
    remoteId：miles@meraki.com
    私有子网：
      -192.168.2.0/24
      -192.168.129.0/24
    秘密：asdf56785678567856785678
    networkTags：
      - 没有
    ipsecPoliciesPreset：默认
```



###免责声明

Merakictl是一个非常强大的工具。使用仪表板速率限制，您具有以下理论上的潜力：


API调用数| 5 | 300 | 180,000 | 1,400,000 |
--- | --- | --- | --- | --- |
**完成时间** | ** 1秒** | ** 1分钟** | ** 1小时** | ** 8小时** |


进行生产更改时请小心。如果不确定该工具，请查看代码。

切勿运行您不理解的程序。在接触生产之前，请在测试环境中练习使用此工具。

创建生产变更计划。在测试环境中实施计划的各个方面，以确保对更改有最高的信心。

####出色的生产变更计划的要素包括：
-**同行审查**要求其他人审查您的测试计划，请他们在您的测试环境中运行它。
-**预先检查**捕获更改前的网络状态。
-**事后检查**捕获更改后的网络状态。
-**备份配置**复制配置，以便在发生回滚时可以重新应用它。
-**回滚过程**请勿轻易采取此步骤，否则会出错。
最糟糕的情况是更改失败并且没有经过测试的可靠回滚计划。
-**指数变更时间表**不要一次做所有事情。从单个网络开始，
监视它，给它时间正常运行，然后，如果没有错误，则调度下5个网络，然后是10、25、50、100等。
-**故障阈值**一批网络可接受的故障更改百分比
在取消所有计划的更改之前？通常，根据您的规模，可接受1-5％的费用。
一切都需要根本原因分析和计划修改。