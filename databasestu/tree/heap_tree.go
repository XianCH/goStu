package tree

import (
	"fmt"
)

func main() {
	arr := []int{4, 10, 3, 5, 1}
	fmt.Println("原始数组:", arr)

	heapSort(arr)

	fmt.Println("排序后的数组:", arr)
}

// 堆排序函数
func heapSort(arr []int) {
	// 构建最大堆
	buildMaxHeap(arr)

	// 从最后一个元素开始，依次将其与堆顶元素交换并调整堆
	for i := len(arr) - 1; i >= 0; i-- {
		// 交换堆顶元素（最大值）和当前元素
		arr[0], arr[i] = arr[i], arr[0]

		// 调整堆，排除已排序的部分
		heapify(arr, i, 0)
	}
}

// 构建最大堆
func buildMaxHeap(arr []int) {
	n := len(arr)
	// 从最后一个非叶子节点开始，依次进行堆化
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

// 堆化函数，将以i为根节点的子树调整为最大堆
func heapify(arr []int, n, i int) {
	largest := i // 初始化最大值为根节点
	left := 2*i + 1
	right := 2*i + 2

	// 如果左子节点存在且大于根节点，更新最大值索引
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点存在且大于根节点，更新最大值索引
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是根节点，交换它们，并递归调整子树
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
