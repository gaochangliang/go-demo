package main

import "fmt"

func add(n int) func(int) int {
	sum := n
	f := func(x int) int {
		var i int = 2
		sum += i * x
		return sum
	}
	return f
}

func main() {
	//Here, since f1 and f2 are both independent, they do not share a common sum
	f1 := add(10)
	f2 := add(20)
	n11 := f1(3)
	n12 := f1(6)
	n21 := f2(4)
	n22 := f2(8)

	fmt.Println("n11", n11)
	fmt.Println("n12", n12)
	fmt.Println("n21", n21)
	fmt.Println("n22", n22)
}

/*So we say "closure = function + referenced environment"
When each call to the add function returns a new instance of the closure, these instances are isolated from each other and each contains a different reference environment site at the time of the call. Unlike functions, closures can have multiple instances at runtime, and different combinations of reference environments and the same function can produce different instances
*/

/*所以我们说“闭包=函数+引用环境”

当每次调用add函数时都将返回一个新的闭包实例，这些实例之间是隔离的，分别包含调用时不同的引用环境现场。不同于函数，闭包在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例。

闭包是包含自由变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。由于自由变量包含在代码块中，所以只要闭包还被使用，那么这些自由变量以及它们引用的对象就不会被释放，要执行的代码为自由变量提供绑定的计算环境。

函数a返回函数b时，如果函数b引用了函数a的局部变量，实际的返回值就变成了函数b的函数体加一个闭包，闭包里是函数b用到的全部变量。这么说是简单，但新人也不一定能看懂，作者写的通俗易懂，对新人很友好，那个银八就是矫情，咋不吃肉粥呢
*/
