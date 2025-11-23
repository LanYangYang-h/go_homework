package main

import "fmt"

func isvalid(s string) bool {
	rightcase := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if left, ok := rightcase[c]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}

	}
	return len(stack) == 0
}

func main() {

	testCases := []struct {
		s        string
		expected bool
	}{
		{"()", true},     // 基础有效
		{"()[]{}", true}, // 多种括号混合有效
		{"(]", false},    // 括号类型不匹配
		{"([)]", false},  // 括号顺序不匹配
		{"{[]}", true},   // 嵌套匹配有效
		{"(", false},     // 左括号多余
		{")", false},     // 右括号多余
		{"", true},       // 空字符串（题目未明确，但通常视为有效）
	}

	// 执行测试并输出结果
	for _, tc := range testCases {
		result := isvalid(tc.s)
		fmt.Printf("字符串 \"%s\" 是否有效：%t（预期：%t）\n", tc.s, result, tc.expected)
	}
}
