package main

import (
	"fmt"
)

func main() {
	in := gen()
	out := fac(in)
	for v := range out {
		fmt.Println(v)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for k := 3; k < 13; k++ {
				out <- k
			}
		}
		close(out)
	}()
	return out
}

func fac(c chan int) chan int {
	out := make(chan int)
	go func() {
		for k := range c {
			out <- factorial(k)
		}
		close(c)
	}()
	return out

}

func factorial(c int) int {
	value := 1
	for i := 2; i <= c; i++ {
		value *= i
	}
	return value
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/

/*
挑战 #1:
-- 改变上面的代码，使之同时并行地执行100个阶乘计算。
--使用 "流水线 "模式来完成这个任务

发帖讨论。
-- 在构建解决方案时，你对并发代码的工作有什么认识？
--例如，你有什么见解，帮助你理解与并发性工作？
-- 在这里发表你的答案，并阅读其他答案：https://goo.gl/uJa99G
*/
