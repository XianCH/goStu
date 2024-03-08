package cal

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
	}

	// 测试用例2：测试负数相加
	result = Add(-2, 3)
	expected = 1
	if result != expected {
		t.Errorf("Add(-2, 3) returned %d, expected %d", result, expected)
	}
}

func BenchmarkAdd(b *testing.B) {
	// 进行基准测试
	for i := 0; i < b.N; i++ {
		// 这里我们调用 Add 函数，你可以根据需要进行替换
		Add(2, 3)
	}
}

func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Sum(nums)
	}
}

func BenchmarkGaussSum(b *testing.B) {
	n := 10
	for i := 0; i < b.N; i++ {
		GaussSum(n)
	}
}


func BenchmarkFibonacciRecursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FibonacciRecursive(20) // 计算第 20 个斐波那契数
    }
}

func BenchmarkFibonacciIterative(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FibonacciIterative(20) // 计算第 20 个斐波那契数
    }
}

