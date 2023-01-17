# go-kenshiro

あたたたたたたたたたた

## Usage

```
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ch := Kenshiro(ctx)
for i := 0; i < 10; i++ {
	fmt.Println(<-ch)
}
```

## Installation

```
go get github.com/mattn/go-kenshiro
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
