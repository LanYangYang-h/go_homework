package main

import "fmt"

// removeDuplicates 从有序数组中原地删除重复项，并返回新长度
func removeDuplicates(nums []int) int {
	// 1. 处理边界情况：如果数组为空，返回 0
	if len(nums) == 0 {
		return 0
	}

	// 2. 初始化慢指针 i
	i := 0

	// 3. 快指针 j 从 1 开始遍历数组
	for j := 1; j < len(nums); j++ {
		// 4. 当快指针找到与慢指针不同的元素时
		if nums[j] != nums[i] {
			// 慢指针向后移动一位，准备接收新元素
			i++
			// 将快指针指向的新元素赋值给慢指针的新位置
			nums[i] = nums[j]
		}
		// 如果元素相同，快指针继续向后走，慢指针不动
	}

	// 5. 新数组的长度是慢指针的索引 + 1
	return i + 1
}

func main() {
	// 测试用例
	testCases := [][]int{
		{1, 1, 2, 4, 5, 5, 6},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{},
		{5},
		{2, 2, 2, 2},
	}

	for _, nums := range testCases {
		fmt.Printf("原始数组: %v\n", nums)
		length := removeDuplicates(nums)
		fmt.Printf("删除重复项后长度: %d\n", length)
		fmt.Printf("删除重复项后的数组前 %d 个元素: %v\n", length, nums[:length])
		fmt.Println("---------------------------------")
	}
}
