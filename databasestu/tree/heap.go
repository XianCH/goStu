package main

import (
    "fmt"
    "testing"
)

// heapSort 使用堆排序算法对切片进行升序排序。
func heapSort(arr []int) {
    n := len(arr)

    // 构建最大堆
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }

    // 从堆中一个一个取出元素
    for i := n - 1; i >= 0; i-- {
        // 将当前根节点（最大值）与最后一个节点交换
        arr[0], arr[i] = arr[i], arr[0]

        // 在减小的堆上重新调整堆
        heapify(arr, i, 0)
    }
}

// heapify 用于在堆中对指定节点进行堆化。
func heapify(arr []int, n, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2

    // 找到左子节点比根节点大的情况
    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    // 找到右子节点比根节点大的情况
    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    // 如果最大值不是根节点，则交换它们，并递归堆化受影响的子树
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

// 测试堆排序
func TestHeapSort(t *testing.T) {
    testCases := []struct {
        input    []int
        expected []int
    }{
        {[]int{5, 2, 9, 3, 6}, []int{2, 3, 5, 6, 9}},
        {[]int{12, 11, 13, 5, 6, 7}, []int{5, 6, 7, 11, 12, 13}},
        {[]int{5, 5, 5, 5, 5}, []int{5, 5, 5, 5, 5}},
    }

    for _, tc := range testCases {
        t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
            heapSort(tc.input)
            if !isEqual(tc.input, tc.expected) {
                t.Errorf("Expected %v, got %v", tc.expected, tc.input)
            }
        })
    }
}

// isEqual 用于比较两个切片是否相等。
func isEqual(slice1, slice2 []int) bool {
    if len(slice1) != len(slice2) {
        return false
    }
    for i := range slice1 {
        if slice1[i] != slice2[i] {
            return false
        }
    }
    return true
}

func main() {
    arr := []int{5, 2, 9, 3, 6}
    fmt.Println("Original array:", arr)
    heapSort(arr)
    fmt.Println("Sorted array:", arr)
}
