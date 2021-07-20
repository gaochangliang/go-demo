## 引言

在Go服务器中，每个传入的请求都在自己的goroutine中处理。请求处理程序经常启动额外的goroutine来访问后端，如数据库和RPC服务。处理一个请求的一组goroutine通常需要访问请求的特定值，如最终用户的身份、授权令牌和请求的截止日期。当一个请求被取消或超时时，所有在该请求上工作的goroutines应该迅速退出，以便系统可以回收他们正在使用的任何资源。

在谷歌，我们开发了一个上下文包，它可以很容易地将请求范围的值、取消信号和最后期限跨越API边界传递给所有参与处理请求的goroutine。该包以context的形式公开提供。这篇文章描述了如何使用该包，并提供了一个完整的工作实例。

## context

The core of the `context` package is the `Context` type:

```
// 一个Context携带一个截止日期、取消信号和请求范围的值
// 跨越API的界限。它的方法对多个goroutines同时使用是安全的。
// goroutines同时使用。
type Context interface {
    // Done返回一个通道，当这个Context被取消或超时时，该通道将被关闭。
    // 或超时而关闭。
    Done() <-chan struct{}。

    // Err表示这个上下文被取消的原因，在Done通道
    // 被关闭。
    Err() error

    // Deadline返回此Context被取消的时间，如果有的话。
    Deadline() (deadline time.Time, ok bool)

    // Value返回与key相关的值，如果没有则返回nil。
    Value(key interface{}) interface{}。
}
```

Done方法返回一个通道，作为代表Context运行的函数的取消信号：当通道被关闭时，函数应该放弃它们的工作并返回。Err方法返回一个错误，表明Context被取消的原因。

接收取消信号的函数通常不是发送信号的那个。特别是，当一个父操作为子操作启动goroutines时，这些子操作不应该能够取消父操作。相反，WithCancel函数（如下所述）提供了一种取消新Context值的方法。

一个Context对于多个goroutine同时使用是安全的。代码可以将一个单一的Context传递给任意数量的goroutine，并取消该Context以向所有的goroutine发出信号。

Deadline方法允许函数决定它们是否应该开始工作；如果剩下的时间太少，可能就不值得了。代码也可以使用截止日期来设置I/O操作的超时。

Value允许一个Context携带请求范围的数据。该数据必须是安全的，可以被多个goroutine同时使用。

##  派生的context

`context`包提供了从现有的`Context`值派生新的`Context`值的函数。这些值形成一棵树：当一个`Context`被取消时，所有从它派生的Context也被取消。

`Background`是任何`Context`树的根；它永远不会被取消。

```
//Background返回一个空的Context。它永远不会被取消，没有截止日期。
//并且没有任何值。Background通常在main、init和测试中使用。
//并作为传入请求的顶级Context。

func Background() Context
```
`WithCancel`和`WithTimeout`返回派生的`Context`值，这些值可以比父`Context`更快地被取消。与传入请求相关的`Context`通常在请求处理程序返回时被取消。`WithCancel`对于使用多个副本时取消多余的请求也很有用。`WithTimeout`对于设置对后端服务器的请求的最后期限很有用。

```
// WithCancel返回一个父类的副本，一旦父类的Done通道被关闭，该副本就被关闭。
// parent.Done被关闭或取消被调用。
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// 一个CancelFunc会取消一个Context。
type CancelFunc func()

// WithTimeout返回一个parent的副本，该副本的Done通道在parent.Done关闭后立即关闭。
// parent.Done被关闭，cancel被调用，或者超时结束。新的
// Context的Deadline是now+timeout和父类的Deadline中较早的一个，如果有的话。
// 任何。如果定时器仍在运行，取消函数将释放其
// 资源。
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

`WithValue`提供了一种将请求范围内的值与上下文相关联的方法。

```
// WithValue返回一个父类的副本，其Value方法返回key的val。
func WithValue(parent Context, key interface{}, val interface{}) Context
```

### Context 原理

`Context` 的调用应该是链式的,通过`WithCancel`,`WithDeadline`,`WithTimeout`或`WithValue`派生出新的 `Context`.当父 `Context `被取消时,其派生的所有 Context 都将取消.

通过`context.WithXXX`都将返回新的` Context` 和 `CancelFunc`.调用` CancelFunc` 将取消子代,移除父代对子代的引用,并且停止所有定时器.未能调用 `CancelFunc` 将泄漏子代,直到父代被取消或定时器触发.`go vet`工具检查所有流程控制路径上使用` CancelFuncs`.

### 遵循规则

遵循以下规则,以保持包之间的接口一致,并启用静态分析工具以检查上下文传播.

1. 不要将` Contexts` 放入结构体,相反`context`应该作为第一个参数传入,命名为`ctx`. `func DoSomething（ctx context.Context,arg Arg）error { // ... use ctx ... }`
2. 即使函数允许,也不要传入`nil`的` Context`.如果不知道用哪种 `Context`,可以使用`context.TODO()`.
3. 使用`context`的`Value`相关方法只应该用于在程序和接口中传递的和请求相关的元数据,不要用它来传递一些可选的参数
4. 相同的 `Context` 可以传递给在不同的`goroutine`;`Context `是并发安全的.

## 所有方法

```
func Background() Context
func TODO() Context

func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

看到Context是一个接口,想要使用就得实现其方法.在`context`包内部已经为我们实现好了两个空的`Contex`t,可以通过调用Background()和TODO()方法获取.一般的将它们作为Context的根,往下派生.



## WithCancel 例子

WithCancel返回一个带有新的Done通道的父类副本。当返回的cancel函数被调用时，或者当父级上下文的Done通道被关闭时，返回的新的上下文的Done通道会被关闭，以先发生的为准。

取消这个上下文(cancle函数)会释放与之相关的资源，所以代码应该在这个上下文中运行的操作完成后立即调用取消。

```

```

