package src

import "sort"

func triangleNumber(nums []int) int {
	count:=0
	sort.Ints(nums)
	length := len(nums)
	for a:=0;a<=length-2;a++{
		for b:=a+1;b<length-1;b++{
			c :=b+1
			for nums[a]+nums[b]>nums[c]{
				count++
				c++
				if c >= length{
					break
				}
			}
		}
	}
	return count
}