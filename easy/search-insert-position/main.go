package main

func searchInsert(nums []int, target int) int {
	var length int = len(nums)
	
	if length < 1 || target < nums[0] {
		return 0
	}
	
	if target > nums[length - 1]{
		return length
	}

	var lowerBoundary int = 0
	var upperBoundary int = length - 1

	for {
		offset := (upperBoundary - lowerBoundary)/2
		centerValue := nums[lowerBoundary + offset]
		switch {
		case centerValue == target:
			return lowerBoundary + offset
		case target == nums[lowerBoundary]:
			return lowerBoundary
		case target == nums[upperBoundary]:
			return upperBoundary
		case upperBoundary == lowerBoundary || offset == 0:
			return lowerBoundary + 1
		case target > centerValue:
			lowerBoundary += offset
		case target < centerValue && target > nums[lowerBoundary]:
			upperBoundary = lowerBoundary + offset
		case target < nums[lowerBoundary]:
			return lowerBoundary
		}
	}

	return -1
}

