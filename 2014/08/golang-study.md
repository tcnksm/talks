Distirbute tool built by Golang
====

at [Golang勉強会](http://connpass.com/event/7814/)

## Introduction

- [tcnksm/ghr](): Easy to ship your project on Github to your user
- [tcnkms/cli-init](): The easy way to start building Golang command-line application
- [tcnksm/dmux](): docker + tmux = dmux
- [tcnksm/docc](): docc open your project documentation

## YAPC::Asia

[コマンドラインツールについて語るときに僕の語ること - YAPC::Asia Tokyo 2014](http://yapcasia.org/2014/talk/show/b49cc53a-027b-11e4-9357-07b16aeab6a4)

## TL;DR

Go言語で作ったツールを複数プラットフォームに向けて配布する方法

## Agenda

- Go以前/以後の世界
    - http://mitchellh.com/abandoning-rubygems
    - http://blog.codegangsta.io/blog/2013/07/21/creating-cli-applications-in-go/
- Cross-Compile
    - davecheney/golang-crosscompile
    - laher/goxc
    - mitchellh/gox
- Release
    - Go言語のツールをクロスコンパイルしてGithubにリリースする
    - Wercker で Go のプロジェクトをクロスコンパイルし、GitHub にリリースする
    - Go言語のビルド生活を drone.ioで幸せに暮らす
    - tcnksm/ghr
- Installation
    - OSXはHomebrew，Debina系に向けてdebパッケージ，Red Hat系に向けてRPMパッケージ
    - HerokuとGithubを使った統一的なツール配布
        - Terraform, http://www.terraform.io/
    - OSX: Homebrew
- Demo (ghr)
    - Cross-Compile
        - `script/package.sh`
    - Release
        - `ghr v0.1.2 dist`
    - Install
        - `terraform apply`


## Cross-Compiling

### Basic

```bash
$ GOOS=linux GOARCH=amd64 go build hikarie.go
```

### Tools

- davecheney/golang-crosscompile
    - bash実装
- laher/goxc
    - 多機能 (クロスコンパイル，パッケージング，bintrayへのアップロード）
- mitchellh/gox
    - シンプル，並列実行

```bash
gox \
    -os="darwin linux windows" \
    -arch="386 amd64" \
    -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}} "
```

- 出力先をtext/template形式で指定できる

## Release

作成したパッケージングをリリースする

### Bintray

```bash
$ curl -T <FILE.EXT> \
       -utcnksm:<API_KEY> \
       https://api.bintray.com/content/tcnksm/dmux/dmux/<VERSION_NAME>/<FILE_TARGET_PATH>
```

Hashicorpスタイル

### Github Release

```bash
$ curl --fail -X POST https://uploads.github.com/repos/${OWNER}/${REPO}/releases/${RELEASE_ID}/assets?name=${ARCHIVE_NAME} \
    -H "Accept: application/vnd.github.v3+json" \
    -H "Authorization: token ${GITHUB_TOKEN}" \
    -H "Content-Type: ${CONTENT_TYPE}" \
    --data-binary @"${ARCHIVE}"
```

### CI as a Service

テスト -> Cross-compile -> リリースを同時に実行

- Wercker で Go のプロジェクトをクロスコンパイルし、GitHub にリリースする
- Go言語のビルド生活を drone.ioで幸せに暮らす

### ghr

```bash
$ ghr v0.1.0 pkg/dist/v0.1.0
--> Uploading: pkg/dist/v0.1.0/ghr_0.1.0_darwin_386.zip
--> Uploading: pkg/dist/v0.1.0/ghr_0.1.0_darwin_amd64.zip
--> Uploading: pkg/dist/v0.1.0/ghr_0.1.0_linux_386.zip
--> Uploading: pkg/dist/v0.1.0/ghr_0.1.0_linux_amd64.zip
--> Uploading: pkg/dist/v0.1.0/ghr_0.1.0_windows_386.zip
--> Uploading: pkg/dist/v0.1.0/ghr_0.1.0_windows_amd64.zip
```


## Installation

自分のプラットフォームに合ったパッケージをダウンロード／展開してPATHの通ったところに置いてもらう

-> もっとツールを使い初めてもらうまでの敷居を下げておきたい


### Homebrew


```bash
$ brew tap tcnksm/ghr
$ brew install ghr
```

```ruby
require "formula"

class Ghr < Formula
  homepage "https://github.com/tcnksm/ghr"
  version 'v0.1.1'

  if Hardware.is_64_bit?
    url "https://github.com/tcnksm/ghr/releases/download/v0.1.1/ghr_v0.1.1_darwin_amd64.zip"
    sha1 "e5c793001f004b670df77cb57bc033c89201b485"
  else
    url "https://github.com/tcnksm/ghr/releases/download/v0.1.1/ghr_v0.1.1_darwin_386.zip"
    sha1 "642b827ae3d3b52dbc2060630e76604b074e139d"
  end

  depends_on :arch => :intel

  def install
    bin.install 'ghr'
  end

  def caveats
    msg = <<-'EOF'
 ________  ___  ___  ________
|\   ____\|\  \|\  \|\   __  \
\ \  \___|\ \  \\\  \ \  \|\  \
 \ \  \  __\ \   __  \ \   _  _\
  \ \  \|\  \ \  \ \  \ \  \\  \|
   \ \_______\ \__\ \__\ \__\\ _\
    \|_______|\|__|\|__|\|__|\|__|

EOF
  end
end
```

### Heroku

```bash
$ L=/usr/local/bin/hk && curl -sL -A "`uname -sp`" https://hk.heroku.com/hk.gz | zcat >$L && chmod +x $L
```

具体的な動作

1. ユーザがHerokuアプリに対してリクエストを投げる
1. アプリはリクエストに基づきプラットフォームを判定し，それに合ったGithub Release上のパッケージへのリダイレクトを返す
1. ユーザはプラットフォームに合ったパッケージをGithub Releaseから得る

ワンライナー

1. 環境変数LにインストールしたいPATHを指定する
1. curlでは，-Lオプションでリクエストで30Xの場合にリダイレクトするようにし，-Aでユーザエージェントを指定する
1. zcatでzipを展開してLに吐き出す


## DEMO

- Cross-Compile
    - `script/package.sh`
- Release
    - `ghr v0.1.2 dist`
- Install
    - `terraform apply`
