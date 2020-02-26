/*
 * Created: 2020-02-26 20:55:17
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

func removeElement(nums []int, val int) int {
	var count int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		a    []int
		val  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := removeElement(tt.a, tt.val)
			if ans != tt.want {
				t.Errorf("removeElement(%v) got:%d want:%d", tt.a, ans, tt.want)
			}
		})
	}
}
