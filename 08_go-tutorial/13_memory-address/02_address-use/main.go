package main

import "fmt"

// 米 转化为 码
const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Println("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meter is ", yards, " yards.")
}
