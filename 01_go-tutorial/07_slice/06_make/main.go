package main

func main() {

	customerNumber := make([]int, 3)
	// 3 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerNumber[0] = 1
	customerNumber[1] = 100
	customerNumber[2] = 1000

	//customerNumber[3] = 10000 //err

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this
	greeting[0] = "good morning"
	greeting[1] = "hello"
	greeting[2] = "dias"
	greeting[3] = "haha" //err

}

//The index of the slice should be less than the length
