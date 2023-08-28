package problems

func TwoSum(nums []int, target int) []int {
	for index, element := range nums {
		for sub, subel := range nums {
			if index != sub && element+subel == target {
				return []int{index, sub}
			}
		}
	}
	return []int{}
}
