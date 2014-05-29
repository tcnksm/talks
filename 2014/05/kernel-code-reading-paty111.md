Build your own PaaS by Docker
====

at [Kernel Code Reading Party #111](http://kernel.doorkeeper.jp/events/10433)


## TL;DR

Dockerを使ってOSSなPlatform-as-a-Service(PaaS)を構築することを目指すプロジェクトの紹介

## PaaS

- Platform-as-a-Service
- Heroku, Google App Engine, dotCloud
- Softwareを稼働させるためのプラットフォームを提供

## How Heroku works

- `heroku create`
    - Prepare stack (Cedar, Bamboo)
        - Cedar: based on Ubuntu 10.04
- `git push`
    1. Prepare Dyno (隔離されたアプケーションの実行環境)
        - Dyno: Lightweight container
        - based on **LXC (Linux Container)**
    2. slug compiler prepares your code for execution on Dyno
        - slug compiler: collection of scripts called buildpack
        - buildpack: fetch all dependencyies and language runtime, compile
        - 例えばRailsならGemfileを元にGemのインストール
    3. Procfileをもとにアプケーションの起動

## Docker?

Dockerによってコンテナがより簡単に作れるようになった

## Build a PaaS by Docker

### Heroku ++

- [Flynn]()
- [Deis]()

### Single Host

- [dokku]()
- [building]()

## Flynn & Deis

- Heroku ++

## dokku

Docker powered mini-Heroku. The smallest PaaS implementation you've ever seen. written by bash


### 要素技術

Heroku

- Dyno (LXC)
- slug compiler (buildpack)

dokku

- Docker image (progrium/buildstep)
- Buildpack

## How dokku works

`git push`
    1. Prepare Docker image (progrium/buildstep)
        - git push されたアプリケーションをイメージ内にぶっ込む
    2. /build/builderスクリプトによるアプリケーションのビルド
        - bashスクリプト
        - 各言語のbuildpackのインストール
        - アプリケーションの言語の判別
        - 判別された言語のbuildpackをつかって依存パッケージのインストール及びコンパイル
    3. Procfileをもとにアプケーションの起動
        - `docker run -d $IMAGE /bin/bash -c "/start.sh web"`


```bash
# start.sh
ruby -e "require 'yaml';puts YAML.load_file('Procfile')['\$1']" | bash
```

```yaml
web: bundle exec rackup config.ru -p $PORT -s thin
```

## building

A simple dev tool CLI for building Docker containers for any app using Heroku Buildpacks

### 要素技術

Heroku

- Dyno (LXC)
- slug compiler (buildpack)

dokku

- Docker image (progrium/buildstep)
- Buildpack

building

- Docker image (ctlc/buildstep forked from progrium/buildstep)
- Buildpack

### dokku vs. building

- git push なし

### 使ってみる

```bash
$ gem install building
$ building -p 3000 tcnksm/myapp
```
### What building do for you

- 専用のDockerfileの作成
- Dockerイメージのビルド
- Dockerコンテナの起動

### Dockerfile by building

```
FROM ctlc/buildstep

ADD . /app
RUN /build/builder
CMD /start web
```

カレントディレクトリのアプケーションをイメージにぶっ込んで，`build/builder`スクリプトを実行しているだけ

## Demo

```bash
$ building -p 3000 tcnksm/martini
$ docker run -d -p 3000 -e "PORT=3000" tcnksm/martini:latest
```

## Reference

- [PaaSに何が起きているのか？](http://www.infoq.com/jp/news/2014/02/paas-future)
- [Inside Dokku in 5 minutes](http://banyan.me/slides/20140116/slides.html)
- [The Start of the Age of Flynn](http://progrium.com/blog/2014/02/06/the-start-of-the-age-of-flynn/)
- [Flynn vs. Deis: The Tale of Two Docker Micro-PaaS Technologies](http://www.centurylinklabs.com/flynn-vs-deis-the-tale-of-two-docker-micro-paas-technologies/)
