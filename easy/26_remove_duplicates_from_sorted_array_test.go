/*
 * Created: 2020-02-24 19:02:41
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
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var index int = 1
	var lastNum int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > lastNum {
			nums[index] = nums[i]
			lastNum = nums[i]
			index++
		}
	}
	return index
}

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := removeDuplicates(tt.a)
			if ans != tt.want {
				t.Errorf("removeDuplicates(%v) got:%d want:%d", tt.a, ans, tt.want)
			}
		})
	}
}
