package Array

func removeElement(nums []int, val int)int{
	left := 0
	for _, val := range nums{
		if val != nums[left]{
			nums[left] = val
			left++
		}
	}
	return left
}

func moveZero(nums []int) []int {
	right, left, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	return nums
}

func reverseStrings(s []byte) []byte {
	for left, right := 0, len(s) -1; left < right ; left++ {
		s[left], s[right] = s[right], s[left]
	}
	return s
}

func twoSum(num []int, target int)(int, int){
	left, right := 0, len(num)
	for left < right {
		sum := num[left] + num[right]
		if sum < target{
			left++
		} else if sum > target {
			right++
		} else {
			break
		}
	}
	return left, right
}