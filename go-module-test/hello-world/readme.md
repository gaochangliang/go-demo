## Steps
代码编写完成后
```bash
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
go: to add module requirements and sums:
        go mod tidy
             
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
         
$ go test
PASS
ok      example.com/hello       0.014s     
```