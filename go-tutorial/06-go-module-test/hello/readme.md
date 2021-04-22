[toc]

## 介绍

本demo中涉及两个文件 `hello.go`,`hello_test.go` 根目录 `hello`

## 1 查看环境变量
```bash
$ go env
set GO111MODULE=on
...
set GOPROXY=https://goproxy.cn,direct
...
```
**关心的变量如下**

- GO111MODULE=on 使用go module包管理工具
- GOPROXY=https://goproxy.cn,direct  设置的代理

**可以使用如下命令设置变量**

```bash
$ go env -w GO111MODULE=on

$ go env -w GOPROXY=https://goproxy.cn,direct
```

## 2  code

**code**

`hello.go`


```go
func Hello() string {
	return "Hello, world."
}
```

`hello_test.go`

```go
import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
```

## 3 go mod init

```bash
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
go: to add module requirements and sums:
        go mod tidy

$ go mod tidy


$ go test
PASS
ok      example.com/hello       0.023s

$ cat go.mod
module example.com/hello

go 1.16
```

## 4 Adding a dependency

**修改  `hello.go`**

```go
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```

**go test**

```bash
$ go test
hello.go:3:8: no required module provides package rsc.io/quote; to add it:
        go get rsc.io/quote

$  go get rsc.io/quote
go get: added rsc.io/quote v1.5.2


$ go test
PASS
ok      example.com/hello       0.021s

$ cat go.mod
module example.com/hello

go 1.16

require rsc.io/quote v1.5.2 // indirect

$ go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0


$ cat go.sum
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=


```

****

## 5 upgrade depency

```bash
$  go get rsc.io/sampler
go get: upgraded rsc.io/sampler v1.3.0 => v1.99.99


$ go test
--- FAIL: TestHello (0.00s)
    hello_test.go:8: Hello() = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Hello, world."
FAIL
exit status 1
FAIL    example.com/hello       0.022s


$ go list -m -versions rsc.io/sampler
rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99


$ go get rsc.io/sampler@v1.3.1
go get: downgraded rsc.io/sampler v1.99.99 => v1.3.1

$ go test
PASS
ok      example.com/hello       0.022s


$ cat go.mod
module example.com/hello

go 1.16

require (
        rsc.io/quote v1.5.2 // indirect
        rsc.io/sampler v1.3.1 // indirect
)
```

## 6  Adding a dependency on a new major version

`hello.go`

```
package hello

import (
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
    return quote.Hello()
}

func Proverb() string {
    return quoteV3.Concurrency()
}
```

`test_hello.go`

```
package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
```

**go test**

```bash
$ go test
hello.go:5:2: no required module provides package rsc.io/quote/v3; to add it:
        go get rsc.io/quote/v3
        

$ go get rsc.io/quote/v3
go get: added rsc.io/quote/v3 v3.1.0

$ go test
PASS
ok      example.com/hello       0.022s

```

**go.mod**

```bash
$ cat go.mod
module example.com/hello

go 1.16

require (
        rsc.io/quote v1.5.2 // indirect
        rsc.io/quote/v3 v3.1.0 // indirect
        rsc.io/sampler v1.3.1 // indirect
)
```



```bash
$  go list -m rsc.io/q...
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0

$ go doc rsc.io/quote/v3
package quote // import "rsc.io/quote/v3"

Package quote collects pithy sayings.

func Concurrency() string
func GlassV3() string
func GoV3() string
func HelloV3() string
func OptV3() string
```

update  ` hello.go`

```go
package hello

import (
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quoteV3.HelloV3()
}

func Proverb() string {
	return quoteV3.Concurrency()
}

```



```bash
$ go test
PASS
ok      example.com/hello       0.021s

```

## 7 Removing unused dependencies

```bash
$  go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.1

$ cat go.mod
module example.com/hello

go 1.16

require (
        rsc.io/quote v1.5.2 // indirect
        rsc.io/quote/v3 v3.1.0 // indirect
        rsc.io/sampler v1.3.1 // indirect
)

$ go mod tidy


$ go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.1

$ cat go.mod
module example.com/hello

go 1.16

require (
        rsc.io/quote/v3 v3.1.0
        rsc.io/sampler v1.3.1 // indirect
)
```



