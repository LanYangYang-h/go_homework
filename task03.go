package main

import "fmt"

// plusOne 给用数组表示的数字加一
func plusOne(digits []int) []int {
	n := len(digits)
	carry := 1 // 初始进位为 1，因为我们要加一

	// 从数组末尾开始向前遍历
	for i := n - 1; i >= 0; i-- {
		sum := digits[i] + carry

		digits[i] = sum % 10 // 当前位的结果
		carry = sum / 10     // 更新进位

		// 如果没有进位了，可以提前退出循环
		if carry == 0 {
			break
		}
	}

	// 遍历结束后，如果还有进位，说明需要在数组开头添加一个新的高位
	if carry > 0 {
		// 创建一个新的数组，长度为 n+1
		newDigits := make([]int, n+1)
		newDigits[0] = carry // 新的最高位是进位
		// 将原数组的内容复制到新数组的后面
		copy(newDigits[1:], digits)
		return newDigits
	}

	// 没有进位，直接返回修改后的原数组
	return digits
}

func main() {
	// 测试用例
	testCases := [][]int{
		{1, 2, 3},    // 预期输出: [1 2 4]
		{4, 3, 2, 1}, // 预期输出: [4 3 2 2]
		{9},          // 预期输出: [1 0]
		{9, 9, 9},    // 预期输出: [1 0 0 0]
		{0},          // 预期输出: [1]
		{1, 9, 9},    // 预期输出: [2 0 0]
	}

	for _, tc := range testCases {
		fmt.Printf("输入: %v\n", tc)
		result := plusOne(tc)
		fmt.Printf("输出: %v\n\n", result)
	}
}
