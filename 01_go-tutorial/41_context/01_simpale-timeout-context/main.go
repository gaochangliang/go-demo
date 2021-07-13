package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx := context.Background()
	ctx, _ := context.WithTimeout(parentCtx, 100*time.Millisecond)

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

/*
通过调用标准库 context.WithTimeout 方法针对 parentCtx 变量设置了超时时间，
并在随后调用 select-case 进行 context.Done 方法的监听，最后由于达到截止时间。
因此逻辑上 select 走到了 context.Err 的 case 分支，
最终输出 context deadline exceeded。

*/

/*
WithCancel：基于父级 context，创建一个可以取消的新 context。// WithCancel返回一个父类的副本，一旦父类的Done通道被关闭，该副本就被关闭。
// parent.Done被关闭或取消被调用。

WithDeadline：
基于父级 context，创建一个具有截止时间（Deadline）的新 context

WithTimeout：基于父级 context，创建一个具有超时时间（Timeout）的新 context。
WithTimeout返回一个父类的副本，该副本的Done通道在以下情况下被关闭
parent.Done被关闭，cancel被调用，或者超时结束。新的
Context的Deadline是now+timeout和父类的Deadline中较早的一个，如果有的话。
任何。如果定时器仍在运行，取消函数会释放其资源。

WithValue返回一个父类的副本，其Value方法返回key的val。

Background：创建一个空的 context，一般常用于作为根的父级 context。
TODO：创建一个空的 context，一般用于未确定时的声明使用。
WithValue：基于某个 context 创建并存储对应的上下文信息。
*/

/*
Deadline：获取当前 context 的截止时间。
Done：获取一个只读的 channel，类型为结构体。可用于识别当前 channel 是否已经被关闭，其原因可能是到期，也可能是被取消了。关于 Done 方法，你必须要记住的知识点就是：如果 Done 没有被 close，Err 方法返回 nil；如果 Done 被 close，Err 方法会返回 Done 被 close 的原因
Err：获取当前 context 被关闭的原因。
Value：获取当前 context 对应所存储的上下文信息
*/
