package main

import "fmt"

func main() {
	values := []string{"abc", "bcd", "efg"}
	done := make(chan bool)

	for _, value := range values {
		v := value
		go func() {
			fmt.Println("value", v)
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
/*
To bind the current value of v to each closure as it is launched,
one must modify the inner loop to create a new variable each iteration.
One way is to pass the variable as an argument to the closure.

In this example, the value of v is passed as an argument to the anonymous function.
That value is then accessible inside the function as the variable u.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines

*/

/*
/*
为了在每个闭包启动时将v的当前值与之绑定。
我们必须修改内循环，以便在每个迭代中创建一个新的变量。
一种方法是将变量作为一个参数传递给闭包。

在这个例子中，v的值被作为一个参数传递给匿名函数。
然后，该值可以在函数内部作为变量u来访问。

源于此。
https://golang.org/doc/faq#closures_and_goroutines
*/
