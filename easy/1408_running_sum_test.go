/*
 * Created: 2021-02-13 16:40:16
 * Author : Win-Man
 * Email : gang.shen0423@gmail.com
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Description:
 */

package easy

import (
	"fmt"
	"testing"
	"github.com/Win-Man/go-leetcode/pkg"
)

func runningSum(nums []int) []int {
	var val int = 0
	for i := 0; i < len(nums); i++ {
		val += nums[i]
		nums[i] = val
	}
	return nums
}

func TestRunningSum(t *testing.T) {
	var tests = []struct {
		x    []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
		{[]int{2}, []int{2}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.x)
		t.Run(testname, func(t *testing.T) {
			ans := runningSum(tt.x)
			if !pkg.CompareArrayInt(ans,tt.want) {
				t.Errorf("runningSum(%v) got %v,want %v", tt.x, ans, tt.want)
			}
		})
	}
}
