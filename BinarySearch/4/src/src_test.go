package src_test

import "testing"

func minNum(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func findKthNum(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return minNum(nums1[0], nums2[0])
	}
	i := minNum(len(nums1), k/2)
	j := minNum(len(nums2), k/2)
	if nums1[i-1] > nums2[j-1] {
		return findKthNum(nums1, nums2[j:], k-j)
	} else {
		return findKthNum(nums2, nums1[i:], k-i)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	return (float64)(findKthNum(nums1, nums2, (len1+len2+1)/2)+
		findKthNum(nums1, nums2, (len1+len2+2)/2)) / 2.0
}

func TestFunction(t *testing.T) {
	arr1 := []int{2, 2, 2, 2}
	arr2 := []int{2, 2, 2}
	t.Log(findMedianSortedArrays(arr1, arr2))
}
