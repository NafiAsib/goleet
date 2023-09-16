package leetcode

func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= target {
			target = i
		}
	}

	return target == 0
}
