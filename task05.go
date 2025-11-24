package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 1. 按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 合并区间
	merged := make([][]int, 0)
	for _, interval := range intervals {
		// 如果结果集为空，直接加入
		if len(merged) == 0 {
			merged = append(merged, interval)
			continue
		}

		// 获取结果集最后一个区间
		last := merged[len(merged)-1]
		// 当前区间的 start <= 最后一个区间的 end，说明有重叠
		if interval[0] <= last[1] {
			// 更新最后一个区间的 end 为较大值
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			// 无重叠，直接加入
			merged = append(merged, interval)
		}
	}

	return merged
}

func main() {
	// 测试用例
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {2, 3}},
		{{1, 10}, {2, 3}, {4, 5}},
	}

	for _, intervals := range testCases {
		fmt.Printf("输入: %v\n", intervals)
		fmt.Printf("合并后: %v\n\n", merge(intervals))
	}
}
