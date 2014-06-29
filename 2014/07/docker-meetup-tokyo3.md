The Age of Flynn
====

at [Docker meetup Tokyo #3](http://connpass.com/event/6998/)

## Talk overview

- Introduction(0.5)
    - Precondition
- TL;DR
- What is Flynn ?
- Background of Flynn (背景) (1)
    - How Heroku works
    - QA. Why PaaS by Docker ?
- OSS PaaS by Docker (関連プロジェクト) (2)
    - Cloud Foundry, dokku, building, Deis, Flynn, (Bazooka)
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

## なぜDocker meetupでFlynnの話をするのか?

"Docker is fun, but not enough"

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

- Heroku slug = Docker image
    - 言語のランタイムやアプリケーションの依存を含んだファイルシステム
- Heroku Dyno = Docker container

- Dockerによりカジュアルにコンテナ仮想化を扱えるようになった
- ファイルシステムをイメージとして保存，配布できるようになった
（Buildpackという標準の仕様があり，それが公開されている）


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


## Why Flynn ?

### vs Heroku

### vs dokku

- bashによる実装
- DokkuはPaaSの最低限の機能を持つ


## How flynn works (1)
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

##


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
