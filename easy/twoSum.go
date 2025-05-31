package easy

func twoSum(nums []int, target int) []int {
	// Create a map to store nums and indices
	numMap := make(map[int]int)

	for i, num := range nums {
		// Calculate pair for num -  the number that, when added to num, equals target
		pair := target - num

		// Check if the pair exists in our map
		if index, ok := numMap[pair]; ok {
			return []int{index, i}
		}
		// Add the current number and its index to the map
		numMap[num] = i
	}
	return nil
}
