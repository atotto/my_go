## 複素数から実部と虚部を取り出したい

builtinのrealとimagを使いましょう:

```go
	c := complex(1, 2)
	r := real(c)
	i := imag(c)

	fmt.Printf("(%g+%gi)\n", r, i)

	// Output:
	// (1+2i)
```

http://play.golang.org/p/tRGjb6AtCK
