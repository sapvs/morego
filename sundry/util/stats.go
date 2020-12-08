package util

import "sort"

func Mean(nums []float64) float64 {
	var sum, coutn float64 = 0, (float64)(len(nums))
	for _, f := range nums {
		sum += f
	}
	return sum / coutn
}

func Median(nums []float64) float64 {
	sort.Float64s(nums)
	tot := len(nums)
	if len(nums)%2 != 0 {
		return nums[len(nums)/2+1]
	} else {
		return Mean(nums[tot/2-1 : tot/2+1])
	}

}

func Mode(nums []float64) float64 {
	sort.Float64s(nums)
	tot := len(nums)
	
	return 0.0

}
