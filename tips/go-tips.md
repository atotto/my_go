# Golang tips

## 実装

### fmt.Println でいちいち fmt パッケージをインポートするのが面倒くさい

println 関数が組み込みで使えます。ただし、一時的な使用に留めましょう。

参考： http://golang.org/ref/spec#Bootstrapping

<blockquote class="twitter-tweet" lang="ja"><p>fmt.Println() なんかを println() で略記する方法はデバッグなどだけで使いましょう <a href="https://twitter.com/search/%23Go言語">#Go言語</a> <a href="http://t.co/yYbHLbtk" title="http://play.golang.org/p/y5XX4RDTW5">play.golang.org/p/y5XX4RDTW5</a> <a href="http://t.co/dUxtmL0I" title="http://golang.org/ref/spec#Bootstrapping">golang.org/ref/spec#Boots…</a></p>&mdash; @golangjpさん (@golangjp) <a href="https://twitter.com/golangjp/status/291473554514509825">2013年1月16日</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

### 不要な戻り値を処理したい　

_ (アンダースコア)に代入しましょう

復習: http://go-tour-jp.appspot.com/#34 

### 例外(Exception)みたいなのやりたい

Goには例外はありません。
代わりに、panicがあります。ランタイムエラーを扱うときに使います。
Javaの例外とは違い、回復を望まないような時に使いましょう。
一般的なエラーは返り値のerrorで返しましょう。

参考: http://golang.org/doc/articles/defer_panic_recover.html

ちなみに、 recover はリカバリするのが安全だという確信がなければ使わないようにすべきです。

### finallyのような動きをさせたい

defer文を使いましょう。Javaでいうところのfinallyと似た使い方もできます。

参考： http://golang.org/doc/articles/defer_panic_recover.html

## go コマンド

### example_test.go にあるサンプルを直接実行したい

標準パッケージなどによく  example_test.go が付属しています。この中の ExampleXXX 関数を直接実行するには次の方法で実行できます：

    $ go test -test.run ExampleXXX

参考: http://golang.org/cmd/go/#Description_of_testing_flags

### benchmarkを取りたい

とりあえず、参考URLみて。 

TODO: あとで詳しく書くかも？ 

参考: http://golang.org/pkg/testing/#Benchmarks

### 


## パッケージ

### 外部パッケージはどこで探せばいいのか。。。

http://godoc.org/ を活用しましょう。Web上に存在するほとんどの外部パッケージを検索できるはずです。
