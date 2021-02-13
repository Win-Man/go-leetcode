/*
 * Created: 2021-02-13 17:55:06
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

func maximumWealth(accounts [][]int) int {
	var maxVal int = 0
	for _, vi := range accounts {
		var accounti int = 0
		for _, vj := range vi {
			accounti += vj
		}
		if accounti > maxVal {
			maxVal = accounti
		}
	}
	return maxVal
}

func TestMaximumWealth(t *testing.T) {
	var tests = []struct {
		x    [][]int
		want int
	}{
		{[][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{[][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
		{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.x)
		t.Run(testname, func(t *testing.T) {
			ans := maximumWealth(tt.x)
			if ans != tt.want {
				t.Errorf("maximumWealth(%v) got %d, want %d", tt.x, ans, tt.want)
			}
		})
	}
}
