package main

import "fmt"

func main() {
	word := "kobe"
	letter := rune(word[0])
	fmt.Println("letter", letter)

	letter1 := rune(word[1])
	fmt.Println("letter1", letter1)

	letter2 := rune(word[2])
	fmt.Println("letter2", letter2)

	letter3 := rune(word[3])
	fmt.Println("letter1", letter3)

}

//letter4 := rune(word[4])
//fmt.Println("letter1",letter4)

//add these line will get index out of range [4] with length 4
