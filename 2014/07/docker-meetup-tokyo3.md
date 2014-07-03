The Age of Flynn
====

at [Docker meetup Tokyo #3](http://connpass.com/event/6998/)

## Talk overview

- Introduction(0.5)
    - Precondition
- TL;DR
- What is Flynn ?
- Background of Flynn (背景) (2)
    - How Heroku works
    - Why PaaS by Docker ?
- OSS PaaS by Docker (関連プロジェクト) (2)
    - dokku, building, Deis, Flynn, (Bazooka)
    - How dokku works
- Why Flynn ? (1)
    - vs Heroku
    - vs dokku
- Feature of flynn
- How flynn works (1)
    - Figure ...
- Architecture of Flynn (2)
    - Layer0
    - Layer1
- Layer0 (2)
    - flynn/flynn-host
        - flynn/sampi
    - flynn/discoverd
- Layer1 (2)
    - ...
- How Flynn use Docker (2)
    - Everything in Container
    - flynn/slugbuilder
    - flynn/slugrunner
- Demo (5)
- Road map of Flynn, Future of Flynn(0.5)
    - Container Independence
- Conlusion(0.5)
- References


## Introduction

- [SOTA](http://deeeet.com/writing/)

###  I build tool

- [tcnksm/dmux]() - docker + tmux = dmux
- [tcnksm/rbdock]() - Generate Dockerfile for Ruby or Rails, Sinatra
- [tcnksm/cli-init]() - The easy way to start building Golang command-line application
- [tcnksm/docc](): docc open your project documentation

## はじめに

- 本職はアプリケーションエンジニア
- Dockerは個人的に家で触って遊んでいる
- カジュアルユーザの立場から話す

### 前提

- Dockerの詳しい話はしない
- Dockerがどのようにうごくか分かっていることを前提とする

## TL;DR

Dockerの応用例の1つであるOSSのPaaSを構築することを目指すFlynnの紹介

- なぜFlynnのようなプロジェクトが登場してきたか
    - Dockerとの関連は？
- Flynnは既存のPaaSの何を変えるのか?
- FlynnがDockerをどのようにつかっているのか?
- 他のOSSのPaaSにはどんなものがあるか?
- Flynnのアーキテクチャー
- デモ

## What is Flynn

- Platform-as-a-Service
- OSS
- Crowdfunded
- Go言語

### Why Flynn's talk in Docker meetup ?

"Docker is great, but not enough"

- Dockerは面白し，触るととても感動する
- ただこれをどう使っていくのが良いかはまだまだわからない
- FlynnはDockerをいかに使うかのよい例

## Background

最初にDockerによるPaaSというプロジェクトがなぜ登場してきたかを簡単に説明する．

### How PaaS (Heroku) works

まず，PaaSで一番有名なHerokuの動作をそのワークフローで復習する．

- `heroku create`
    - Stackと呼ばれるベースとなるOSを準備する
          - e.g., Cedar stack
- `git push heroku master`
    - アプリケーションがデプロイされる
    - アプリケーションをビルドしてslug（SquashFS）を作成する
        - slug compiler
            - 各言語のBuildpackの集合
            - 依存関係のインストールなどを行う
              - e.g., RubyならGemfileをもとにrubygemsをインストール
        - slug
            - ソースと依存ライブラリ，言語のランタイムを含んだ圧縮されたファイルシステム(SquashFS)
    - アプリケーションの実行環境（Dyno）を準備する
        - Dyno
            - **LXCをベースにしたContainer環境**
    - Dynoにslugをロードする
    - Procfileをもとにアプリケーションを起動する
        - Procfile
            - プロセスの起動コマンドを記述
                - e.g., web: bundle exec rails server -p $PORT
- `heroku ps:scale web=2`(Heroku)
    - webのプロセスを増やす = 実行Dynoの数を増やす
        - https://s3-eu-west-1.amazonaws.com/jon-assettest/dynos.jpg
- `heroku run bash`
    - Dynoにログインする
        - 新しくDynoが準備され，最新のSlugが読み込まれる
        - 変更は他のDynoに影響を与えない

"Heorkuの重要な構成要素の１つはLXCベースの実行環境Dyno"

### Why PaaS by Docker ?

- Dockerによりカジュアルにコンテナ仮想化を扱えるようになった
- ファイルシステムをイメージとして保存，配布できるようになった
（Buildpackという標準の仕様があり，それが公開されている）


- Heroku slug = Docker image
    - 言語のランタイムやアプリケーションの依存を含んだファイルシステム
- Heroku Dyno = Docker container


## OSS PaaS by Docker

そういうわけで，Dockerを使ってPaaSを構築するプロジェクトが多く登場してきた．

- dokku
    - 100行のbashで書かれたシンプルなPaaS実装
- building
   - dokku - git-receive
- Deis
   - ≒ Flynn
   - マルチホストのPaaS
   - on CoreOS

### How dokku works

- bashによる実装
- DokkuはPaaSの最低限の機能を持つ

## Goal of Flynn

**"The product that ops provides to developers"**

- Easy deployment
    - git pushで
    - Docker containerで
- Run anything
    - どんな言語でも
    - どんなフレームワークでも
- Painless Scaling
    - 新たにノードを追加するだけで，スケールもしくはクラスタを追加できる
- Simple and composable
    - 簡単に各コンポーネントを変更したり，切り替えたりできる
        - すべての機能がモジュールとして実装
        - APIとして実装

## Why Flynn ?

Herokuの簡便さとAmazon EC2のような自由度を兼ね備えたPaaS

### vs. Heroku

- Heroku++
    - http://progrium.com/blog/2014/02/06/the-start-of-the-age-of-flynn/

### vs. dokku

- dokku ++
    - マルチホストに対応
    - 分散システムの上にのったdokkuとも見なすことができる
    - 個々の機能が全てモジュール化され，dockerコンテナになっている

## How flynn works

## Architecture of Flynn

Flynnのアーキテクチャはシンプルかつ理解しやすいようにデザインされている

- ほとんど全てのシステムがDockerコンテナでコンテナ内で動く
    - Flynnを構成するコンポーネントは，Flynnにデプロイされるサービスやアプリケーションと
変わらないと考えることができる

**flynn/flynn-demo**を見るとよくわかる，インストールはすべてdocker pullで済ませている

```ruby
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "flynn-base"
  ...

  config.vm.provision "shell", inline: <<SCRIPT
    # layer0
    docker pull flynn/host
    docker pull flynn/discoverd
    docker pull flynn/etcd

    # layer1
    docker pull flynn/postgres
    docker pull flynn/controller
    docker pull flynn/gitreceive
    docker pull flynn/strowger
    docker pull flynn/shelf
    docker pull flynn/slugrunner
    docker pull flynn/slugbuilder
    docker pull flynn/bootstrap

    docker run -e=DISCOVERD=${IP_ADDR}:1111 flynn/bootstrap
SCRIPT
end
```

Flynnは2つのレイヤーで構成される

- **layer0 (The Grid)**
    - 低位なリソースフレームワーク層
    - この上の全てのアプリケーションやサービスの基礎
    - コンテナ管理, サービスディスカバリー, タスクスケジューラー, 分散型KVS（etcd）
- **layer1**
    - 高位なコンポーネント層
    - Git-receive，Heroku Buildpack，DB，HTTP Routing，

## Layer0

### Overview fo Layer0

クラスタ内の全てのホストで実行

- [D] flynn/discoverd
    - https://github.com/flynn/discoverd
    - サービスディスカバリー
- [H] flynn/flynn-host
    - https://github.com/flynn/flynn-host
    - コンテナ管理，タスクスケジューラー

### [D] flynn/discoverd

- サービスディスカバリー
    - クラスタの構築
    - メンバーの参加/離脱イベントを通知
    - クラスタのリーダの選出
- クライアント (flynn/go-discoverd)
    - クラスタへのホストの登録/削除
    - 各ホストのアドレスやメタ情報，新しいホストの参加/離脱のイベントの購読
- バックエンドに**etcd**を利用
    - ZooKeeperなどに入れ替え可能

### [H] flynn/flynn-host

- Dockerコンテナ管理
    - HTTP API経由で複数ホストの複数のDockerコンテナを管理
- 役割
    - リーダー
        - クラスタ全体のホストの一覧とそれらが実行しているジョブの管理
        - 新しいジョブの登録
    - それ以外
        - リーダからのジョブをDockerコンテナで実行
        - 指定されたジョブの停止
        - 実行中のジョブ，指定されたジョブの情報の返答
- [S] ジョブスケジューラー
    - flynn/flynn-host/sampi
        - スケジューラーフレームワーク
        - シンプルな Mesos
- ジョブ

```go
type Job struct {
	ID string

	// Job attributes
	Attributes map[string]string
	// Number of TCP ports required by the job
	TCPPorts int

	Config     *docker.Config
	HostConfig *docker.HostConfig
}

```

```go
// https://github.com/flynn/go-dockerclient/blob/master/docker.go

type Config struct {
	Hostname        string
	Domainname      string
	User            string
	Memory          int64 // Memory limit (in bytes)
	MemorySwap      int64 // Total memory usage (memory + swap); set `-1' to disable swap
	CpuShares       int64 // CPU shares (relative weight vs. other containers)
	AttachStdin     bool
	AttachStdout    bool
	AttachStderr    bool
	PortSpecs       []string // Deprecated - Can be in the format of 8080/tcp
	ExposedPorts    map[string]struct{}
	Tty             bool // Attach standard streams to a tty, including stdin if it is not closed.
	OpenStdin       bool // Open stdin
	StdinOnce       bool // If true, close stdin after the 1 attached client disconnects.
	Env             []string
	Cmd             []string
	Dns             []string
	Image           string // Name of the image as it was passed by the operator (eg. could be symbolic)
	Volumes         map[string]struct{}
	VolumesFrom     string
	WorkingDir      string
	Entrypoint      []string
	NetworkDisabled bool
	Name            string `json:"-"`
}
```

## Layer1

### Overview fo Layer1

- [G] flynn/gitreceived
    - git pushをうけることに特化したSSHサーバ
- [C] flynn/flynn-controller
    - HTTP APIサーバー
    - (≒ Heroku Platform API)
- [SB] flynn/slugbuilder
    - DockerとBuildpackで（Heroku的な）slugの作成
- [SR] flynn/slugrunner
　　　- slugbuilderで作成されたslugの実行
- [R] flynn/strowger
    - TCP/HTTP ルータ（リバースプロキシ）
- [DB] flynn/flynn-postgres
    - Flynn用のPostgreSQL
- flynn/flynn-cli
    - コマンドラインクライアント
    - (≒ heroku-toolbelt)


### [C] flynn/flynn-controller

HTTP APIサーバー

- Flynn上で動く全てのアプリケーションをHTTP APIで操作する
- HerokuのPlatform APIにインスパイアされている
- flynn-cliクライアントから利用できる

### [G] flynn/gitreceived

git pushをうけることに特化したSSHサーバ

1. git pushによりアプリケーションのコードがデプロイされる
2. ssh-keyによる認証を行う
3. pushされたアプリケーションのtarを引数に`receiver`スクリプトを実行する

### flynn/flynn-gitreceived

gitreceivedを使ってアプリケーションのビルドと実行を行う

`receiver`スクリプトはflynn-controllerに対して以下を要請する

1. git pushされたアプリケーションのtarをslugbuilderに渡してビルドする（slugの作成）
    - slug(.tgz)をshelf（ファイルサーバ）上に作成する
2. 作成されたビルドをslugrunnerで実行する

### [SB] flynn/slugbuilder

DockerとBuildpackで（Heroku的な）slugの作成する

Herokuの場合（slug-compiler）
-> buildpackでアプリケーションのビルド
-> ビルドをslug（SquashFS:Dynoで利用可能な圧縮ファイルシステム）として保存

Flynnの場合
-> Dockerコンテナ内（クリーン！！）でBuildpackを使ってアプケーションのビルド
-> ビルドを.tgz形式に固めて保存


```bash
$ id=$(git archive master | docker run -i -a stdin flynn/slugbuilder) # buildpackによるビルド
$ docker wait $id
$ docker cp $id:/tmp/slug.tgz . # コンテナ内のslug.tgzを取り出す
```

### [SR] flynn/slugrunner

slugbuilderで作成された（Heroku的な）slugを実行する

Herokuの場合
-> Dynoにslugをロードする
-> Procfileをもとにアプリケーションを起動する

Flynnの場合
-> Dockerコンテナ内にslug(.tgz)をロードする
-> ProcfileをもとにDockerコンテナを起動する


直接.tgzを標準入力からコンテナに渡す

```bash
$ cat myslug.tgz | docker run -i -a stdin -a stdout flynn/slugrunner start web
```

slugのURLを与える

```bash
$ docker run -e SLUG_URL=http://example.com/slug.tgz -i -t flynn/slugrunner start web
```

### [R] flynn/strowger

TCP/HTTP ルータ

https://commons.wikimedia.org/wiki/File:HebdrehwaehlerbatterieOrtsvermittlung_4954.jpg

- ランダムロードバランシングのためのリバースプロキシとして動作する
- サービスディスカバリーによりバックエンドで何が起動したかを常に監視する
    - e.g., 複数のslugrunner（アプリケーション）に対してランダムにリクエストを振り分ける
- vs. HAProxy/nginx
    - サービスディスカバリーによる動的な設定変更が可能
　　　    - HAProxyやnginxは設定更新のために新しいプロセスが必要になる
    - (だたしまだプロトタイプなので，プロダクション環境ではサービスディスカバリー+HAProxyの方がよい)

### flynn/flynn-cli

最後に出力されるコマンドをコピペしてサーバを登録する

```bash
flynn server-add -g localhost:2201 \
    -p boAKKeTdRqNVRdsmVmXCA8jwucmYLWKTZLNIUK+qcmc= \
    default https://localhost:8081 \
    61ae3daf3b95ab2f2542ee66980c84ba
```

SSH keyを登録する

```bash
flynn key-add
```

アプリケーションをデプロイする

```bash
flynn create example
git push flynn master
```

webプロセスを作る

```bash
flynn scale web=1
```

ルートを設定する(?)

```bash
flynn route-add-http localhost:8080
```

スケールする

```bash
flynn scale web=3
```

### Other services in Layer1

- flynn/shelf
    - シンプルで，高速なHTTPのファイルサービス
    - (≒シンプルなAmazon S3)
- flynn/flynn-bootstrap
    - Layer1の起動
- flynn/taffy
   - レポジトリをpullしてFlynnにデプロイ

## DEMO

- AWSとDigitalOceanのアカウントがあれば利用可能
- 5分程度で立ち上がる


### serverの登録（server.go）

```bash
$ flynn server-add \
    -g lxr.flynnhub.com \ #githost
    -p GHj+Q74JItrqrf+kn7jHCRX3rBEe6cuzkolRcohQrHQ= \ # SHA256 of the server's TLS cert
    "flynn-demo" \ # server-name
    https://lxr.flynnhub.com \ # server url
    69e3ea9510205c532c4133b1b6c7e2c9 # key
```

Usage

```bash
$ server-add [-g <githost>] [-p <tlspin>] <server-name> <url> <key>
```

中身を見る（消せば何度でも利用可能）

```bash
cat ~/.flynnrc
```

### SSH-KEYを登録する

```bash
$ flynn key-add
```

### アプリケーションの作成（app.go）

中身，git remoteをしつつ，コントローラーに伝えている

```go
exec.Command("git", "remote", "remove", "flynn").Run()
exec.Command("git", "remote", "add", "flynn", gitURLPre(serverConf.GitHost)+app.Name+gitURLSuf).Run()
```

```bash
flynn create node-demo
Created node-demo
```

### アプリケーションのデプロイ

nodeのサンプルを上げる

```bash
$ git push flynn master
The authenticity of host 'lxr.flynnhub.com (107.170.153.129)' cant be established.
RSA key fingerprint is 6b:d2:2b:16:85:f7:2b:cf:7c:60:6b:02:fd:24:87:23.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'lxr.flynnhub.com,107.170.153.129' (RSA) to the list of known hosts.
Counting objects: 18, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (17/17), done.
Writing objects: 100% (18/18), 1.84 KiB | 0 bytes/s, done.
Total 18 (delta 6), reused 0 (delta 0)
-----> Building node-demo...
-----> Node.js app detected
-----> Requested node range: 0.10.x
-----> Resolved node version: 0.10.29
-----> Downloading and installing node
-----> Writing a custom .npmrc to circumvent npm bugs
-----> Installing dependencies
       npm WARN package.json node-example@0.0.1 No description
       npm WARN package.json node-example@0.0.1 No repository field.
       npm WARN package.json node-example@0.0.1 No README data
-----> Cleaning up node-gyp and npm artifacts
-----> Building runtime environment
-----> Discovering process types
       Procfile declares types -> web
-----> Compiled slug size is 5.2M
-----> Creating release...
=====> Application deployed
To ssh://git@lxr.flynnhub.com/node-demo.git
 * [new branch]      master -> master
```

### ルートの表示

```bash
$ flynn routes
ROUTE                            SERVICE        ID
http:node-demo.lxr.flynnhub.com  node-demo-web  http/65ae6884514622f0e7dcd85f71724814
```

### ルートの追加（route.go）

```bash
$ flynn route-add-http docker-meetup.lxr.flynnhub.com
$ flynn routes
ROUTE                                SERVICE        ID
http:docker-meetup.lxr.flynnhub.com  node-demo-web  http/8aecdfad5adb70ddff6a2ac32b79da4d
http:node-demo.lxr.flynnhub.com      node-demo-web  http/65ae6884514622f0e7dcd85f71724814
```

### ルートの削除

```bash
$ flynn route-remove http/baf06246e4ec7f3486a89a2e313205ba
```

### プロセスの確認

```bash
flynn ps
```

### ログの確認

```bash
flynn log e4cffae4ce2b-8cb1212f582f498eaed467fede768d6f
```

### スケールする

```bash
$ flynn scale web=8
```

### Docker コンテナに入る

```bash
flynn run bash
```
### 緊急対応

```bash
ssh n0.lxr.flynnhub.com
```
-> 無理



## Roadmap of Flynn

"Production環境で安定して動作させることが最優先"
    - テストカバレッジの向上と，セキュリティに注力

### 今後の予定 (2014.07現在)

- 数週間以内にFlynn Beta（社内環境向け）をリリース
- 夏の終わりにFlynn 1.0をリリース


## Future of Flynn

"Container Indepnedence"

- FlynnのユーザはDocker Incや他の企業に結びつけられるべきではない

### pinkerton

- a tool for using Docker images with other container runners
    - Docker(libcontainer), systemd+nspawn, LXC, libvirt-lxc, Lmctfy, OpenVZ
- https://cloud.githubusercontent.com/assets/13026/3220006/68b7cc1a-effb-11e3-8386-64e34b9c54d8.png


## Conclusion


## References

- Flynn - Open source Platform as a Service powered by Docker
    - https://flynn.io/
- The Start of the Age of Flynn
    - http://progrium.com/blog/2014/02/06/the-start-of-the-age-of-flynn/
- Unveiling Flynn, a new PAAS based on Docker
    - http://jpetazzo.github.io/2013/11/17/flynn-docker-paas/
- 5by5 | The Changelog #99: Flynn, Tent, open source PaaSes and more with Jeff Lindsay and Jonathan Rudenberg
    - http://5by5.tv/changelog/99
- 5by5 | The Changelog #115: Flynn updates with Jonathan Rudenberg and Jeff Lindsay
    - http://5by5.tv/changelog/115
- Container Independence
    - https://flynn.io/blog/container-indepedence
- Bazooka: Continuous Deployment at SoundCloud - Google Slides
    - https://docs.google.com/presentation/d/1ni1BFiVMTLN_8q34si-qPfl0F6w2W3oPVKqSMfcs_XA/edit#slide=id.p
- Deis: Evolution of a Docker PAAS
    - http://gabrtv.github.io/deis-dockercon-2014/#/
- Flynn vs. Deis: The Tale of Two Docker Micro-PaaS Technologies | CenturyLink Labs
    - http://www.centurylinklabs.com/flynn-vs-deis-the-tale-of-two-docker-micro-paas-technologies/
- Inside Dokku in 5 minutes
    - http://banyan.me/slides/20140116/slides.html
- Discoverd - r7km/s
    - http://r7kamura.github.io/2014/06/24/discoverd.html
- Flynn Host - r7km/s
    - http://r7kamura.github.io/2014/06/26/flynn-host.html
- Beyond Flynn, or Flynn-as-a-Worldview
    - http://progrium.com/blog/2014/07/01/beyond-flynn-or-flynn-as-a-worldview/
