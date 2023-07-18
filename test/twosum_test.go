package test

import (
	"LeetCodeGo/problems"
	_ "LeetCodeGo/problems"
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := problems.TwoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error in Test Case 1. Expected: %v, but got: %v", expected, result)
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	result := problems.TwoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error in Test Case 2. Expected: %v, but got: %v", expected, result)
	}
}
