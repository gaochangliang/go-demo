package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(10)
	fmt.Println("\n-- Unsorted ---\n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n-- sorted ---\n\n", slice)
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

//冒泡排序只会操作相邻的两个数据。
//每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。如果不满足就让它俩互换。
//一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作。
func bubbleSort(sli []int) {
	var (
		n      = len(sli)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if sli[i] > sli[i+1] {
				sli[i], sli[i+1] = sli[i+1], sli[i]
				//一次循环过程中有交换意味着可能还需要冒泡，数据还没有排序好
				//没有交换意味着所有数据都排序好了
				//当某次冒泡操作已经没有数据交换时，说明已经达到完全有序，不用再继续执行后续的冒泡操作
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}

		n = n - 1 //每一次循环结束意味着找到一个数放在数组的最右边，最大index需要要减一，
	}
}

//原地排序算法:
//冒泡的过程只涉及相邻数据的交换操作，只需要常量级的临时空间，空间复杂度为 O(1)

//稳定的排序算法
//为了保证冒泡排序算法的稳定性，当有相邻的两个元素大小相等的时候，我们不做交换，
//相同大小的数据在排序前后不会改变顺序，所以冒泡排序是稳定的排序算法

//时间复杂度
//最好时间复杂度：要排序的数据已经是有序的了，我们只需要进行一次冒泡操作，就可以结束了，所以最好情况时间复杂度是 O(n)。
//最坏时间复杂度，要排序的数据刚好是倒序排列的，我们需要进行 n 次冒泡操作，所以最坏情况时间复杂度为 O(n2)

//平均时间复杂度
