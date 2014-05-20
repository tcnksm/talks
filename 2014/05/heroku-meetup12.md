Heroku Meetup #12
====

http://herokujp.doorkeeper.jp/events/10902

# Deploy Go Web Application through wercker

# TL;DR

Go言語によるWeb ApplicationをHerokuにデプロイ
（使えるツールやフレームワークの紹介）

# Marniti

Go製の軽量Webフレームワーク

- Sinatraやexpressの影響
- Sinatraと同様にMiddlewareでauthやsession等の機能を追加できる

```ruby
require 'sinatra'

get '/' do
  "Hello!"
end
```

```go
package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello!"
  })
  m.Run()
}
```

# Deployの準備

- 依存関係の解決
- Procfileの準備
- Buildpack

# Godep

- Goパッケージの依存関係を管理
- RubyのBundler

```bash
$ go get github.com/kr/godep
$ godep save
```

依存関係のリストがGodeps/Godeps.jsonに書き込まれる．

# Procfile

```bash
$ echo "web: $(basename `pwd`)" > Procfile
```

（Goはディレクトリ名がバイナリ名になる）


# Go Heroku Buildpack

- HerokuがサポートするRuby，Clojure，Nodeは内部でそれぞれのBuildpackを利用してビルドする
- `git push`されたアプリケーションを判断して，専用のbuildpackを利用してアプリケーションをビルドする．
- 3rd partyのものも利用できる．

```bash
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git
```

## Deploy

あとはいつも通りにデプロイ

```bash
$ git push heroku master
```

## Wercker（ワーカー）

- 無料で使えるCI as a Service
- プライベートレポジトリでも無料で使える
- Githubへのpushを契機にテストが走る

-> テストが通ったらHerokuにデプロイするようにする

## Golang App to Wercker

- Githubのレポジトリを登録すると，自動で`wercker.yml`が作成される
    - Go言語も対応
- テストが通ったらHerokuにデプロイするようにする
    - HerokuのApp KeyとApplicationを登録
    - 以下を追記
    
```bash
deploy:
  steps:
      - heroku-deploy
```

- https://app.wercker.com/#deploystep/535687bd6472c2051c0019ec

## Heroku on Docker

- Heroku on Docker
    - HerokuのBuildpackを使うアプリケーション用のDockerコンテナを作成
- Dokkuよりシンプル
- Goのビルドパックはあらかじめ存在

```bash
$ gem install building
$ building -p 3000 tcnksm/martini
```

```bash
Step 2 : RUN /build/builder
 ---> Running in 41e9c84f2ab1
 Go app detected
 -----> Installing go1.2.1... done
 -----> Running: godep go install -tags heroku ./...
 -----> Discovering process types
 Procfile declares types -> web
```

```bash
$ docker build -t tcnksm/martini:latest .
$ docker run -d -p 3000 -e "PORT=3000" tcnksm/martini:latest
```

## デモ

## まとめ

- Go Web ApplicationをHerokuにdeploy
- Front系よりHTTP APIサービスの構築に使えそう
    - e.g., logspout（Dockerコンテナ用のlogルーティングツール）
-     

