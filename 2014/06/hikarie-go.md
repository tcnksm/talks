Distribute golang tool for multiple platform
====

at [hikarie.go](http://connpass.com/event/6579/)

## Introduction

- [tcnksm/dmux](): docker + tmux = dmux
- [tcnksm/docc](): docc open your project documentation

## TL;DR

Go言語で作ったツールを複数プラットフォームに配布する

## Why golang?

- 自分が作ったツールを使ってもらうための障壁を下げる
    - ruby ? gem ?

例えば，OSXなら

```bash
$ brew tap tcnksm/docc
$ brew install docc
```

と言いたい．

## Hashicorp

で，Go言語の主要プレーヤーであるHashicorpはあらゆるツールを
複数プラットフォーム向けに配布していて見習いたい．

e.g., serf, packer


## How ?

1. Cross-compile
1. Upload
1. Install

## Cross-Compiling

### Basic

```bash
$ GOOS=linux GOARCH=amd64 go build hikarie.go
```

### mitchellh/gox

を使うと並列でビルドできる．

```bash
$ gox
```

### packaging

```bash
for PLATFORM in $(find ./pkg); do
    PLATFORM_NAME=$(basename ${PLATFORM})
    ARCHIVE_NAME=dmux_${VERSION}_${PLATFORM_NAME}

    pushd ${PLATFORM}
    zip ${DIR}/pkg/dist/${ARCHIVE_NAME}.zip ./*
    popd
done
```

## Hosting

### Bintray

```bash
$ curl -T <FILE.EXT> \
       -utcnksm:<API_KEY> \
       https://api.bintray.com/content/tcnksm/dmux/dmux/<VERSION_NAME>/<FILE_TARGET_PATH>
```

## Installation

### Goal

```bash
$ brew tap tcnksm/docc
$ brew install docc
```

### How ?

1. homebrew-<packege>でGitHubレポジトリを作る
2. <package>.rbでformulaを作成

### Formula


```ruby
require "formula"

class Docc < Formula
  homepage "https://github.com/tcnksm/docc"
  version '0.1.1'

  if Hardware.is_64_bit?
    url "https://github.com/tcnksm/docc/releases/download/v0.1.1/docc_0.1.1_darwin_amd64.zip"
    sha1 "eaad2915415c5ceb3e3fb613420be62a856da46a" # $ shasum docc_0.1.0_darwin_amd64.zip
  else
    url "https://github.com/tcnksm/docc/releases/download/v0.1.1/docc_0.1.1_darwin_386.zip"
    sha1 "6d8558d8589d7dcf034bf76035da1868248ddccc"
  end

  depends_on :arch => :intel

  def install
    bin.install 'docc'
  end
end
```

## Cool tool ?

本家にPull Requestを投げましょう

homebrew/Library/Formula

## DEMO

```bash
$ gox
```
