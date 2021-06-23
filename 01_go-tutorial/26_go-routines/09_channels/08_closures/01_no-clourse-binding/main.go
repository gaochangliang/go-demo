package main

import "fmt"

func main() {
	values := []string{"abc", "bcd", "efg"}
	done := make(chan bool)

	for _, value := range values {
		go func() {
			fmt.Println("value", value)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}

/*
Some confusion may arise when using closures with concurrency.

One might mistakenly expect to see a, b, c as the output.
What you'll probably see instead is c, c, c. This is because
each iteration of the loop uses the same instance of the variable v,
so each closure shares that single variable. When the closure runs,
it prints the value of v at the time fmt.Println is executed,
but v may have been modified since the goroutine was launched.
To help detect this and other problems before they happen,
run go vet

*/

/*
在使用闭包与并发时，可能会出现一些混淆。

人们可能会错误地期望看到a, b, c作为输出。
你可能看到的是c, c, c. 这是因为
循环的每个迭代都使用同一个变量v的实例。
所以每个闭包都共享这个单一的变量。当闭包运行时。
它打印出执行fmt.Println时的v的值。
但自从goroutine启动以来，v可能已经被修改了。
为了帮助在这一问题和其他问题发生之前发现它们。
运行go vet。

*/
