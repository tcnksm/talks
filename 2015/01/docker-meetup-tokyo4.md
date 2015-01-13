# CoreOS

at [Docker meetup Tokyo #4](http://connpass.com/event/10318/)


## Overview

- Introduction (0.5)
- TL;DL (0.5, 1)
- What docker provide
- What docker can't resolve
- Why CoreOS (0.5, 1.5)
- What is CoreOS (1, 2.5)
    - Data center as a computer
- Features of CoreOS (4, 6.5)
    - Minimal Design
    - Docker
    - No Package manager & No runtimes ruby/python
    - Update
    - Clustering
- Technologies of CoreOS (4, 10.5)
    - Docker
    - etcd (イーティーシーディ)
    - fleet
    - cloud-config
- What is Good point of CoreOS ?
- How to play with CoreOS (10, 20.5)
    - Workflow
    - Test on Vagrant
    - Cluster by Terrafrom
    - Load-balancing
    - Container Metrics (Heapster)
    - *Logging*
    - Automated deploy through DockerHub
        - Git push -> automated build -> captainhook
- DEMO (4.5, 25)
- Conclusion

## Introduction

## TL;DL

## Docker provides

- Dockerイメージとして実行環境ごとアプリケーションをパッケージ化
- Docker Registry (DockerHub)によるイメージの配布
- 高速なアプリケーション起動・リソースの隔離化


```bash
$ docker pull tcnksm/docker-meetup-demo-web:1
$ docker run -p 80:80 tcnksm/docker-meetup-demo-web:1
```

## Docker can not solve

しかし，Dockerのみでは解決できない問題も多い．例えば，Dockerはシングルホストでは非常に強力だが，マルチホストな環境では非常に弱い．

問題

- サービスディスカバリー
    - コンテナAによるサービスははどのホストで動いている？
- オーケストレーション
    - 複数ホストにまたがったときにコンテナ同士を如何に協調させるか
- スケジューリング
    - コンテナをどのホストにデプロイするのか
    - せっかくDockerでホストを抽象化できたのに....
- ホストの環境の統一
    - Dockerの実行ホストに差が出てしまっては意味がない...
x    - Dockerで開発環境と本番の環境を一致させることができてバンザイ！ってなってもこの周辺ツールに誤差が出る．それだと意味がない
- 死活監視
    - コンテナが死んだらどうする？フェイルオーバーは？     

解決

- サービスディスカバリー
    - fleet
- オーケストレーション
    - etcdによるコンテナの設定共有
- スケジューリング
    - fleet
- ホストの環境の統一
    - cloud-config
    - shellshockみたいなのが起こったら...
- 死活監視
    - fleet

そしてとにかくシンプル！

## Why CoreOS

- CoreOS
    - オーケストレーションやサービスディスカバリの問題をデフォルトで解決
    - とにかくシンプル
        - Docker実行環境を簡単に準備できる
        - 簡単にスケールできる
        - 耐障害性もある
- コンテナオーケストレーションと言えばKubernetesだが，KubernetesはToo muchな場合も多い
    -  KubernetesはIaaSやPaaSの中間に位置する，基盤サービス
    - 一般的なウェブサービスのコンテナデプロイに使うのにはToo much過ぎる．
    - ローカルでテストとかめっちゃ大変そう（e.g., PaaSを立てる）

## What is CoreOS ?

GoogleやFacebook，Twitterといった企業が実現している柔軟かつスケーラブル，耐障害性の高いインフラの構築を目的としたLinuxディストリビューション

簡単に言ってしまうと，複数のホストを規模によって束ねて，そのどっかにコンテナをデプロイして面倒見ておいてくれるってのを簡単にやってくれる

## Features of CoreOS

- Minimal Operating system
- Painless Updating
- Docker containers
- Clustered by default


### Minimal Operating system

- 従来のLinuxディストリビューションが機能を追加することでその価値を高めていったのに対して，CoreOSは必要最低限まで機能を削ぎ落としていることに価値がある
- RAMの使用量は114MBで一般的なLinuxディストリビューションと比較しても40%も少ない
- DigitalOceanやAmazon EC2，OpenStackなどあらゆるプラットフォームで動作する

### Painless Updating

- CoreOSは安全かつ容易なOSアップデート機構を持っている
- OmahaというChromeOSやChromeの更新に利用されているUpdate Engineを利用
- RootFSを丸ごと入れ替えることでアップデートを行う
- これによりShellShockのような脆弱性が発見されても，いちいちパッチを当てるといったことやらずに済む

具体的なアップデートの仕組みを紹介する

### Docker

CoreOSは専用のパッケージマネージャーをもたない．またRubyやPythonといった言語のRuntimeも持たない．アプリケーションは全てDockerコンテナとして起動させる．

- これによりプロセスの隔離と，安全なマシンリソースの共有，アプリケーションのポータビリティという恩恵を受けることができる（Dockerのものの恩恵）

Dockerを使うことでホストは以下のようにいくつもの部屋が存在するように見えるようになる．


### Clustered by default

- CoreOSはクラスタリングの機構を標準で持っている．
- Datacenter as a Computer
- データセンターの大量のサーバー群からクラスタを構築してまるでそれが1つのコンピュータとして扱えるようにすることをゴールとしている

Docker + etcd + fleetでクラスタされたホストは以下のように見える．

## Technologies of CoreOS

### Docker

CoreOSは専用のパッケージマネージャーをもたないまたRubyやPythonといった言語のRuntimeも持たない．アプリケーションは全てDockerコンテナとして起動させる．

### cloud-config

- CoreOSを初期化・カスタマイズする仕組み
- cloud-initにインスパイア
- cloud-initから必要最低限な設定項目を抜き出したサブセット
- CoreOS特有の設定項目ももつ
- YAML形式で既述を行う
- 同じ設定ファイルから同じ環境を再現できる

```yaml
#cloud-config

coreos:
  etcd:
    discovery: https://discovery.etcd.io/<YOUR_TOKEN>
    addr: $private_ipv4:4001
    peer-addr: $private_ipv4:7001
  fleet:
    public-ip: $private_ipv4
    metadata: role=web,provider=digitalocean
  units:
    - name: etcd.service
      command: start
    - name: fleet.service
      command: start
  update:
    group: alpha
    reboot-strategy: best-effort
```
    
### etcd (イーティーシーディ)

- クラスタを形成するために，CoreOSはetcdという分散Key-Valuesストアを使い，各種設定をノード間で共有する
- etcdってのは”/etc distributed”の


### fleet

- クラスタ内で動いているコンテナのスケジューリングと管理を行う
- Fleetによるデプロイはめちゃ簡単
- Containerがどこのホストで動いているかを意識しなくてもよい
- Unit fileを書いてどんどんstartしていけばよい
- 人によってやり方が変わることがない

## How to play with CoreOS

CoreOSが良さそうなのはわかった．では実際に運用するときにはどうすればよいのか，を簡単に紹介する．

- 実際のWorkflowは？Dev環境とProduction環境
- CoreOSそのものを管理するには？
- サービスディスカバリーはどうやるの？
- 例としてDemoコンテナを複数立てて，それにロードバランシングする方法を紹介する

### Workflow

- Playstationチームの例
- Vagrantで簡単にローカルにクラスタ環境を構築できる．プロダクションと同様のcloud-configを使えば同様の環境を作ることができる
- サービスの起動はfleetの向き先を変えれば良い

### Build Cluster by Terrafrom

[http://deeeet.com/writing/2015/01/07/terraform-coreos/](http://deeeet.com/writing/2015/01/07/terraform-coreos/)

### Load-balancing with confd

例としてDemoコンテナを複数立てて，それにロードバランシングする方法を紹介する．

## DEMO

at [tcnksm/docker-meetup-4-demo](https://github.com/tcnksm/docker-meetup-4-demo)

## Conslusion

