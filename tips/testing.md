## Tests:テストを書く

パッケージ名 `eg`
テストパッケージ名は `eg_test` とする

例 https://gist.github.com/atotto/4570162

TODO:書く

参考:

* http://atotto.hatenadiary.jp/entry/2013/01/19/112012


## Benchmarks:ベンチマークを測る

TODO:書く

## Examples:関数の使い方を書く

パッケージを作成したならば、使い方の例を書きたくなるでしょう。Goには、例を書くためのステキな仕組みがあります。

通常、 `example_test.go` に例を列挙します。以下に例の例を示します：

```
func ExampleFuncA() {
        fmt.Println("hello")
        // Output:
        // hello
}
```

これは、 `FuncA` 関数に対する例を示しています。

`// Output:` でその出力結果を提示します。
Output: 以下に書いたものがテスト仕様となりますので:

    $ go test

とすることで `// Output:` で書かれた内容と比較されます（標準出力には内容は出力されません。結果のみ出力されます）。

なお、 `// Output:` を書かなければ、Example関数は `go test` で実行されませんので注意が必要です。


例の関数の命名規則は、`Exapmle`を付け、それぞれ`F`は関数名、`T`はtype、`T_M`はtype T上のメソッド`M`を指定できます：

```
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```

ひとつの関数に複数の例を書きたい場合は、`suffix`に小文字で例の名前を入れます：

```
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
func ExampleT_M_suffix() { ... }
```

```
func Example() { ... }
```

と書くと、そのパッケージ全体としての使用例を書くことができます。


参考資料:

* http://golang.org/pkg/testing/#hdr-Examples


## テストデータを置きたいんだけど

`testdata`ディレクトリにデータを入れましょう。
名前を`testdata`としておけば、godoc上へ表示されませんので便利です。
