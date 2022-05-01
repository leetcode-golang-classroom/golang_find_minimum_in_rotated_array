package sol

func findMin(nums []int) int {
	nLen := len(nums)
	if nLen == 1 {
		return nums[0]
	}
	left := 0
	right := nLen - 1
	for left <= right {
		mid := (left + right) / 2
		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if mid < nLen-1 && nums[mid+1] < nums[mid] {
			return nums[mid+1]
		}
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= nLen {
		left = 0
	}
	return nums[left]
}
