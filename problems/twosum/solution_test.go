package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	// Test Case 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0, 1}
	result1 := TwoSum(nums1, target1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Error in Test Case 1. Expected: %v, but got: %v", expected1, result1)
	}
}

func TestTwoSum2(t *testing.T) {
	// Test Case 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	expected2 := []int{1, 2}
	result2 := TwoSum(nums2, target2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Error in Test Case 2. Expected: %v, but got: %v", expected2, result2)
	}
}
