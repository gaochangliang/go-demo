package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()

	//定义规则具体参考网上文档
	c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("every 3 seconds executing")
	})

	c.Start()

	//主要用于阻塞使用
	//select {
	//case <-time.After(time.Second * 10):
	//	return
	//}

	t1 := time.NewTimer(time.Second * 3)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 3)
		}
	}
}
