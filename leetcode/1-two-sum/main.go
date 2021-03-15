package main

import "fmt"

func main() {
	var target = 10
	var sli = []int{1, 2, 5, 8, 0, 10}
	fmt.Println(TwoSum(sli, target))
}

/*
假设map有值
遍历循环后没找到，就依次将数组的value作为map的k, 数组的index 作为map的value
接着根据value := target - sli[i]找key
*/

func TwoSum(sli []int, target int) []int {
	if len(sli) <= 1 {
		return nil
	}
	m := make(map[int]int, len(sli))
	for i := 0; i < len(sli); i++ {
		value := target - sli[i]
		if _, ok := m[value]; ok {
			return []int{m[value], i}
		}
		m[sli[i]] = i
	}
	return nil
}
