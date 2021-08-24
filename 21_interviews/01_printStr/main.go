/*
Alternate printing of numbers and words

Use two goroutines to print the sequence alternately,
with one goroutine printing numbers and the other goroutine printing words.
One goroutine prints the numbers and the other goroutine prints the words, t
he final result is as follows.

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

*/

package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	strCh, numberCh := make(chan bool), make(chan bool)

	wait := sync.WaitGroup{}

	go func() {
		i := 1
		// notice the use of for and select
		for {
			select {
			case <-numberCh:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				strCh <- true
			}
		}

	}()

	wait.Add(1)
	go func(group *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYX"
		//index can use i
		i := 0
		for {
			select {
			case <-strCh:
				if i >= strings.Count(str,"") -1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str,"") {
					i = 0
				}
				fmt.Print(str[i:i+1])
				i++
				numberCh <- true
			}
		}
	}(&wait)
	numberCh <-true
	wait.Wait()

}

//The transaction print string usually has two global channels through which both sides determine the order of printing
