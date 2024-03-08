package cal

func Add(a, b int) int {
	return a + b
}

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func GaussSum(n int) int {
	return n * (n + 1) / 2
} // 递归方式计算斐波那契数列的第 n 个数

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}
// 迭代方式计算斐波那契数列的第 n 个数
func FibonacciIterative(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
