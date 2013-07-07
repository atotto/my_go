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

### JavaでいうところのStringBufferなのが欲しい

byte.BufferのWriteStringを使いましょう。

    package main
    
    import "bytes"
    
    func main() {
        var buffer bytes.Buffer
    
        for i := 0; i < 1000; i++ {
            buffer.WriteString("a")
        }
    
        println(buffer.String())
    }

参考: http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go


### CPUフル活用したい

Goの並行性を最大限に活かすためにはGOMAXPROCSを設定してあげる必要があります。

一番簡単な例は：

    runtime.GOMAXPROCS(runtime.NumCPU())

をmain関数へ記述することで、そのマシンのCPU数を設定します。

参考: http://golang.org/pkg/runtime/#GOMAXPROCS


## go コマンド

### example_test.go にあるサンプルを実行してみたい

標準パッケージなどによく  example_test.go が付属しています。この中の ExampleXXX 関数を直接実行するには次の方法で実行できます：

    $ go test -test.run ExampleXXX

ただし、テスト実行するだけで、ExampleXXXに書かれている標準出力はコンソールに出力されません。出力結果は `Output:` 以下に書かれている内容です。もし `Output:` の内容と異なれば、そのテストは失敗します。

参考:

* http://golang.org/cmd/go/#Description_of_testing_flags
* http://golang.org/pkg/testing/#hdr-Examples

### benchmarkを取りたい

とりあえず、参考URLみて。 
参考: http://golang.org/pkg/testing/#Benchmarks

例えば、BenchXXXなベンチマーク関数をメモリベンチマークと共にとる場合は：

    $ go test -test.bench Bench -test.benchmem

### test用の実行ファイルをつくっておきたい

通常、Goのテストは：

    $ go test

で実行しますが、`-c`オプションをつけることで実行ファイルを生成することができます：

    $ go test -c

で`<package_name>.test`というファイルが生成され、実行することでテストが走ります。

参考：

* http://golang.org/cmd/go/#hdr-Test_packages

## パッケージ

### 便利そうな外部パッケージはどこで探せばいいのか。。。

http://godoc.org/ を活用しましょう。Web上に存在するほとんどの外部パッケージを検索できるはずです。


### Doxygenみたいに自分のパッケージもWebで見たい。

$GOPATH以下に配置すれば、 `godoc` コマンドを使って：

    $ godoc -http=:8000

とすれば自分のマシン上にgoのドキュメントページ(http://localhost:8000/pkg)を開設できます。

## 学習

### テストコードの書き方がわからん。

$GOROOT/src/pkg 以下を参考にしましょう。

`*_test.go` は全てテストコードです。
