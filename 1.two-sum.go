/*
 * [1] Two Sum
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (43.07%)
 * Total Accepted:    94.5K
 * Total Submissions: 219.5K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
 *
 * 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 *
 */
package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if index, ok := m[complement]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}

	return nil
}
