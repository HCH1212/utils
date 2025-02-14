package sort

// 还是建议用官方排序
func Sort(arr []int) {
	l := len(arr)
	if l < 20 {
		BubbleSort(arr)
		return
	} else if l < 1000000 {
		QuickSort(arr, 0, l-1)
		return
	} else {
		HeapSort(arr)
	}
}

// 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break // 如果没有交换，说明数组已经有序，提前退出
		}
	}
}

// 选择排序
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

// 插入排序
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 快速排序
func QuickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		QuickSort(arr, low, p-1)
		QuickSort(arr, p+1, high)
	}
}
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 堆排序
func HeapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		// 如果堆顶已是最大值，则无需交换
		if arr[0] != arr[i] {
			arr[0], arr[i] = arr[i], arr[0]
			heapify(arr, i, 0)
		}
	}
}
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// 归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right)) // 使用合适的容量来避免多次分配内存
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)  // 将左边剩余的部分追加到结果
	result = append(result, right[j:]...) // 将右边剩余的部分追加到结果
	return result
}
