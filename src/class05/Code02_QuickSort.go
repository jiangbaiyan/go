package class05

import (
	"math/rand"
)

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	process2(arr, 0, len(arr)-1)
}

func process2(arr []int, L int, R int) {
	if L >= R {
		return
	}
	// 快排3.0 首先随机选一个值与数组最后一个元素(划分值)交换
	randIndex := L + rand.Intn(R-L+1)
	arr[randIndex], arr[R] = arr[R], arr[randIndex]
	// eqL, eqR为等于区的左右边界
	eqL, eqR := netherlandsFlagFlag(arr, L, R)
	// 处理左边
	process2(arr, L, eqL-1)
	// 处理右边
	process2(arr, eqR+1, R)
}

// netherlandFlag 在L...R上做荷兰国旗划分
// arr[R]为划分值
// <arr[R]在左边, =arr[R]在中间 >arr[R]在右边 返回=区的左右边界
func netherlandsFlagFlag(arr []int, L int, R int) (int, int) {
	if L > R {
		return -1, -1
	}
	if L == R {
		return L, R
	}
	less := -1 // 小于区
	more := R  // 大于区 (注意这里不是R + 1, 因为最后一个数为划分值)
	index := L // 当前数
	for index < more {
		if arr[index] < arr[R] {
			arr[less+1], arr[index] = arr[index], arr[less+1]
			less++
			index++
		} else if arr[index] > arr[R] {
			arr[more-1], arr[index] = arr[index], arr[more-1]
			more--
			// 注意这里index不++, 因为右边换过来的数还没看过, 需要停住看一下
		} else {
			index++
		}
	}
	arr[R], arr[more] = arr[more], arr[R]
	return less + 1, more
}
