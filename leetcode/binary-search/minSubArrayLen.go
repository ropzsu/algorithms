/* https://leetcode.com/problems/minimum-size-subarray-sum/description/
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

For example, given the array [2,3,1,2,4,3] and s = 7,
the subarray [4,3] has the minimal length under the problem constraint.

More practice:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
*/

package leetcode

// Two Pointers
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLen, sum := 0, nums[0]
	for l, r := 0, 0; r < len(nums); {
		if sum >= s {
			if tmp := r - l + 1; tmp < minLen || minLen == 0 {
				minLen = tmp
			}
			sum -= nums[l]
			l++
		} else if r++; r < len(nums) {
			sum += nums[r]
		}
	}
	return minLen
}
