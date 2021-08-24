/*
Implement an algorithm that does not use additional data structures and storage space,
Given a string, return a string that is the flipped string.
Make sure the length of the string is less than or equal to 5000.
*/

package main

import "fmt"

func main() {
	str := "123456"
	res,_ := reverseString(str)
	fmt.Println("reverseString",res)
}

func reverseString(str string) (string, bool) {
	s := []rune(str)
	if len(s) > 5000 {
		return str, false
	}

	//The length of the string is 1/2 of the length,
	//and the value is assigned back and forth, minus i.
	for i := 0 ; i < len(s)/2; i++ {
		s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
	}

	return string(s),true

}
