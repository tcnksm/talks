# Overview of CoreOS

[@CoreOS Meetup Tokyo #1](http://coreos-meetup-tokyo.connpass.com/event/12596/)

## Goal

他の発表についていけるように簡単にCoreOSとは何か，なぜ登場したのかを簡単におさらいしておく．

## Overview

- What is CoreOS
    - Googleやtwitterのような企業が実現しているスケーラブルかつ耐障害性のあるインフラを構築するために設計された新しいLinuxディストリビューション
- Who develop CoreOS
    - CEO, Alex Polvi
    - CTO, Brandon Philips
- Why CoreOS is started (What is initial motivation of CoreOS)
    - インターネットのセキュリティを根本的に改善したい
    - セキュリティはいかにソフトウェアをアップデートできるかにかかっている（完璧なソフトウェアは存在しない）
    - 自動的にアップデートするOSを作ろう！
    - 最新のソフトウェアを利用することでセキュリティだけではなく，信頼性，パフォーマンスを享受することができる
- What is features of CoreOS
    - ミニマルなデザイン（RAMの使用量は114MBしかない）
        - 従来のLinuxディストリビューションが機能を追加することで価値を高めていったのに対してCoreOSは必要最低限まで機能を削ぎ落としていることに価値がある
    - ChromeやChromeOSのように自動でOSがアップデートされる(RootFSを入れ替える)
    - アプリケーションはすべてコンテナ内で動く(言語のruntimeやパッケージマネージャーもない)
    - 標準でクラスタリング機構をもつ
        - "Data center as a Computer"
        - "サーバーの台数は減ることはないby Mitchell Hashimoto"
- What technologies CoreOS uses
    - Docker
    - etcd（イーティーシーディー），分散K&V
    - fleet (flt), 分散init system
    - rkt, Container runtime
        - appc/spec
    - flannel, networking
- Who uses CoreOS, or CoreOS technology
    - Deis, PaaS
    - CloudFoundry, PaaS
    - memsql, for distibuted test
    - Playstation, for dev environment
- What is new of CoreOS
    - Google Venture等から1200万ドルを調達
    - Tectonic
    - CoreOS Fest    
