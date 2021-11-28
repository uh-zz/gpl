## USAGE

### ページ内リンクを取得

```
go build github.com/uh-zz/gpl/ch01/fetch.go
go build -o findlinks1 github.com/uh-zz/gpl/ch05/findlinks1/main.go

./fetch https://go.dev | ./findlinks1
```

### ページアウトラインを取得

```
go build github.com/uh-zz/gpl/ch01/fetch.go
go build -o outline github.com/uh-zz/gpl/ch05/outline/main.go

./fetch https://go.dev | ./outline
```
