Heroku Meetup #12
====

http://herokujp.doorkeeper.jp/events/10902

# Deploy Go Web Application through wercker

# TL;DR

Go言語によるWeb ApplicationをHerokuにデプロイ
使えるツールやフレームワークの紹介

# Marniti

Go製の軽量Webフレームワーク

- Sinatraやexpressの影響
- Middlewareでauthやsession等の機能を追加できる

```go
package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
```

## Building

- Heroku on Docker
    - buildpackを使ったDockerfileの生成
    - ローカルプロジェクトをコンテナにADD
- Dokkuよりシンプル

```bash
$ gem install building
```
