package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

/*

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

WithDeadline返回一个父级上下文的副本，其最后期限调整为不晚于d。

如果父级上下文的最后期限已经早于d，WithDeadline(parent, d)在语义上等同于parent。

WithDeadline返回的上下文的Done通道在最后期限到期时、返回的cancel函数被调用时、或者父级上下文的Done通道被关闭时（以先发生者为准）被关闭。

取消这个上下文会释放与之相关的资源，所以代码应该在这个上下文中运行的操作完成后立即调用取消。

*/
