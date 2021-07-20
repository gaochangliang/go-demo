package main

import (
	"context"
	"fmt"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	//The reason why the previous concurrent process does not care about the context, because it will automatically end, here is a loop, so you need to use the context
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	//返回一个通道，但是在一直写数据，所以这个通道没有发生错误，
	//发生错误的情况是当读的时候确没有人写数据
	//当n值等于5的时候，执行cancle函数会关闭WithCancel生成的上下文的通道，从而释放相应的资源
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

/*

WithCancel返回一个带有新的Done通道的父类副本。当返回的cancel函数被调用时，
或者当父级上下文的Done通道被关闭时，返回的新的上下文的Done通道会被关闭，以先发生的为准。

取消这个上下文(cancle函数)会释放与之相关的资源，所以代码应该在这个上下文中运行的操作完成后立即调用取消。

CancelFunc告诉一个操作要放弃它的工作。

CancelFunc并不等待工作的停止。

一个CancelFunc可以被多个goroutine同时调用。在第一次调用后，对CancelFunc的后续调用不做任何事情。

*/

/*
Background返回一个非零的空的Context。

它永远不会被取消，没有值，也没有截止日期。它通常被主函数、初始化和测试使用，并作为传入请求的顶层上下文。

TODO返回一个非零的空的Context。

当不清楚应该使用哪个Context或者它还不可用时（因为周围的函数还没有被扩展到接受Context参数），代码应该使用context.TODO。
*/
