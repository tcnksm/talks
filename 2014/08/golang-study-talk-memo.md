How to think Go
====


非同期プログラミングの経験があるとよい

## 例外とエラー

- この条件を満たしてないと実行できないという場合に使う
- os.Stat() を見る
- fmt.Errorf()を使えばよい

スコープをどうするのか

```go
if err := hoge(); err != nil {
}

var err error
err = hoge()
if err != nil {
}
```

## 構造体
　
- 継承ではない
- 委譲


コードの再利用は？

- 合成する
- 継承せずに小さな部品に繋げる
- 親から生成するのではなく子供から結びつけるようにする

```go
func (s State) String() string {
  return s
}

type Foo struct {
  State
}

f := Foo{}
f.String() # Stateに委譲される
```

ラップする

- interface{}の定義を満たしていけばよい
- 利用するAPIから考える
- 考えるべきはAPI
- 例えば草食動物ではなく，草を食べるというinterfaceを考えてメソッドをそろえる
- あくまでAPI=できることベースで
