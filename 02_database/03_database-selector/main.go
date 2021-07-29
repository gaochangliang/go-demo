package main

import (
	"database/03/core/store"
	"fmt"
	"time"
)

func main() {
	//这里只是简单的了解设计这种模式，字段数据没有参考意义
	db := store.NewStorer("mysql",
		store.Opt{
			ConnStr:     "127.0.0.1",
			IdleTimeout: time.Second,
		},
	)
	fmt.Println("createDB", db)
}
