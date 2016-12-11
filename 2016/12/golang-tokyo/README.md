# golang.tokyo #2

25 min talk and 5 min questions

- Introduction
- Table Driven Test
    - シンプルな例
    - Table Drivenテストは強力
    - 簡単にテストケースを追加できる
    - e.g., http.Handlerも落とし込める Done
    - e.g., コンパイラのテストも落とし込める Done
    - テストしやすいコードとは
        - TableDriveTestに落とし込めるコードである
        - どんな複雑なアプリケーションもシンプルなTable-Drivenの粒度/もしくは書き方に落とし込めることが大切
- テストしにくいコードとは?
    - 入力がコントロールできていない（変更できない、変更しにくい）
        - グローバル変数
            - テストに限らず難しいのはこの値がどこから来たのか・どこで設定されてるのかが追えないとしんどい
            - Gopherは暗黙ではなく明示的であることを嗜好する（すべてをgodefで追いたい）           
            - シンプルな例（Config pathを避ける） Done
            - テストないで上書きすれば良い．テストを並列に動かした場合にRaceが起こる可能性がある
            - グローバル変数はどこで使うべきか（明確な思想ーDefault値、そしてmainでのみ使う）
            - 似た話でflagはどうか? flagをどこに定義するか問題
        - プログラム外からの入力（flag, env, stdin, args）
            - ユーザの入力をテストしたい https://github.com/tcnksm/go-input // Done
            - 環境変数（12factor） // Done
            - const, flag, envはなるべくmain関数（もしくはそれに順ずるもの）でのみ使う
    - 出力を取り出せない/取り出しにくい（io.Writerをちゃんと使おう）
        - void的関数のテストはしんどい mainのみをvoidにする!
        - ファイルへの出力（Save config）
        - Stdout/Stderrは暗黙的な設定である（切り替えられる！）・やりすぎても意味ない（コンソールへの出力が機能であればやるべき)
        - logは依存である   
        - ステータスコードをテストしたい
    - 依存を抽象化できていない（外界との境界をinterfaceにする）
        - DBを準備しないと...
        - Synmetric API testing
- 設計パターンのベストプラクティス
    - const env flag などをまとめて
    - inerfaceを含めて
- デザインの指標
    - 依存の流れを明確にしよう．各依存がしっかりテストされているのが大切
    - 設定を変更できない振る舞いをテストするのはしんどい．毎回同じでも設定可能にする
- Conclusion



TODO

- [[https://testing.googleblog.com/2010/12/test-sizes.html][Test Sizes]]
- [[https://speakerdeck.com/twada/php-conference-2016][PHP7で堅牢なコードを書く - 例外処理、表明プログラミング、契約による設計]]


* What's new in golang testing package?

* Go1.7

Released on August 2016.

- [[][Subtests / Sub benchmarks]]
- [[http://deeeet.com/writing/2016/08/02/go1_7-subtest/][Go1.7のSubtestsとSub-benchmarks]]

* Go1.8

[[https://golang.org/dl/#go1.8beta1][Go1.8 beta1]] was released on Dec (expected to be released in February 2017).

- [[https://beta.golang.org/doc/go1.8#more_context][T.Context]]

