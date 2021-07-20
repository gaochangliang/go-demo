package main

import (
	"context"
	"fmt"
)

type favContextKey string

func main() {

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}

/*

func WithValue(parent Context, key, val interface{}) Context

WithValue返回一个父类的副本，其中与key相关的值是val。

使用context Values只能用于传输进程和API的请求范围的数据，不能用于向函数传递可选参数。

提供的键必须是可比较的，并且不应该是字符串类型或任何其他内置类型，

以避免使用上下文的包之间发生碰撞。WithValue的用户应该为键定义自己的类型。

为了避免在分配给一个接口{}时进行分配，上下文键通常具有具体的结构{}类型。或者，导出的上下文关键变量的静态类型应该是一个指针或接口。

*/
