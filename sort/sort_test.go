package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func Test(t *testing.T) {
	arr := generateRandomArray(1000000)
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	t1 := time.Now()
	sort.Ints(arrCopy)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

	t1 = time.Now()
	Sort(arr)
	t2 = time.Now()
	fmt.Println(t2.Sub(t1))
	fmt.Println(reflect.DeepEqual(arr, arrCopy))
}

func TestSort(t *testing.T) {
	arr := generateRandomArray(120)
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	//BubbleSort(arr)
	//InsertionSort(arr)
	//SelectionSort(arr)
	//QuickSort(arr, 0, len(arr)-1)
	//HeapSort(arr)
	arr = MergeSort(arr)
	fmt.Println(reflect.DeepEqual(arr, arrCopy))
}

func generateRandomArray(length int) []int {
	rand.Seed(time.Now().UnixNano()) // 用当前时间来初始化随机数种子
	arr := make([]int, length)

	// 填充数组
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(1000) // 生成 0 到 999 之间的随机整数
	}

	return arr
}
