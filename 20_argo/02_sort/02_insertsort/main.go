//insertsort

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(10)
	fmt.Println("\n-- Unsorted ---\n\n", slice)
	insertSort(slice)
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

func insertSort(sli []int) {
	//想不到变量值可以用n代替
	var n = len(sli)
	for i := 0; i < n; i++ {
		j := i
		for j > 0 {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
			j = j - 1
		}
	}
}

//插入排序算法:
//在插入排序算法中，我们假设列表的某一部分已经被排序，而另一部分仍然没有排序。
//数组中的数据分为两个区间，已排序区间和未排序区间。
//初始已排序区间只有一个元素，就是数组的第一个元素。
//插入算法的核心思想是取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。
//重复这个过程，直到未排序区间中元素为空，算法结束

//原地排序算法:
//插入排序算法的运行并不需要额外的存储空间，所以空间复杂度是 O(1)

//稳定的排序算法:
//我们可以选择将后面出现的元素，插入到前面出现元素的后面，这样就可以保持原有的前后顺序不变，所以插入排序是稳定的排序算法

//插入排序的时间复杂度
//最好时间复杂度:如果要排序的数据已经是有序的，我们并不需要搬移任何数据。
//如果我们从尾到头在有序数据组里面查找插入位置，每次只需要比较一个数据就能确定插入的位置。O(n)

//如果数组是倒序的，每次插入都相当于在数组的第一个位置插入新的数据，
//所以需要移动大量的数据，所以最坏情况时间复杂度为 O(n2)。
