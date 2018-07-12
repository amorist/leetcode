//Package twosum https://leetcode.com/problems/two-sum/description/
package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	var ans []int
	for i := 0; i < len(nums); i++ {
		gap := target - nums[i]
		if _, ok := m[gap]; ok && m[gap] != i {
			ans = append(ans, i)
			ans = append(ans, m[gap])
			break
		}
	}
	return ans
}
