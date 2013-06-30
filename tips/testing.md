# テストを書く


# 使い方の例を書く

パッケージを作成したならば、使い方の例を書きたくなるでしょう。Goには、例を書くためのステキな仕組みがあります。

通常、 `example_test.go` に例を列挙します。以下に例の例を示します：

```
func ExampleFuncA() {
        fmt.Println("hello")
        // Output: hello
}
```

これは、 `FuncA` 関数に対する例を示すことになります。

`// Output:` でその出力結果を提示できます。
Output: 以下に書いたものは、それそのものがテスト仕様となります。

    $ go test -run Example  

で Example* 全てを実行し、 Output: で書かれたものと比較されます（標準出力には出力されません）。

### 参考資料:

* http://golang.org/pkg/testing/#hdr-Examples
