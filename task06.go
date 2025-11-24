package main

import "fmt"

// twoSum 哈希表法，最优解
func twoSum(nums []int, target int) []int {
	// 创建一个 map 用于存储 (数值, 下标)
	numMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算需要配对的另一个数
		complement := target - num

		// 检查 complement 是否已经在 map 中
		if j, exists := numMap[complement]; exists {
			// 如果存在，则找到了解，返回两个数的下标
			return []int{j, i}
		}

		// 如果不存在，将当前数和其下标存入 map
		numMap[num] = i
	}

	// 如果遍历完整个数组都没有找到，返回 nil
	return nil
}

func main() {
	// 测试用例
	nums := []int{2, 7, 11, 15}
	target := 13

	result := twoSum(nums, target)

	if result != nil {
		fmt.Printf("数组 %v 中，和为 %d 的两个数的下标是: %v\n", nums, target, result)
		fmt.Printf("验证: %d + %d = %d\n", nums[result[0]], nums[result[1]], target)
	} else {
		fmt.Println("未找到符合条件的两个数")
	}
}
