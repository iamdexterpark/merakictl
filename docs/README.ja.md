
# merakictl
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/ddexterpark/merakictl)

*多言語ドキュメント：[英語](../README.md)[日本語](README.ja.md)、[简体中文](README.zh-cn.md).*

コミュニティが開発したGolangSDKと、Meraki DashboardAPI用のコマンドラインツール。
ベンダーがサポートするAPIインターフェースについては、すばらしい[Meraki Dashboard API Pythonライブラリ](https://github.com/meraki/dashboard-api-python) を参照してください。

## Merakictlをダウンロード
このSDKは、プログラミングの経験がなくても、CLIベースのアプリケーションとして使用できます。
クロスプラットフォーム(Linux / Mac / Win)の静的バイナリにコンパイルされます。

現在、アルファ版を使用しています：[Merakictlをダウンロード](https://github.com/ddexterpark/merakictl/releases)

##ソースコードから実行

####インストール
[Go](http://golang.org)プログラミング言語をインストールします。

####パスを設定
```bash
export GOPATH = $ HOME / go
PATH = $ PATH：$ GOPATH / binをエクスポートします
```

####プロジェクトをダウンロード

```bash
github.com/ddexterpark/merakictlを入手してください
```

#### CLIのコンパイル(オプション)
```シェルスクリプト
    ＃Linux / MacOS
    main.goをビルドします
    mv main / usr / local / sbin / merakictl
    export PATH = / usr / local / bin：/ usr / local / sbin： "$ PATH"

    ＃Windows64ビット
    env GOOS = windows GOARCH = amd64 go build main.go
    
    ＃Windows32ビット
    env GOOS = windows GOARCH = 386 go build main.go
```
    
＃＃ 環境変数

少なくとも、このツールを使用するには、APIキーの環境変数を設定する必要があります。
API呼び出しをカスタマイズするために設定できるオプションのenv変数もいくつかあります。

####必須

** MERAKI_API_TOKEN **
```シェルスクリプト
      バッシュ-
      エクスポートMERAKI_API_TOKEN = 1234567890987654321
      エコー$ MERAKI_API_TOKEN
      
      パワーシェル -
            setx MERAKI_API_TOKEN "1234567890987654321"
            エコー％MERAKI_API_TOKEN％
```
####オプション
 
** MERAKI_API_URL **
```シェルスクリプト
        デフォルト= 'https：//api.meraki.com/api/'
        中国= 'https：//api.meraki.cn/api/'
```

** MERAKI_API_VERSION **

デフォルトバージョンはv1です。このツールは、2022年に廃止されるため、v0のサポートが制限されています。
すべてのエンドポイントがv0で機能するわけではありません。
 
```シェルスクリプト
    デフォルト= 'v1'
```
    
##構文

```シェルスクリプト
    merakictl [COMMAND] [SUBCOMMAND] [TARGET] [flags]
```

 
#### [コマンド]
 
操作|構文|説明|
--- | --- | ---
| **クレーム** | merakictlクレーム[注文/シリアル/ライセンス] [ターゲット] |注文、ライセンス、シリアル番号をダッシュ​​ボードに請求します。 |
| **作成** | merakictl create [COMMAND] [SUBCOMMAND] [TARGET] [flags] |新しいリソースを作成(POST)します。 |
| **取得** | merakictl get [COMMAND] [SUBCOMMAND] [TARGET] [flags] | APIリソースを表示する(GET)操作。 |
| **更新** | merakictl update [COMMAND] [SUBCOMMAND] [TARGET] [flags] |更新(PUT)対象のリソース。 |
| **削除** | merakictl remove [COMMAND] [SUBCOMMAND] [TARGET] [flags] |ダッシュボードからリソースを削除するための破壊的(DELETE)API呼び出し。 |
| **バージョン** | merakictlバージョン| Merakictlのバージョンと関連リリース情報を表示します。 |

#### [サブコマンド]

サブコマンドは、Merakiダッシュボード階層を反映しています。

オブジェクト|説明
--- | --- |
|組織|ネットワークのコレクション|
|ネットワーク|デバイスのコレクション|
|デバイス|メラキ製品|


構文例：
```シェルスクリプト
     merakictl get org list
     merakictl getorgnetworks
     merakictlはネットワークアラート設定を取得します
     merakictl get mxl3ファイアウォールルール
     merakictl get ms port
     merakictl get mr ssid
```
これらのコマンドを呼び出すには、フラグを利用してターゲットのorg / network / device / switchport / ssid / etc ...を指定する必要があります。

#### [フラグ]

フラグタイプ|長い|ショート|説明|
--- | --- | --- | ---
|グローバル| **-エクスポート** | -e | getコマンドを介してMerakiAPIから構成を抽出するためのグローバルフラグ。 |
|グローバル| **-入力** | |設定をDashboardAPIに渡すためにyamlファイルを使用するためのグローバルフラグ|
|グローバル| **-diff ** | |ダッシュボード構成で構成ファイルを比較するためのグローバルフラグ|
|グローバル| **-詳細** | -v |トラブルシューティングのためのhttpリクエストとレスポンスを表示するグローバルフラグ。 |
|グローバル| **-組織** | -o |組織IDのグローバルフラグ。 |
|グローバル| **-ネットワーク** | -n |ネットワークIDのグローバルフラグ。 |
|グローバル| **-ホスト名** | |デバイスホスト名のグローバルフラグ。 |

＃＃＃ データ・モデル
完全なリファレンスガイドについては、examplesディレクトリを参照してください。
createを実行する前に、ダッシュボードから現在の構成をエクスポートすることを強くお勧めします。
コマンドを更新または削除します。

ほとんどの場合、APIコマンドは上書き操作を実行します。** yamlファイルにキャプチャされていないダッシュボード構成は置き換えられるリスクがあります。**


####コマンドの例
```シェルスクリプト
     merakictl update vpn --input vpnconf.yaml --org 12345678
```

#### vpnconf.yaml
```yaml
---
仲間：
  -名前：私の仲間1
    publicIp：123.123.123.1
    privateSubnets：
      -192.168.1.0 / 24
      -192.168.128.0 / 24
    シークレット：asdf1234
    ipsecPolicies：
      ikeCipherAlgo：
        -トリプルDES
      ikeAuthAlgo：
        --sha1
      ikeDiffieHellmanGroup：
        --group2
      ikeLifetime：28800
      childCipherAlgo：
        --aes128
      childAuthAlgo：
        --sha1
      childPfsGroup：
        - 無効
      childLifetime：28800
    networkTags：
      - すべて
  -名前：私の仲間2
    publicIp：123.123.123.2
    remoteId：miles@meraki.com
    privateSubnets：
      -192.168.2.0 / 24
      -192.168.129.0 / 24
    秘密：asdf56785678567856785678
    networkTags：
      - 無し
    ipsecPoliciesPreset：デフォルト
```



###免責事項

Merakictlは非常に強力なツールです。ダッシュボードのレート制限を使用すると、理論的に次のことが可能になります。


API呼び出しの数| 5 | 300 | 180,000 | 1,400,000 |
--- | --- | --- | --- | --- |
**完了する時間** | ** 1秒** | ** 1分** | ** 1時間** | ** 8時間** |


生産の変更を行う際には注意してください。ツールについて不明な点がある場合は、コードを確認してください。

わからないプログラムは絶対に実行しないでください。本番環境に触れる前に、テスト環境でこのツールを使用する練習をしてください。

生産変更計画を作成します。計画のあらゆる側面をテスト環境に実装して、変更に対する最高レベルの信頼性を確保します。

####優れた生産変更計画の要素は次のとおりです。
-**ピアレビュー**他の誰かにテスト計画をレビューしてもらい、テスト環境で実行するように依頼します。
-**事前チェック**変更前のネットワークの状態をキャプチャします。
-**事後チェック**変更後のネットワークの状態をキャプチャします。
-**バックアップ構成**ロールバックが発生した場合に再適用できるように、構成をコピーします。
-**ロールバック手順**この手順を軽く行わないでください。問題が発生します。
考えられる最悪の状況は、変更が失敗し、テスト済みの信頼できるロールバック計画がないことです。
-**指数関数的変更スケジュール**すべてを一度に実行しないでください。単一のネットワークから始めて、
それを監視し、正常に動作する時間を与えます。次に、問題がない場合は、次の5つのネットワーク、次に10、25、50、100などをスケジュールします。
-**失敗のしきい値**失敗した変更の何パーセントがネットワークのバッチで許容可能か
スケジュールされたすべての変更がキャンセルされる前に？通常、規模にもよりますが、1〜5％が許容されます。
それ以上のものは、根本原因の分析と計画の変更が必要です。